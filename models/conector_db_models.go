package models

import (

	//"../config"
	// "github.com/gin-gonic/gin"

	"../config"
)

var (
	db  = config.DBInit()
	idb = config.InDB{DB: db}
	// idb = inDB

)
