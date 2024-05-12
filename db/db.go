package db

import (
	"fmt"

	video "github.com/mjou-ut/video-service/models"
)

var DB = make(map[string]video.Video)

func GetVideo(id string) (*video.Video, error) {
	if v, found := DB[id]; found {
		return &v, nil
	}
	return nil, fmt.Errorf("video not found")
}

func GetVideos() []video.Video {
	videos := []video.Video{}
	index := 0
	for _, v := range DB {
		v.Tasks = nil
		videos = append(videos, v)
		// max 100 videos
		if index > 100 {
			break
		}
		index++
	}
	return videos
}

func SaveVideo(v video.Video) {
	DB[v.ID] = v
}

func deleteVideo(id string) {
	delete(DB, id)
}
