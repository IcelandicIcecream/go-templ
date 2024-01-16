package handler

import (
	"github.com/icelandicicecream/htmx-go/model"
	"github.com/icelandicicecream/htmx-go/views/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "benjaminsinidol@gmail.com",
	}
	return render(c, user.Show(u))
}
