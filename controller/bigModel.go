package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/user"

	"github.com/gin-gonic/gin"
)

func AddRouterForBigModelRecordController(
	rg *gin.RouterGroup,
	ub user.UserWithBigModel,
) {
	ctl := BigModelRecordController{
		bs: app.NewBigModelRecordService(ub),
	}

	rg.POST("/v1/d1/bigmodel", ctl.AddBigModel)
}

type BigModelRecordController struct {
	baseController
	bs app.BigModelRecordService
}

// @Summary Check
// @Description add user query bigmodel record
// @Tags  D1
// @Param	type	path	string	true	"owner of project"
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
