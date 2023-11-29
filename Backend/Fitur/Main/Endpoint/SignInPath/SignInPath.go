package signinpath

import (
	signindomains "Backend/Main/Features/SignIn/SignInDomains"
	signinpresentation "Backend/Main/Features/SignIn/SignInPresentation"

	"github.com/gin-gonic/gin"
)

func SignInPath(r *gin.Engine) {
	r.POST(signindomains.APIEndPointSignIn, signinpresentation.SignIn)
}
