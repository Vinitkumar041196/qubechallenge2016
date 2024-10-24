package server

import (
	"distributor-manager/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutDistributor godoc
//
//	@Summary		Create or Update Distributor
//	@Description	Can be used to update basic details and region permissions
//	@Accept			json
//	@Produce		json
//	@Param			request	body		types.Distributor	true	"distributor data"
//	@Success		200		{object}	types.SuccessResponse
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		500		{object}	types.ErrorResponse
//	@Router			/distributor [put]
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

// Get Distributor godoc
//
//	@Summary		Get Distributor
//	@Description	Get the latest details for distributor by code
//	@Produce		json
//	@Param			code	path		string	true	"distributor code"
//	@Success		200		{object}	types.Distributor
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		500		{object}	types.ErrorResponse
//	@Router			/distributor/{code} [get]
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

// Check Distributor Serviceable godoc
//
//	@Summary		Checks if distributor can provide service to a given region
//	@Description	Checks if distributor can provide service to a given region by matching own and ancestor's permissions
//	@Accept			json
//	@Produce		json
//	@Param			request	body		types.IsServiceableRequest	true	"distributor code and region"
//	@Success		200		{object}	types.IsServiceableResponse
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		500		{object}	types.ErrorResponse
//	@Router			/distributor/is_serviceable [post]
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

// Delete Distributor godoc
//
//	@Summary		Delete Distributor
//	@Description	Delete distributor by code
//	@Produce		json
//	@Param			code	path		string	true	"distributor code"
//	@Success		200		{object}	types.SuccessResponse
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		500		{object}	types.ErrorResponse
//	@Router			/distributor/{code} [delete]
func (srv *Server) DeleteDistributor(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "code cannot be empty"})
		return
	}

	err := srv.app.DeleteDistributor(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{Message: "success", Code: code})
}
