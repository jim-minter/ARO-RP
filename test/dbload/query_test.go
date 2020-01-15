package dbload

import (
	"context"
	"strconv"
	"testing"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	utillog "github.com/Azure/ARO-RP/pkg/util/log"
)

func TestQuery(t *testing.T) {
	ctx := context.Background()
	log := utillog.GetLogger()

	collc, c, err := get(ctx, log)
	if err != nil {
		t.Fatal(err)
	}

	partitions, err := collc.PartitionKeyRanges(ctx, "OpenShiftClusters")
	if err != nil {
		t.Fatal(err)
	}

	var countTotal int
	for _, r := range partitions.PartitionKeyRanges {
		result := c.Query("", &cosmosdb.Query{
			Query: `SELECT VALUE COUNT(1) FROM OpenShiftClusters doc WHERE doc.openShiftCluster.properties.provisioningState IN ("Creating", "Deleting", "Updating") AND (doc.leaseExpires ?? 0) < GetCurrentTimestamp() / 1000`,
		}, &cosmosdb.Options{
			PartitionKeyRangeID: r.ID,
		})
		// because we aggregate count we don't expect pagination in this query result,
		// so we gonna call Next() only once per partition.
		var data struct {
			api.MissingFields
			Document []string `json:"Documents,omitempty"`
		}
		err := result.NextRaw(ctx, &data)
		if err != nil {
			t.Fatal(err)
		}

		count, err := strconv.Atoi(data.Document[0])
		if err != nil {
			t.Fatal(err)
		}
		countTotal = countTotal + count
	}

	t.Log(countTotal)
}
