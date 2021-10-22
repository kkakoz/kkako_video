package handler

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "kkako_video/api/video/v1"
	"kkako_video/internal/video/domain"
)

type VideoHandler struct {
	v1.UnimplementedVideoServiceServer
	videoLogic domain.IVideoLogic
}

func (v VideoHandler) GetVideo(ctx context.Context, req *v1.GetVideoReq) (*v1.GetVideoRes, error) {
	video, err := v.videoLogic.GetVideo(ctx, req.Id)
	res := &v1.GetVideoRes{}
	err = copier.Copy(res, video)
	return res, err
}

func (v VideoHandler) GetVideos(ctx context.Context, req *v1.VideoNewsReq) (*v1.VideoNewsRes, error) {
	panic("implement me")
}

func (v VideoHandler) AddVideo(ctx context.Context, req *v1.VideoNewsReq) (*v1.VideoNewsRes, error) {
	panic("implement me")
}

