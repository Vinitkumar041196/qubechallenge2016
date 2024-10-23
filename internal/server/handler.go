package server

import (
	"distributor-manager/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (srv *Server) PutDistributor(c *gin.Context) {
	req := new(types.Distributor)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err = srv.app.PutDistributor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "success", Code: req.Code})
}

func (srv *Server) GetDistributor(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "code cannot be empty"})
		return
	}

	dist, err := srv.app.GetDistributor(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dist)
}

func (srv *Server) DeleteDistributor(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "code cannot be empty"})
		return
	}

	err := srv.app.DeleteDistributor(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "success", Code: code})
}
