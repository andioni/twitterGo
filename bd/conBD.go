package bd

import (
	"context"
	"fmt"

	"twitterGo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DataBaseName string

func ConectarBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&wmajority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion realizad con la base de datos")

	MongoCN = client
	DataBaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
