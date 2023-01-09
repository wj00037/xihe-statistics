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
	rg.POST("/v1/train/increase", ctl.Get)

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

// @Summary GetIncrease
// @Description get trained user counts during the time
// @Tags  Train
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/train/increase [post]
func (ctl *TrainRecordController) GetIncrease(ctx *gin.Context) {
	req := TrainIncreaseRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, respBadRequestBody)
		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(
			errorBadRequestParam, err,
		))
		return
	}

	dto, err := ctl.tr.GetTrains(cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
