package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForCloudRecordController(
	rg *gin.RouterGroup,
	cr repository.CloudRecord,
) {
	ctl := CloudRecordController{
		cs: app.NewCloudRecodeService(cr),
	}

	rg.GET("/v1/cloud", ctl.Get)
}

type CloudRecordController struct {
	baseController

	cs app.CloudRecordService
}

// @Summary Get
// @Description get cloud record
// @Tags  cloud
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/cloud [get]
func (ctl *CloudRecordController) Get(ctx *gin.Context) {
	dto, err := ctl.cs.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
