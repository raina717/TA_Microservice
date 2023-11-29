package listkabupatenpresentation

import (
	"net/http"

	"github.com/gin-gonic/gin"

	listkabupatendata "Backend/Main/Features/List/ListKabupaten/ListKabupatenData"
	"Backend/core"
	"Backend/models"
)

func ListKabupaten(c *gin.Context) {
	var responseList []listkabupatendata.RequestKabupatenList

	/// Validate Request Bad Request
	// if err := c.ShouldBindJSON(&requestBody); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	query := `SELECT * FROM ms_kabupaten`
	rows, err := models.DB.ConnPool.QueryContext(c, query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// for rows.Next() {
	// 	err := rows.Scan(
	// 		&responseBody.ID,
	// 		&responseBody.Kabupaten,
	// 	)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }

	defer rows.Close()

	for rows.Next() {
		var responseBody listkabupatendata.RequestKabupatenList
		err := rows.Scan(
			&responseBody.ID,
			&responseBody.Kabupaten,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		// Append the current row to the responseList slice
		responseList = append(responseList, responseBody)
	}

	c.JSON(http.StatusOK, gin.H{
		core.StatusParamCodeEndPoint:     http.StatusOK,
		core.StatusParamMessageEndPoint:  http.StatusText(http.StatusOK),
		core.StatusParamResponseEndPoint: responseList})
}
