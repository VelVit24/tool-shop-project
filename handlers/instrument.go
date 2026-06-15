package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) POSTInstruments(c *gin.Context) {
	instr := models.Instrument{}
	if err := c.ShouldBindJSON(&instr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := repository.InsertInstrument(h.DB, &instr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *Handler) PUTInstruments(c *gin.Context) {
	instr := models.Instrument{}
	if err := c.ShouldBindJSON(&instr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	instr.Id = id
	err = repository.UpdateInstrument(h.DB, &instr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *Handler) DELETEInstruments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = repository.DeleteInstrument(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *Handler) GETInstruments(c *gin.Context) {
	cats, err := repository.SelectInstrument(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, cats)
}
