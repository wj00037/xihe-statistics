package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForMediaController(
	rg *gin.RouterGroup,
	repo repository.Media,
) {
	ctl := MediaController{
		s: app.NewMeidaService(repo),
	}

	rg.GET("/v1/media", ctl.Get)
	rg.POST("/v1/media", ctl.Add)
}

type MediaController struct {
	baseController
	s app.MediaService
}

// @Summary Get
// @Description get a media info
// @Tags  Media
// @Param	bigmodel	path	string	true	"type of bigmodel"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/media [get]
func (ctl *MediaController) Get(ctx *gin.Context) {
	dto, err := ctl.s.GetAll()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}

// @Summary Add
// @Description add media data
// @Tags  Media
// @Param  body  body  BigModelCreateRequest  true  "body of bigmodel records"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/media [post]
func (ctl *MediaController) Add(ctx *gin.Context) {
	req := MediaRequest{}
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

	err = ctl.s.Add(&cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData("success"))
}
