package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForFileUploadRecordController(
	rg *gin.RouterGroup,
	fr repository.FileUploadRecord,
) {
	ctl := FileUploadRecordController{
		fs: app.NewFileUploadRecordService(fr),
	}

	rg.GET("/v1/d2", ctl.GetFileUploadRecord)

}

type FileUploadRecordController struct {
	baseController
	fs app.FileUploadRecordService
}

// @Summary Get
// @Description get d2
// @Tags  D2
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d2 [get]
func (ctl *FileUploadRecordController) GetFileUploadRecord(ctx *gin.Context) {
	dto, err := ctl.fs.GetUsersCounts()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
