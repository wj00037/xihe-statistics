package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForBigModelRecordController(
	rg *gin.RouterGroup,
	ub repository.UserWithBigModel,
) {
	ctl := BigModelRecordController{
		bs: app.NewBigModelRecordService(ub),
	}

	rg.GET("/v1/d1/bigmodel/:bigmodel", ctl.Get)
	rg.POST("/v1/d1/bigmodel/increase", ctl.GetIncrease)
	rg.GET("/v1/d1/bigmodel", ctl.GetAll)

}

type BigModelRecordController struct {
	baseController
	bs app.BigModelRecordService
}

// @Summary Get
// @Description get a bigmodel query records
// @Tags  D1
// @Param	bigmodel	path	string	true	"type of bigmodel"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/bigmodel/{bigmodel} [get]
func (ctl *BigModelRecordController) Get(ctx *gin.Context) {
	bigmodel, err := domain.NewBigModel(ctx.Param("bigmodel"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(
			errorBadRequestParam, err,
		))

		return
	}

	bmd, err := ctl.bs.GetBigModelRecordsByType(bigmodel)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, newResponseData(bmd))
}

// @Summary GetIncrease
// @Description get increase d1 user count during time
// @Tags  D1
// @Param  body  body  BigModelQueryWithTypeAndTimeRequest  true  "body of time and bigmodel type"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/bigmodel/increase [post]
func (ctl *BigModelRecordController) GetIncrease(ctx *gin.Context) {
	req := BigModelQueryWithTypeAndTimeRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
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

	dto, err := ctl.bs.GetCountsByTypeAndTimeDiff(cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}

// @Summary GetAll
// @Description get all bigmodel query records
// @Tags  D1
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/bigmodel [get]
func (ctl *BigModelRecordController) GetAll(ctx *gin.Context) {

	bmd, err := ctl.bs.GetBigModelRecordAll()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, newResponseData(bmd))
}
