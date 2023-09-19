package bd

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DataBaseName)
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)

	if err != nil {
		return resultado, false, resultado.ID.Hex()
	}
	return resultado, true, resultado.ID.Hex()

}
