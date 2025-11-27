package controllers

import (
	"net/http"
	"strconv"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Service *services.RoleService
}

func NewRoleController(svc *services.RoleService) *RoleController {
	return &RoleController{Service: svc}
}

func (ctr *RoleController) Create(c *gin.Context) {
	var req entities.CreateRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := ctr.Service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (ctr *RoleController) GetAll(c *gin.Context) {
	roles, err := ctr.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (ctr *RoleController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := ctr.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (ctr *RoleController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req entities.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := ctr.Service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (ctr *RoleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ctr.Service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role berhasil dihapus"})
}
