package xgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tsingson/go-ums/model"
	"github.com/tsingson/go-ums/pkg/services"
)

// HTTPServer  define http server
type HTTPServer struct {
	as *services.AccountStore
}

// NewHTTPServer  initial http server
func NewHTTPServer() *HTTPServer {
	as := services.New()
	return &HTTPServer{
		as: as,
	}
}

// RegisterHandler register handler
func (hs *HTTPServer) RegisterHandler(c *gin.Context) {
	var err error

	var reqObj = model.AccountRequest{}

	if err = c.ShouldBindJSON(&reqObj); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var ac *model.Account
	ac, err = hs.as.Register(reqObj.Email, reqObj.Password)
	if err != nil {
		msg := error.Error(err)
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, ac)

}

// SetupRouter setup router
func (hs *HTTPServer) SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/register", hs.RegisterHandler)
	return r
}
