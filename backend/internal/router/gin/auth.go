package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MangataL/BangumiBuddy/internal/auth"
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
	loginPath = "/login"
)

// CheckCookie checks the cookie in the request
func (a *Auth) CheckCookie(c *gin.Context) {
	cookie, err := a.authenticator.CheckCookie(c.Request)
	if strings.HasSuffix(c.Request.URL.Path, loginPath) {
		if err == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.Abort()
			return
		}
	} else {
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, loginPath)
			c.Abort()
			return
		}
	}

	c.Next()
	if cookie.Name == "" {
		return
	}
	// refresh cookie
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResp struct {
	Message string `json:"msg"`
}

func (a *Auth) Login(c *gin.Context) {
	var req loginReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	cookie, err := a.authenticator.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, loginResp{Message: err.Error()})
		return
	}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.JSON(http.StatusOK, loginResp{})
}

type logoutReq struct {
	Username string `json:"username"`
}

func (a *Auth) Logout(c *gin.Context) {
	var req logoutReq
	if err := c.BindJSON(&req); err != nil {
		return
	}
	if err := a.authenticator.Logout(c.Request.Context(), req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, loginResp{Message: err.Error()})
		return
	}
	c.SetCookie(auth.CookieName, "", -1, "", "", false, false)
	c.JSON(http.StatusOK, loginResp{})
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
	if err := a.authenticator.UpdateUser(c.Request.Context(), req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, loginResp{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginResp{})
}
