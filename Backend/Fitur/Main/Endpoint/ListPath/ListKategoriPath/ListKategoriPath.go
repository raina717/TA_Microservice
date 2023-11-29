package listkategoripath

import (
	listkategoridomains "Backend/Main/Features/List/ListKategori/ListKategoriDomains"
	listkategoripresentation "Backend/Main/Features/List/ListKategori/ListKategoriPresentation"

	"github.com/gin-gonic/gin"
)

func ListKategoriPath(r *gin.Engine) {
	r.GET(listkategoridomains.APIEndPointListKategori, listkategoripresentation.ListKategori)
}
