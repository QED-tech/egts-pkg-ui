package egts

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qed-tech/egts-pkg-debugger/internal/usecase"
	"github.com/qed-tech/egts-pkg-debugger/pkg/server"
)

func (h Handler) Decode(c echo.Context) error {
	pkg := c.QueryParam("package")
	if pkg == "" {
		return c.JSON(http.StatusBadRequest, server.Problem{
			Title:  "Invalid package",
			Status: http.StatusBadRequest,
		})
	}

	out, err := h.decodeUsecase.Decode(c.Request().Context(), pkg)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidHexString):
			return c.JSON(http.StatusBadRequest, server.Problem{
				Title:  "Invalid package",
				Status: http.StatusBadRequest,
			})
		default:
			return c.JSON(http.StatusInternalServerError, server.Problem{
				Title:  "Internal server error",
				Status: http.StatusInternalServerError,
			})
		}
	}

	return c.JSON(http.StatusOK, server.Response{
		Data: out,
	})
}
