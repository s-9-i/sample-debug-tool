package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sample-debug-tool-go/pkg/api/command"
)

type Handler interface {
	command.Handler
	UpdateLevel(c echo.Context) error
	SkipTutorial(c echo.Context) error
}

type handlerImpl struct{}

func New() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) GetBaseCommandGroup() *command.Group {
	return &command.Group{
		Name: "ユーザー",
		URL:  "/user",
		Commands: command.Commands{
			{
				HandlerFunc: h.UpdateLevel,
				URL:         "/updateLevel",
				Name:        "レベル操作",
				Description: "ユーザーのレベルを更新します。",
				Model:       UpdateLevelRequest{},
			},
			{
				HandlerFunc: h.SkipTutorial,
				URL:         "/skipTutorial",
				Name:        "チュートリアル突破",
				Description: "ユーザーがチュートリアルを突破した状態にします。",
				Model:       SkipTutorialRequest{},
			},
		},
	}
}

type UpdateLevelRequest struct {
	UserID string `json:"userID" validate:"required"`
	Level  int32  `json:"level" validate:"required"`
}

func (h *handlerImpl) UpdateLevel(c echo.Context) error {
	r := &UpdateLevelRequest{}
	if err := c.Bind(r); err != nil {
		return err
	}

	// レベル操作の実処理

	return c.NoContent(http.StatusOK)
}

type SkipTutorialRequest struct {
	UserID string `json:"userID" validate:"required"`
}

func (h *handlerImpl) SkipTutorial(c echo.Context) error {
	r := &SkipTutorialRequest{}
	if err := c.Bind(r); err != nil {
		return err
	}

	// チュートリアル突破操作の実処理

	return c.NoContent(http.StatusOK)
}
