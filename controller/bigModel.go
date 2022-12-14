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

	rg.POST("/v1/d1/bigmodel", ctl.AddBigModel)
	rg.GET("/v1/d1/bigmodel/:bigmodel", ctl.Get)
	rg.GET("/v1/d1/bigmodel", ctl.GetAll)

}

type BigModelRecordController struct {
	baseController
	bs app.BigModelRecordService
}

// @Summary Add
// @Description add user query bigmodel record
// @Tags  D1
// @Param  body  body  QueryBigModelRequest  true  "body of bigmodel records"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/bigmodel [post]
func (ctl *BigModelRecordController) AddBigModel(ctx *gin.Context) {
	req := QueryBigModelRequest{}
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

	err = ctl.bs.AddUserWithBigModel(&cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData("success"))
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
