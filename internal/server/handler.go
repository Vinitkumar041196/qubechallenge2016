package server

import (
	"distributor-manager/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Server) PutDistributor(c *gin.Context) {
	req := new(types.Distributor)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	err = srv.app.PutDistributor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{Message: "success", Code: req.Code})
}

func (srv *Server) GetDistributor(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "code cannot be empty"})
		return
	}

	dist, err := srv.app.GetDistributor(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dist)
}

func (srv *Server) CheckIsServiceable(c *gin.Context) {
	req := new(types.IsServiceableRequest)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	isServiceable, err := srv.app.CheckIsServiceable(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}
	isServiceableStr := "NO"
	if isServiceable {
		isServiceableStr = "YES"
	}
	c.JSON(http.StatusOK, types.IsServiceableResponse{Code: req.Code, Region: req.Region, IsServiceable: isServiceableStr})
}
