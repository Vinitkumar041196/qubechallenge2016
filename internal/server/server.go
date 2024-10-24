package server

import (
	"distributor-manager/internal/app"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app    *app.App
	router http.Handler
}

func NewServer(app *app.App) *Server {
	srv := &Server{app: app}
	srv.SetRouter()
	return srv
}

func (srv *Server) SetRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "success"}) })

	router.PUT("/distributor", srv.PutDistributor)
	router.GET("/distributor/:code", srv.GetDistributor)
	router.POST("/distributor/is_serviceable", srv.CheckIsServiceable)

	srv.router = router
}

func (srv *Server) Start() error {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = ":8080"
	}

	httpSrv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      srv.router,
	}

	return httpSrv.ListenAndServe()
}
