package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/platform"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForDownloadRecordController(
	rg *gin.RouterGroup,
	dr repository.DownloadRecord,
	pf platform.PlatForm,
	gl repository.Gitlab,
) {
	ctl := DownloadRecordController{
		ds: app.NewDownloadRecordService(dr),
		gs: app.NewGitLabService(pf, gl),
	}

	rg.GET("/v1/download", ctl.Get)
	rg.GET("/v1/download/clone", ctl.GetClone)
}

type DownloadRecordController struct {
	baseController
	ds app.DownloadRecordService
	gs app.GitLabService
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

// @Summary Get Clones
// @Description get Clone record
// @Tags  download
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/download/clone [get]
func (ctl *DownloadRecordController) GetClone(ctx *gin.Context) {
	dto, err := ctl.gs.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}