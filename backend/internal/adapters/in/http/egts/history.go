package egts

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qed-tech/egts-pkg-debugger/internal/entity"
	"github.com/qed-tech/egts-pkg-debugger/pkg/server"
	"github.com/samber/lo"
)

type PackageHistoryRecord struct {
	Hex string `json:"hex"`
}

func (h Handler) GetHistory(c echo.Context) error {
	ctx := c.Request().Context()

	historyRecords, err := h.getHistoryUsecase.Handle(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get history records", slog.Any("err", err))
		return c.JSON(http.StatusInternalServerError, server.Problem{
			Title:  "Internal server error",
			Status: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, server.Response{
		Data: lo.Map(historyRecords, func(record *entity.HistoryRecord, _ int) *PackageHistoryRecord {
			return &PackageHistoryRecord{
				Hex: record.EgtsPackageHex,
			}
		}),
	})
}
