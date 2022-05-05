package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ksarsecha/gin-demo/entity"
	"github.com/ksarsecha/gin-demo/service"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func (vc *videoController) Save(ctx *gin.Context) (video entity.Video) {
	if err := ctx.BindJSON(&video); err != nil {
		return video
	}
	return vc.service.Save(video)
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func New(service service.VideoService) VideoController {
	return &videoController{service: service}
}
