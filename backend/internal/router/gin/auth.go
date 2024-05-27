package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MangataL/BangumiBuddy/internal/auth"
	"github.com/MangataL/BangumiBuddy/pkg/errs"
)

// Auth is the gin handler for authentication
type Auth struct {
	authenticator auth.Authenticator
}

// NewAuth creates a new Auth
func NewAuth(authenticator auth.Authenticator) *Auth {
	return &Auth{
		authenticator: authenticator,
	}
}

func (a *Auth) CheckToken(c *gin.Context) {
	if err := a.authenticator.CheckAccessToken(c.Request.Context(), getBearerToken(c.Request)); err != nil {
		code, msg := errs.ParseError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": msg})
		return
	}
	c.Next()
}

func getBearerToken(request *http.Request) string {
	return strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")
}

func (a *Auth) Token(c *gin.Context) {
	grantType := c.Request.FormValue("grant_type")
	switch grantType {
	case "password":
		a.authorize(c)
	case "refresh_token":
		a.refreshToken(c)
	default:
		c.JSON(http.StatusBadRequest, tokenError{
			Error:            "unsupported_response_type",
			ErrorDescription: "不支持的授权类型",
		})
	}
}

func (a *Auth) authorize(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	credentials, err := a.authenticator.Authorize(c.Request.Context(), username, password)
	if err != nil {
		writeOAuth2Error(c, err)
		return
	}
	writeCredentials(c, credentials)
}

func writeOAuth2Error(c *gin.Context, err error) {
	code, msg := errs.ParseError(err)
	errType := convertToOAuth2ErrorType(code)
	c.JSON(code, tokenError{Error: errType, ErrorDescription: msg})
}

func convertToOAuth2ErrorType(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "invalid_request"
	case http.StatusUnauthorized:
		return "invalid_grant"
	default:
		return "server_error"
	}
}

type tokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func writeCredentials(c *gin.Context, credentials auth.Credentials) {
	type data struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
	}
	c.JSON(http.StatusOK, data{
		AccessToken:  credentials.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: credentials.RefreshToken,
	})
}

func (a *Auth) refreshToken(c *gin.Context) {
	refreshToken := c.Request.FormValue("refresh_token")
	credentials, err := a.authenticator.RefreshCredentials(c.Request.Context(), refreshToken)
	if err != nil {
		writeOAuth2Error(c, err)
		return
	}
	writeCredentials(c, credentials)
}

func (a *Auth) UpdateUser(c *gin.Context) {
	type updateReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req updateReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	if err := a.authenticator.UpdateUser(c.Request.Context(), req.Username, req.Password); err != nil {
		writeError(c, err)
		return
	}
	c.Status(http.StatusOK)
}
