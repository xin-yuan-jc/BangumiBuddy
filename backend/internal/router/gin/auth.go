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

const (
	tokenPath = "/token"
)

func (a *Auth) CheckToken(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, tokenPath) {
		c.Next()
		return
	}
	if err := a.authenticator.CheckAccessToken(c.Request.Context(), getBearerToken(c.Request)); err != nil {
		code, msg := errs.ParseError(err)
		c.Data(code, "text/plain", []byte(msg))
		return
	}
	c.Next()
}

func getBearerToken(request *http.Request) string {
	return strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")
}

func (a *Auth) Login(c *gin.Context) {
	if grantType := c.Request.FormValue("grant_type"); grantType != "password" {
		c.JSON(http.StatusBadRequest, tokenError{
			Error:            "unsupported_grant_type",
			ErrorDescription: "不支持的授权类型",
		})
	}
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	credentials, err := a.authenticator.Authorize(c.Request.Context(), username, password)
	if err != nil {
		code, msg := errs.ParseError(err)
		errType := convertToErrorType(code)
		c.JSON(code, tokenError{Error: errType, ErrorDescription: msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  credentials.AccessToken,
		"token_type":    "Bearer",
		"refresh_token": credentials.RefreshToken,
	})
}

func convertToErrorType(code int) string {
	switch code {
	case http.StatusUnauthorized:
		return "invalid_request"
	case http.StatusForbidden:
		return "invalid_grant"
	default:
		return "invalid_scope"
	}
}

type tokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type updateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) UpdateUser(c *gin.Context) {
	var req updateReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	accessToken := getBearerToken(c.Request)
	if err := a.authenticator.UpdateUser(c.Request.Context(), accessToken, req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, loginResp{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginResp{})
}
