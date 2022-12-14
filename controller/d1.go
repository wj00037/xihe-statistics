package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForD1Controller(
	rg *gin.RouterGroup,
	bm repository.UserWithBigModel,
	rr repository.UserWithRepo,
) {
	ctl := D1Controller{
		ds: app.NewD1Service(bm, rr),
	}

	rg.GET("/v1/d1", ctl.Get)

}

type D1Controller struct {
	baseController
	ds app.D1Service
}

// @Summary Get
// @Description Get D1
// @Tags  D1
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1 [Get]
func (ctl *D1Controller) Get(ctx *gin.Context) {

	dto, err := ctl.ds.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
