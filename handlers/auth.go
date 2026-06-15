package handlers

import (
	"net/http"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) POSTLogin(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := repository.CheckUser(h.DB, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	token, err := service.GenToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{"token": token})

}

func (h *Handler) POSTRegister(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := repository.InsertUser(h.DB, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	token, err := service.GenToken(id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, gin.H{"token": token})
}
