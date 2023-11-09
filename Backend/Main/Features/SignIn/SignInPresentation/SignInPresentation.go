package signinpresentation

import (
	"net/http"

	"github.com/gin-gonic/gin"

	signindata "Backend/Main/Features/SignIn/SignInData"
	"Backend/core"
	"Backend/models"
)

func SignIn(c *gin.Context) {
	var requestBody signindata.RequestUsersLogin
	var responseBody signindata.ResponseUsersLogin

	/// Validate Request Bad Request
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	/// Validate Response Bad Request
	if requestBody.Email == "" || requestBody.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	email := requestBody.Email
	password := requestBody.Password

	query := `SELECT 
	user.ID, 
	user.First_Name, 
	user.Last_Name, 
	user.Email,
	user.Phone_Number,
	user.Gender,
	user.Address,
	user.Status
	FROM user_data as user
	WHERE user.Email = "` + email + `" AND user.Password = "` + password + `" LIMIT 1`
	rows, err := models.DB.ConnPool.QueryContext(c, query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for rows.Next() {
		err := rows.Scan(
			&responseBody.Id,
			&responseBody.FirstName,
			&responseBody.LastName,
			&responseBody.Email,
			&responseBody.PhoneNumber,
			&responseBody.Gender,
			&responseBody.Address,
			&responseBody.Status,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	/// Validate Response Bad Request
	if responseBody.Id == "" || responseBody.FirstName == "" || responseBody.Email == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User doesn't exist"})
		return
	}

	// if responseBody.IdStatus != 1 {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": GetStatusName(c, responseBody.IdStatus, responseBody)})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		core.StatusParamCodeEndPoint:     http.StatusOK,
		core.StatusParamMessageEndPoint:  http.StatusText(http.StatusOK),
		core.StatusParamResponseEndPoint: responseBody})
}
