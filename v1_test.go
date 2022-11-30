package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

var cli *Client

type User struct {
	Name string
	Age  int
}

func (User) CollName() string {
	return "user"
}

func TestCollection_InsertOne(t *testing.T) {
	db := cli.NewDB("dev")
	u := User{
		Name: "王五",
		Age:  22,
	}
	id, err := db.Coll("user").InsertOne(context.Background(), u)
	if err != nil {
		t.Error(err)
	}
	log.Println(id.Hex())
}

func TestCollection_InsertMany(t *testing.T) {
	db := cli.NewDB("dev")
	u := User{
		Name: "王五",
		Age:  22,
	}
	u1 := User{
		Name: "李四",
		Age:  22,
	}
	ids, err := db.Coll("user").InsertMany(context.Background(), []any{u, u1})
	if err != nil {
		t.Error(err)
	}
	log.Println(ids)
}

func TestCollection_Find(t *testing.T) {
	db := cli.NewDB("dev")
	var ls []User
	err := db.Model(User{}).Find(context.Background(), bson.M{"name": "Petter"}, &ls)
	if err != nil {
		t.Error(err)
	}
	t.Log(ls)
}

func TestCollection_FindByID(t *testing.T) {
	db := cli.NewDB("dev")
	var u User
	err := db.Model(u).FindByID(context.Background(), "636600f02789a22e1d237d46", &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
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
