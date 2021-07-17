package models

import (
	"../structs"
)

func GetUsers(params structs.ParamsGetListUser) ([]structs.DataGetUser, int, error) {

	var (
		data []structs.DataGetUser
		//t      structs.Component
		Count int
	)

	err := idb.DB.Table("users")
	err = err.Where("deleted_at is null")

	if params.Name != "" {
		err = err.Where("users.name like ?", "%"+params.Name+"%")
	}
	err = err.Count(&Count)
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Limit)
	}

	err = err.Find(&data)
	errx := err.Error

	return data, Count, errx
}

func CreateUsers(params structs.ParamsCreateUser) (structs.ParamsCreateUser, error) {
	var t = structs.Component{}
	params.CreatedAt = t.GetTimeNow()
	create := idb.DB.Table("users").Create(&params)
	err := create.Error
	return params, err
}

func UpdateUsers(params structs.ParamsUpdateUser) (structs.ParamsUpdateUser, error) {
	var t = structs.Component{}
	params.UpdatedAt = t.GetTimeNow()
	update := idb.DB.Table("users").Where("id = ?", params.ID).Updates(&params)
	err := update.Error
	return params, err
}

func DeleteUsers(params structs.ParamsUser) (structs.ParamsUser, error) {

	delete := idb.DB.Table("users").Where("id = ?", params.Id).Delete(&params)
	err := delete.Error
	return params, err
}

func Login(params structs.ParamsLogin) ([]structs.DataGetUser, error) {

	var (
		data []structs.DataGetUser
	)

	login := idb.DB.Table("users")
	login = login.Where("deleted_at is null")
	login = login.Where("users.username = ?", params.Username)
	login = login.Where("users.password = ?", params.Password)

	login = login.First(&data)

	errx := login.Error

	return data, errx
}
