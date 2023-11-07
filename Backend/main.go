package main

import (
	"strconv"

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

	// Define the port number
	port := 8080

	// Start the server
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}
