package models

import (
	"database/sql/driver"
	"encoding/json"

	pModels "github.com/lonzzi/leetcode-daily-rank/pkg/models"
)

type SliceType interface {
	int | string
}

type Slices[T SliceType] []T
type Custom[T pModels.RecentACSubmission | pModels.UserCalendar] struct{}

func (s Slices[T]) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Slices[T]) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &s)
}

func (c Custom[T]) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Custom[T]) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &c)
}
