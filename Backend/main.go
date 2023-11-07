package main

import (
	"github.com/gin-gonic/gin"

	signuppath "Backend/Main/EndPoint/SignUpPath"
	"Backend/models"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	models.ConnectDataBase()

	///***************************************** API Refactor **************************************************

	///================= SignIn =================
	// signinpath.SignInPath(r)

	///================= SignUp =================
	signuppath.SignUpPath(r)
}
