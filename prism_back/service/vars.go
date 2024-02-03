package service

import (
	"os"
	"strconv"
)

var (
	adminSession      string = os.Getenv("ADMIN_SESSION")
	userSession       string = os.Getenv("USER_SESSION")
	rootAdminID       string = os.Getenv("ROOT_ADMIN_ID")
	rootAdminPassword string = os.Getenv("ROOT_ADMIN_PASSWORD")
	profileFolder = os.Getenv("PROFILE_IMAGE_FOLDER")
	adminRank, err = strconv.Atoi(os.Getenv("ADMIN_RANK"))
)
