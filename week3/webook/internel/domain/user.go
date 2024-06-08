package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	Nickname string `json:"nickname"`
	Birthday string `json:"birthday"`
	AboutMe  string `json:"aboutMe"`

	Ctime time.Time
}
