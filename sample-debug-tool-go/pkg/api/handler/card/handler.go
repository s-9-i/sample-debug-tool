package card

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sample-debug-tool-go/pkg/api/command"
)

type Handler interface {
	command.Handler
	UpdateRarity(c echo.Context) error
}

type handlerImpl struct{}

func New() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) GetBaseCommandGroup() *command.Group {
	return &command.Group{
		Name: "カード",
		URL:  "/card",
		Commands: command.Commands{
			{
				HandlerFunc: h.UpdateRarity,
				URL:         "/updateRarity",
				Name:        "レアリティ操作",
				Description: "所持カードのレアリティを更新します。",
				Model:       UpdateRarityRequest{},
			},
		},
	}
}

type UpdateRarityRequest struct {
	UserID string `json:"userID" validate:"required"`
	CardID string `json:"cardID" validate:"required" master:"Card" type:"select"`
	Rarity int32  `json:"rarity" validate:"required" enum:"RarityType" type:"radios"`
}

func (h *handlerImpl) UpdateRarity(c echo.Context) error {
	r := &UpdateRarityRequest{}
	if err := c.Bind(r); err != nil {
		return err
	}

	// 所持カードのレアリティ操作の実処理

	return c.NoContent(http.StatusOK)
}
