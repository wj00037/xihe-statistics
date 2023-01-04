package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"
)

func AddRouterForTrainRecordController(
	rg *gin.RouterGroup,
	tr repository.TrainRecord,
) {
	ctl := TrainRecordController{
		tr: app.NewTrainRecordService(tr),
	}

	rg.GET("/v1/train", ctl.Get)
}

type TrainRecordController struct {
	baseController
	tr app.TrainRecordService
}

// @Summary Get
// @Description get trained user counts
// @Tags  Train
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/train [get]
func (ctl *TrainRecordController) Get(ctx *gin.Context) {

	dto, err := ctl.tr.Get()

	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
