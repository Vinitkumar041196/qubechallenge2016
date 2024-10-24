package server

import (
	"distributor-manager/internal/app"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app    *app.App
	router http.Handler
}

// New API server
func NewServer(app *app.App) *Server {
	srv := &Server{app: app}
	srv.SetRouter()
	return srv
}

// Initializing API routes
func (srv *Server) SetRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "success"}) })

	router.PUT("/distributor", srv.PutDistributor)
	router.GET("/distributor/:code", srv.GetDistributor)
	router.POST("/distributor/is_serviceable", srv.CheckIsServiceable)

	srv.router = router
}

// Start HTTP Server
func (srv *Server) Start() error {
	httpSrv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      srv.router,
	}

	log.Println("Starting HTTP server on :8080")

	return httpSrv.ListenAndServe()
}
