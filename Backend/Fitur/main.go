package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	signuppath "Backend/Main/EndPoint/SignUpPath"
	signinpath "Backend/Main/Endpoint/SignInPath"
	"Backend/models"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	models.ConnectDataBase()

	///***************************************** API Refactor **************************************************

	///================= SignIn =================
	signinpath.SignInPath(r)

	///================= SignUp =================
	signuppath.SignUpPath(r)

	// Define the port number
	port := 8080

	// Start the server
	// err := r.Run(":" + strconv.Itoa(port))
	// if err != nil {
	// 	log.Fatal("Failed to start the server: ", err)
	// } else {
	// 	log.Printf("Server is running on http://localhost:%d", port)
	// }

	go func() {
		if err := r.Run("localhost:" + strconv.Itoa(port)); err != nil {
			log.Fatal("Failed to start the server: ", err)
		}
	}()

	// Log a message indicating that the server is running
	log.Printf("Server is running on http://localhost:%d", port)

	// Block the main goroutine (e.g., with a select{}) to keep the program running
	select {}
}
