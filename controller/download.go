package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForDownloadRecordController(
	rg *gin.RouterGroup,
	dr repository.DownloadRecord,
) {
	ctl := DownloadRecordController{
		ds: app.NewDownloadRecordService(dr),
	}

	rg.GET("/v1/download", ctl.Get)
}

type DownloadRecordController struct {
	baseController
	ds app.DownloadRecordService
}

// @Summary Get
// @Description get download record
// @Tags  download
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/download [get]
func (ctl *DownloadRecordController) Get(ctx *gin.Context) {
	dto, err := ctl.ds.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
