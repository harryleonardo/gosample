package tech_curricullum

import (
	"time"
)

type ListData struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	MSISDN      int       `json:"msisdn"`
	Email       string    `json:"email"`
	BirthDate   time.Time `json:"birth_date"`
	CreatedTime time.Time `json:"created_time"`
	UpdateTime  time.Time `json:"update_time"`
	UserAge     int       `json:"user_age"`
}
