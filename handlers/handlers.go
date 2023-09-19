package handlers

import (
	"context"
	"fmt"

	"twitterGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func Manejaradores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {

	path := ctx.Value(models.Key("path")).(string)
	fmt.Println("Procesar" + path + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi

	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
	case "GET":
	case "PUT":
	case "DELETE":
	}

	r.Message = "Method invalid"

	return r
}
