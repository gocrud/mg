package mg

import (
	"context"
	"github.com/gocrud/mg/pipe"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
)

var cli *Client
var db *DB

func init() {
	var err error
	cli, err = NewClient(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		panic(err)
	}
	db = cli.NewDB("platform")
}

type Admin struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
}

func TestAggregate_Pipe(t *testing.T) {
	var md []Admin
	match := pipe.BuildMatcher()
	match.IF(true, bson.E{"_id", ""})
	err := db.Coll("admin").Aggregate(context.Background(), match.Build()).Find(&md)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", md)
}
