package signuppresentation

import (
	"net/http"

	"github.com/gin-gonic/gin"

	signupdata "Backend/Main/Features/SignUp/SignUpData"
	"Backend/models"
)

func SignUp(c *gin.Context) {
	var requestBody signupdata.UserSignUp

	/// Validate Request Bad Request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	/// Check Email And Password already use
	emailRequest := requestBody.Email
	query := `SELECT ID FROM user_data WHERE email = "` + emailRequest + `" `
	rows, err := models.DB.ConnPool.QueryContext(c, query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// for rows.Next() {
	// 	err := rows.Scan(
	// 		&tblUsers.Id,
	// 	)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// 	if tblUsers.Id != "" {
	// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Email already use"})
	// 		return
	// 	}
	// }

	var existingUserID string

	for rows.Next() {
		err := rows.Scan(&existingUserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// If existingUserID is not empty, it means email already exists
	if existingUserID != "" {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		return
	}

	// Insert user data into user_data table
	insertQuery := `
	INSERT INTO user_data (Full_Name, Email, Password, Phone_Number)
	VALUES (?, ?, ?, ?)
	`
	_, err = models.DB.ConnPool.ExecContext(c, insertQuery, requestBody.FullName, emailRequest, requestBody.Password, requestBody.PhoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})

}
