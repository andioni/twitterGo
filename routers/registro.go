package routers

import (
	"context"
	"encoding/json"

	"fmt"

	"twitterGo/bd"
	"twitterGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi

	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el mail"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña valida"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe el uauario"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "No se ha podido insertar " + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "No se ha podido insertar " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)

	return r

}
