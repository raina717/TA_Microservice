package listkategoripresentation

import (
	"net/http"

	"github.com/gin-gonic/gin"

	listkategoridata "Backend/Main/Features/List/ListKategori/ListKategoriData"
	"Backend/core"
	"Backend/models"
)

func ListKategori(c *gin.Context) {
	var responseList []listkategoridata.RequestKategoriList

	/// Validate Request Bad Request
	// if err := c.ShouldBindJSON(&requestBody); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	query := `SELECT * FROM ms_kategori`
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
		var responseBody listkategoridata.RequestKategoriList
		err := rows.Scan(
			&responseBody.ID,
			&responseBody.Kategori,
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
