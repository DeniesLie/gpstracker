package controllers

import (
	"net/http"
	"strconv"

	"github.com/DeniesLie/gpstracker/internal/api/controllers/response"
	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/gin-gonic/gin"
)

type waypointRoutes struct {
	s WaypointService
}

func AddWaypointRoutes(handler *gin.Engine, s WaypointService) {
	r := &waypointRoutes{s}

	h := handler.Group("/waypoints")
	{
		h.GET("/:trackId", r.getByTrackId)
		h.POST("/addBatch", r.addBatch)
	}
}

func (r *waypointRoutes) getByTrackId(c *gin.Context) {
	trackIdParam := c.Param("trackId")
	trackId, convErr := strconv.ParseUint(trackIdParam, 10, 0)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, "trackId must have non-negative integer value")
		return
	}

	waypoints, err := r.s.GetByTrackId(uint(trackId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.Success(waypoints))
}

func (r *waypointRoutes) addBatch(c *gin.Context) {
	batch := []dto.WaypointPost{}
	if err := c.BindJSON(&batch); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.s.AddBatch(batch)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, response.Success[any](nil))
}
