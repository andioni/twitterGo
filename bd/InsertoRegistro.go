package bd

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DataBaseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, error := col.InsertOne(ctx, u)

	if error != nil {
		return "", false, error
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}
