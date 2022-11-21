package mogo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Coll interface {
	CollName() string
}

type Collection struct {
	coll *mongo.Collection
}

func (c *Collection) FindByID(ctx context.Context, id string, data any) error {
	val, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	err = c.coll.FindOne(ctx, bson.M{"_id": val}).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	return err
}

func (c *Collection) Find(ctx context.Context, filter bson.M, data any) error {
	cur, err := c.coll.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, data)
	return err
}

func (c *Collection) FindOne(ctx context.Context, filter bson.M, data any) error {
	err := c.coll.FindOne(ctx, filter).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (c *Collection) InsertOne(ctx context.Context, doc any) error {
	_, err := c.coll.InsertOne(ctx, doc)
	return err
}

func (c *Collection) InsertMany(ctx context.Context, docs []any) error {
	_, err := c.coll.InsertMany(ctx, docs)
	return err
}

func (c *Collection) DeleteOne(ctx context.Context, filter bson.M) error {
	_, err := c.coll.DeleteOne(ctx, filter)
	return err
}

func (c *Collection) DeleteMany(ctx context.Context, filter bson.M) error {
	_, err := c.coll.DeleteMany(ctx, filter)
	return err
}

func (c *Collection) UpdateOne(ctx context.Context, filter bson.M, data any) error {
	_, err := c.coll.UpdateOne(ctx, filter, data)
	return err
}

func (c *Collection) UpdateMany(ctx context.Context, filter bson.M, data any) error {
	_, err := c.coll.UpdateMany(ctx, filter, data)
	return err
}

func (c *Collection) ReplaceOne(ctx context.Context, filter bson.M, data any) error {
	_, err := c.coll.ReplaceOne(ctx, filter, data)
	return err
}

func (c *Collection) Aggregate(ctx context.Context, pipes []bson.D, data any) error {
	cur, err := c.coll.Aggregate(ctx, pipes)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, data)
	return err
}
