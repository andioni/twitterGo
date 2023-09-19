package main

import (
	"context"
	"os"
	"strings"
	"twitterGo/awsgo"
	"twitterGo/bd"
	"twitterGo/handlers"
	"twitterGo/models"
	"twitterGo/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse
	awsgo.InicializoAWS()
	if !ValidoParametros() {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las vaiables de entrada",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error de lectura de secret" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}
	path := strings.Replace(req.PathParameters["twittercursogo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), req.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), req.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// CHECK DB CONNECTION
	err = bd.ConectarBD(awsgo.Ctx)

	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error de conexion a la base de datos" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

	respApi := handlers.Manejaradores(awsgo.Ctx, req)

	if respApi.CustomResp == nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}
	return respApi.CustomResp, nil
}

func ValidoParametros() bool {
	_, traeParam := os.LookupEnv("SecretName")
	if !traeParam {
		return traeParam
	}

	_, traeParam = os.LookupEnv("BucketName")
	if !traeParam {
		return traeParam
	}

	_, traeParam = os.LookupEnv("UrlPrefix")
	if !traeParam {
		return traeParam
	}

	return true
}
