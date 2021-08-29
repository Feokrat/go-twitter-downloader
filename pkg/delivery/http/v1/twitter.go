package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTwitterRoutes(api *gin.RouterGroup) {
	group := api.Group("/twitter")
	{
		group.GET("downloadImages", h.downloadImages)
	}
}

// @Summary Загрузка картинок по ссылке
// @Tags Twitter
// @Description Скачивает картинки твита по ссылке
// @Accept  json
// @Produce  json
// @Param userId query string true "Tweet Url"
// @Success 200 {object} outputGetUserGroups
// @Failure 400 {object} HTTPResponse
// @Failure 500 {object} HTTPResponse
// @Router /api/v1/twitter/downloadImages [get]
func (h *Handler) downloadImages(ctx *gin.Context) {
	tweetUrl := ctx.Query("tweetUrl")
	downloadPath := ctx.Query("downloadPath")
	photos, err := h.services.Twitter.DownloadImages(tweetUrl, downloadPath)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}
