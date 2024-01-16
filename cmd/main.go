package main

import (
	"context"
	"fmt"

	"github.com/icelandicicecream/htmx-go/handler"
	"github.com/icelandicicecream/htmx-go/model"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/styles", "styles")
	app.Use(withUser)

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	err := app.Start(":3000")
	if err != nil {
		panic(err)
	}
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctxKVP := model.KeyValuePair{
			Key:   "user",
			Value: "benjaminsinidol@gmail.com",
		}

		fmt.Println("HELLO")

		ctx := context.WithValue(c.Request().Context(), ctxKVP.Key, ctxKVP.Value)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
