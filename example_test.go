package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var cli *Client
var db *DB

func init() {
	var err error
	cli, err = NewClient(context.Background(), options.Client().ApplyURI(os.Getenv("MOGO")))
	if err != nil {
		panic(err)
	}
	db = cli.NewDB("cloud-desktop")
}

type diskTop struct {
	Id     string `bson:"_id"`
	UserID string `bson:"user_id"`
}

func (diskTop) CollName() string {
	return "device"
}

func TestCollection_Aggregate(t *testing.T) {
	list := make([]diskTop, 0)
	c, err := db.Model(diskTop{}).Aggregate(context.Background()).FindPageList(1, 20, &list)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(c, list)
	}
}
