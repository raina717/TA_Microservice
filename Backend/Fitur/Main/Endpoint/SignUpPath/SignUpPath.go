package signuppath

import (
	signupdomains "Backend/Main/Features/SignUp/SignUpDomains"
	signuppresentation "Backend/Main/Features/SignUp/SignUpPresentation"

	"github.com/gin-gonic/gin"
)

func SignUpPath(r *gin.Engine) {
	r.POST(signupdomains.APIEndPointSignUp, signuppresentation.SignUp)
}
