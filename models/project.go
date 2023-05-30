package models

import "time"

type Project struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	GitType   string     `json:"git_type"`
	GitLink   string     `json:"git_link"`
	FootID    uint       `json:"foot_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
