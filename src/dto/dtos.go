package dto

import (
	"time"
)

type TUserProfile struct {
	UserId      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	Birthday    time.Time `json:"birthday"`
	MailAddress string    `json:"mail_address"`
	Password    string    `json:"password"`
	Profile     string    `json:"profile"`
	RegistData  time.Time `json:"regist_data"`
	CreateUser  string    `json:"create_user"`
	CreateDate  time.Time `json:"create_date"`
	UpdateUser  string    `json:"update_user"`
	UpdateDate  time.Time `json:"update_date"`
}

func (m *TUserProfile) TableName() string {
	return "t_user_profile"
}
