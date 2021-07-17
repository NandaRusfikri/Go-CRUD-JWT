package structs

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name          string
	Phone         string
	Email         string
	EmailVerified bool
	Username      string
	Password      string
	IdNotif       string
	Gender        bool
	Address       string
}

type ParamsCreateUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name" binding:"required"`
	Phone     string `json:"phone" form:"phone" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Gender    bool   `json:"gender" form:"gender"`
	Address   string `json:"address" form:"address"`
	CreatedAt string
}

type ParamsUpdateUser struct {
	ID            int    `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" `
	Phone         string `json:"phone" form:"phone" `
	Email         string `json:"email" form:"email"`
	EmailVerified bool   `json:"email_verified" form:"email_verified"`
	Username      string `json:"username" form:"username" `
	Password      string `json:"password" form:"password" `
	Gender        bool   `json:"gender" form:"gender" `
	IdNotif       string `json:"id_notif" form:"id_notif"`
	Address       string `json:"address" form:"address"`
	UpdatedAt     string
}

type ParamsGetListUser struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Limit  *int   `json:"limit" form:"limit"`
	Offset *int   `json:"offset" form:"offset"`
}
type DataGetUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Gender   bool   `json:"gender"`
}
type ParamsUser struct {
	Id        int `json:"id" form:"id" binding:"required"`
	DeletedAt string
}
type DataDetailUser struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    bool   `json:"gender"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}
type ParamsLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
