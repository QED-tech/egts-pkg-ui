package egts

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qed-tech/egts-pkg-debugger/pkg/server"
)

func (h Handler) GetHistory(c echo.Context) error {
	ctx := c.Request().Context()

	out, err := h.getHistoryUsecase.Handle(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, server.Problem{
			Title:  "Internal server error",
			Status: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, server.Response{
		Data: out,
	})
}
