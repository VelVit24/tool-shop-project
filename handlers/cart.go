package handlers

import (
	"net/http"
	"strconv"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) POSTCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	cart := models.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := repository.InsertCart(h.DB, id_user.(int), cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func (h *Handler) PUTCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	cart := models.Cart{}
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := repository.UpdateCart(h.DB, id_user.(int), cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}
func (h *Handler) DELETECart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = repository.DeleteCart(h.DB, id_user.(int), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(204)
}

func (h *Handler) GETCart(c *gin.Context) {
	id_user, ok := c.Get("user_id")
	if ok != true {
		c.JSON(http.StatusUnauthorized, "invalid token")
	}
	carts, err := repository.SelectCart(h.DB, id_user.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, carts)
}
