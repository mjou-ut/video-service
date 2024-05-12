package video

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type VideoTask struct {
	ID      string `json:"id"`
	StartAt uint32 `json:"startAt"`
	EndAt   uint32 `json:"endAt"`
	URL     string `json:"url"`
}

type Video struct {
	ID          string      `json:"id"`
	CreatedAtMs int64       `json:"createdAtMs"`
	DurationMs  int32       `json:"durationMs"`
	URL         string      `json:"url"`
	Tasks       []VideoTask `json:"tasks,omitempty"`
}

func NewVideo() Video {
	uuid := uuid.New().String()

	v := Video{
		ID:          uuid,
		CreatedAtMs: time.Now().Unix(),
		DurationMs:  rand.Int31n(10000-1000) + 1000,
		URL:         "s3://mybucket/" + uuid,
	}

	v.AddRandomTasks()

	return v
}

func (v *Video) AddTask(task VideoTask) {
	v.Tasks = append(v.Tasks, task)
}

func (v *Video) AddRandomTasks() {
	for i := 0; i < rand.Intn(10-1)+1; i++ {
		uuid := uuid.New().String()
		startAt := rand.Int31n(v.DurationMs - 1000)
		endAt := startAt + rand.Int31n(v.DurationMs-startAt)

		task := VideoTask{
			ID:      uuid,
			StartAt: uint32(startAt),
			EndAt:   uint32(endAt),
			URL:     "s3://mybucket/" + v.ID + "/" + uuid,
		}
		v.AddTask(task)
	}
}
