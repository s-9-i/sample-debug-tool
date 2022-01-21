package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"sample-debug-tool-go/pkg/api/command"
	"sample-debug-tool-go/pkg/api/handler/card"
	"sample-debug-tool-go/pkg/api/handler/user"
)

func Serve() {
	e := echo.New()
	e.Use(middleware.CORS())

	commandHandlers := []command.Handler{
		user.New(),
		card.New(),
	}

	// デバッグAPI群
	groups := make(command.Groups, 0, len(commandHandlers))
	for _, commandHandler := range commandHandlers {
		group := commandHandler.GetBaseCommandGroup()
		command.InitCommand(group)
		for _, command := range group.Commands {
			e.POST(group.URL+command.URL, command.HandlerFunc)
		}
		groups = append(groups, group)
	}

	// デバッグAPI一覧取得API
	e.GET("/api/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, groups)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
