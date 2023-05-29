package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/DeniesLie/gpstracker/internal/api/controllers/response"
	"github.com/DeniesLie/gpstracker/internal/core/dto"
)

type trackRoutes struct {
	s TrackService
}

func AddTrackRoutes(handler *gin.Engine, s TrackService) {
	r := &trackRoutes{s}

	h := handler.Group("/tracks")
	{
		h.GET("", r.getAll)
		h.GET("/:id/info", r.getInfo)
		h.POST("", r.create)
		h.POST("/:id/complete", r.complete)
		h.DELETE("/:id", r.delete)
	}
}

func (r *trackRoutes) getAll(c *gin.Context) {
	tracks, err := r.s.GetAll()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.Success(tracks))
}

func (r *trackRoutes) getInfo(c *gin.Context) {
	trackIdParam := c.Param("id")
	trackId, convErr := strconv.ParseUint(trackIdParam, 10, 0)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, "trackId must have non-negative integer value")
		return
	}

	info, err := r.s.Getinfo(uint(trackId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.Success(info))
}

func (r *trackRoutes) create(c *gin.Context) {
	trackPost := dto.TrackPost{}
	if err := c.BindJSON(&trackPost); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	track, err := r.s.Create(trackPost)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, response.Success(track))
}

func (r *trackRoutes) complete(c *gin.Context) {
	trackIdParam := c.Param("id")
	trackId, convErr := strconv.ParseUint(trackIdParam, 10, 0)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, "trackId must have non-negative integer value")
		return
	}

	err := r.s.Complete(uint(trackId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.Success[any](nil))
}

func (r *trackRoutes) delete(c *gin.Context) {
	trackIdParam := c.Param("id")
	trackId, convErr := strconv.ParseUint(trackIdParam, 10, 0)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, "trackId must have non-negative integer value")
		return
	}

	err := r.s.Delete(uint(trackId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.Success[any](nil))
}
