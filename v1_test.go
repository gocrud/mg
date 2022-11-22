package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func init() {
	var err error
	cli, err = NewClient(context.Background(), options.Client().ApplyURI("mongodb://root:123456@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/admin?replicaSet=rs0"))
	if err != nil {
		panic(err)
	}
}

var cli *Client

type User struct {
	Id   primitive.ObjectID
	Name string
	Age  int
}

func (User) CollName() string {
	return "user"
}

func TestCollection_InsertOne(t *testing.T) {
	db := cli.NewDB("testdb")
	u := User{
		Name: "王五",
		Age:  22,
	}
	err := db.Coll("user").InsertOne(context.Background(), u)
	if err != nil {
		t.Error(err)
	}
}

func TestCollection_Find(t *testing.T) {
	db := cli.NewDB("testdb")
	var ls []User
	err := db.Model(User{}).Find(context.Background(), bson.M{"name": "Petter"}, &ls)
	if err != nil {
		t.Error(err)
	}
	t.Log(ls)
}

func TestCollection_FindByID(t *testing.T) {
	db := cli.NewDB("testdb")
	var u User
	err := db.Model(u).FindByID(context.Background(), "636600f02789a22e1d237d46", &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

func TestClient_Transaction(t *testing.T) {
	db := cli.NewDB("testdb")
	u := User{
		Name: "ZhangSan",
		Age:  25,
	}
	u2 := User{
		Name: "LiSi",
		Age:  30,
	}
	err := cli.Transaction(func(sessionContext mongo.SessionContext) error {
		if err := db.Model(u).InsertOne(sessionContext, u); err != nil {
			return err
		}
		if err := db.Model(u).InsertOne(sessionContext, u2); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCollection_UpdateOne(t *testing.T) {
	db := cli.NewDB("testdb")
	err := db.Model(User{}).UpdateOne(context.Background(), bson.M{"name": "LiSi"}, bson.M{"$set": bson.M{"age": 100}})
	if err != nil {
		t.Error(err)
	}
}

func TestCollection_FindOne(t *testing.T) {
	db := cli.NewDB("testdb")
	u := User{}
	err := db.Model(u).FindOne(context.Background(), bson.M{"name": "LiSi"}, &u)
	if err != nil {
		t.Error(err)
	}
}

func TestGetStrId(t *testing.T) {
	t.Log(GetStrId())
}
