package handlers

import (
	"net/http"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) PostLogin(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CheckUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	token, err := service.GenToken(user.Id, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponce{Error: "token generation error"})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"role":  user.Role,
	})

}

func (h *AuthHandler) PostRegister(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponce{Error: err.Error()})
		return
	}
	token, err := service.GenToken(user.Id, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponce{Error: "token generation error"})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"role":  user.Role,
	})
}

func (h *AuthHandler) GetCheckEmail(c *gin.Context) {
	email := c.Query("email")
	err := h.service.CheckEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponce{Error: err.Error()})
		return
	}
	c.Status(200)
}
func (h *AuthHandler) GetCheckPhone(c *gin.Context) {
	phone := c.Query("phone")
	err := h.service.CheckPhoneUnique(phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponce{Error: err.Error()})
		return
	}
	c.Status(200)
}
