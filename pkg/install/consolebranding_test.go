package install

import (
	"context"
	"testing"

	v1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/client-go/operator/clientset/versioned/fake"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestUpdateConsoleBranding(t *testing.T) {
	ctx := context.Background()

	// read up on how to use Reactors here to be able to customise the fake
	// behaviour in more detail
	operatorcli := fake.NewSimpleClientset(&v1.Console{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
		},
		Status: v1.ConsoleStatus{
			OperatorStatus: v1.OperatorStatus{
				Conditions: []v1.OperatorCondition{
					{
						Type:   "DeploymentAvailable",
						Status: v1.ConditionTrue,
					},
				},
			},
		},
	})

	i := &Installer{
		log:         logrus.NewEntry(logrus.StandardLogger()),
		operatorcli: operatorcli,
	}

	err := i.updateConsoleBranding(ctx)
	if err != nil {
		t.Error(err)
	}

	console, err := operatorcli.OperatorV1().Consoles().Get("cluster", metav1.GetOptions{})
	if err != nil {
		t.Error(err)
	}

	if console.Spec.Customization.Brand != "azure" {
		t.Error(console.Spec.Customization.Brand)
	}
}
