package xgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tsingson/go-ums/model"
	"github.com/tsingson/go-ums/pkg/services"
)

type HttpServer struct {
	as *services.AccountStore
}

// var once sync.Once

func NewHttpServer() *HttpServer {
	as := services.New()
	return &HttpServer{
		as: as,
	}
}

func (hs *HttpServer) RegisterHandler(c *gin.Context) {
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

func (hs *HttpServer) SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/register", hs.RegisterHandler)
	return r
}
