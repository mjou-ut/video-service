package db

import (
	"fmt"

	video "github.com/mjou-ut/video-service/models"
)

var DB = make(map[string]video.Video)

func GetVideo(id string) (video.Video, error) {
	if _, ok := DB[id]; !ok {
		return video.Video{}, fmt.Errorf("video not found")
	}
	return DB[id], nil
}

func GetVideos() []video.Video {
	videos := []video.Video{}
	index := 0
	for _, v := range DB {
		v.Tasks = nil
		videos = append(videos, v)
		// max 10 videos
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
