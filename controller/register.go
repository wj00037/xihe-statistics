package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"
)

func AddRouterForRegisterRecordController(
	rg *gin.RouterGroup,
	rr repository.RegisterRecord,
) {
	ctl := RegisterRecordController{
		rs: app.NewRegisterRecordService(rr),
	}

	rg.GET("/v1/d0", ctl.GetRegisterRecord)
}

type RegisterRecordController struct {
	baseController
	rs app.RegisterRecordService
}

// @Summary Get
// @Description get user register records
// @Tags  D0
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d0 [get]
func (ctl *RegisterRecordController) GetRegisterRecord(ctx *gin.Context) {

	dto, err := ctl.rs.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
