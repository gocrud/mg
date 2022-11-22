package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	cli *mongo.Client
}

func NewClient(ctx context.Context, clientOptions *options.ClientOptions) (*Client, error) {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) NewDB(name string) *DB {
	return &DB{db: c.cli.Database(name)}
}

func (c *Client) Transaction(fn func(sessionContext mongo.SessionContext) error) error {
	session, err := c.cli.StartSession()
	if err != nil {
		return err
	}
	defer func() {
		session.EndSession(context.Background())
	}()

	var f = func(sessionContext mongo.SessionContext) (interface{}, error) {
		err := fn(sessionContext)
		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), f)
	return err
}
