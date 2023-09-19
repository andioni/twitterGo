package handlers

import (
	"context"
	"fmt"

	"twitterGo/jwt"
	"twitterGo/models"
	"twitterGo/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Manejaradores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {

	path := ctx.Value(models.Key("path")).(string)
	fmt.Println("Procesar" + path + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi

	r.Status = 400

	isOk, statusCode, msg, _ := validoAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch path {
		case "registro":
			routers.Registro(ctx)
		}
	case "GET":
	case "PUT":
	case "DELETE":
	}

	r.Message = "Method invalid"

	return r
}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}
	fmt.Println("Token ok")
	return true, 200, msg, *claim
}
