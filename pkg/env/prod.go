package env

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/deploy/generator"
	basekeyvault "github.com/Azure/ARO-RP/pkg/util/azureclient/keyvault"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/compute"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/dns"
	"github.com/Azure/ARO-RP/pkg/util/instancemetadata"
	"github.com/Azure/ARO-RP/pkg/util/refreshable"
)

type prod struct {
	instancemetadata.InstanceMetadata
	ServiceKeyvaultInterface

	keyvault basekeyvault.BaseClient

	acrName             string
	clustersKeyvaultURI string
	domain              string
	zones               map[string][]string

	fpCertificate        *x509.Certificate
	fpPrivateKey         *rsa.PrivateKey
	fpServicePrincipalID string

	clustersGenevaLoggingCertificate *x509.Certificate
	clustersGenevaLoggingPrivateKey  *rsa.PrivateKey

	log     *logrus.Entry
	envType Type
}

func newProd(ctx context.Context, log *logrus.Entry, instancemetadata instancemetadata.InstanceMetadata) (*prod, error) {
	kvAuthorizer, err := RPAuthorizer(azure.PublicCloud.ResourceIdentifiers.KeyVault)
	if err != nil {
		return nil, err
	}

	p := &prod{
		InstanceMetadata: instancemetadata,

		keyvault: basekeyvault.New(kvAuthorizer),

		log:     log,
		envType: Prod,
	}

	p.ServiceKeyvaultInterface, err = NewServiceKeyvault(ctx, p)
	if err != nil {
		return nil, err
	}

	rpAuthorizer, err := RPAuthorizer(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	err = p.populateDomain(ctx, rpAuthorizer)
	if err != nil {
		return nil, err
	}

	p.clustersKeyvaultURI, err = getVaultURI(ctx, instancemetadata, generator.ClustersKeyVaultTagValue)
	if err != nil {
		return nil, err
	}

	err = p.populateZones(ctx, rpAuthorizer)
	if err != nil {
		return nil, err
	}

	fpPrivateKey, fpCertificates, err := p.GetCertificateSecret(ctx, RPFirstPartySecretName)
	if err != nil {
		return nil, err
	}

	p.fpPrivateKey = fpPrivateKey
	p.fpCertificate = fpCertificates[0]
	p.fpServicePrincipalID = "f1dd0a37-89c6-4e07-bcd1-ffd3d43d8875"

	clustersGenevaLoggingPrivateKey, clustersGenevaLoggingCertificates, err := p.GetCertificateSecret(ctx, ClusterLoggingSecretName)
	if err != nil {
		return nil, err
	}

	p.clustersGenevaLoggingPrivateKey = clustersGenevaLoggingPrivateKey
	p.clustersGenevaLoggingCertificate = clustersGenevaLoggingCertificates[0]

	if p.ACRResourceID() != "" { // TODO: ugh!
		acrResource, err := azure.ParseResourceID(p.ACRResourceID())
		if err != nil {
			return nil, err
		}
		p.acrName = acrResource.ResourceName
	} else {
		p.acrName = "arointsvc"
	}

	return p, nil
}

func (p *prod) ACRResourceID() string {
	return os.Getenv("ACR_RESOURCE_ID")
}

func (p *prod) ACRName() string {
	return p.acrName
}

func (p *prod) populateDomain(ctx context.Context, rpAuthorizer autorest.Authorizer) error {
	zones := dns.NewZonesClient(p.SubscriptionID(), rpAuthorizer)

	zs, err := zones.ListByResourceGroup(ctx, p.ResourceGroup(), nil)
	if err != nil {
		return err
	}

	if len(zs) != 1 {
		return fmt.Errorf("found %d zones, expected 1", len(zs))
	}

	p.domain = *zs[0].Name

	return nil
}

func (p *prod) populateZones(ctx context.Context, rpAuthorizer autorest.Authorizer) error {
	c := compute.NewResourceSkusClient(p.SubscriptionID(), rpAuthorizer)

	skus, err := c.List(ctx, "")
	if err != nil {
		return err
	}

	p.zones = map[string][]string{}

	for _, sku := range skus {
		if !strings.EqualFold((*sku.Locations)[0], p.Location()) ||
			*sku.ResourceType != "virtualMachines" {
			continue
		}

		p.zones[*sku.Name] = *(*sku.LocationInfo)[0].Zones
	}

	return nil
}

func (p *prod) ClustersGenevaLoggingSecret() (*rsa.PrivateKey, *x509.Certificate) {
	return p.clustersGenevaLoggingPrivateKey, p.clustersGenevaLoggingCertificate
}

func (p *prod) ClustersKeyvaultURI() string {
	return p.clustersKeyvaultURI
}

func (p *prod) Domain() string {
	return p.domain
}

func (p *prod) FPAuthorizer(tenantID, resource string) (refreshable.Authorizer, error) {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, err
	}

	sp, err := adal.NewServicePrincipalTokenFromCertificate(*oauthConfig, p.fpServicePrincipalID, p.fpCertificate, p.fpPrivateKey, resource)
	if err != nil {
		return nil, err
	}

	return refreshable.NewAuthorizer(sp), nil
}

// ManagedDomain returns the fully qualified domain of a cluster if we manage
// it.  If we don't, it returns the empty string.  We manage only domains of the
// form "foo.$LOCATION.aroapp.io" and "foo" (we consider this a short form of
// the former).
func (p *prod) ManagedDomain(domain string) (string, error) {
	if domain == "" ||
		strings.HasPrefix(domain, ".") ||
		strings.HasSuffix(domain, ".") {
		// belt and braces: validation should already prevent this
		return "", fmt.Errorf("invalid domain %q", domain)
	}

	domain = strings.TrimSuffix(domain, "."+p.Domain())
	if strings.ContainsRune(domain, '.') {
		return "", nil
	}
	return domain + "." + p.Domain(), nil
}

func (p *prod) Zones(vmSize string) ([]string, error) {
	zones, found := p.zones[vmSize]
	if !found {
		return nil, fmt.Errorf("zone information not found for vm size %q", vmSize)
	}
	return zones, nil
}

func (d *prod) CreateARMResourceGroupRoleAssignment(ctx context.Context, fpAuthorizer refreshable.Authorizer, resourceGroup string) error {
	// ARM ResourceGroup role assignments are not required in production.
	return nil
}

func (p *prod) Type() Type {
	return p.envType
}
