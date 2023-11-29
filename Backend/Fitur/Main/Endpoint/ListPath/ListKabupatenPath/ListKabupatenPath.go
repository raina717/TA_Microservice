package listkabupatenpath

import (
	listkabupatendomains "Backend/Main/Features/List/ListKabupaten/ListKabupatenDomains"
	listkabupatenpresentation "Backend/Main/Features/List/ListKabupaten/ListKabupatenPresentation"

	"github.com/gin-gonic/gin"
)

func ListKabupatenPath(r *gin.Engine) {
	r.GET(listkabupatendomains.APIEndPointListKabupaten, listkabupatenpresentation.ListKabupaten)
}
