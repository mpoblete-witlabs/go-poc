package controller

import (
	"fmt"
	"net/http"
	"poc-go/models"
	"poc-go/service"

	"github.com/gin-gonic/gin"
)

type EmpleadoController interface {
	AgregarEmpleado(*gin.Context)
	ObtenerEmpleado(*gin.Context)
	ActualizarEmpleado(*gin.Context)
}

type empleadoController struct {
	empleadoService service.EmpleadoService
}

func NewEmpleadoController(s service.EmpleadoService) EmpleadoController {
	return empleadoController{
		empleadoService: s,
	}
}

func (u empleadoController) AgregarEmpleado(c *gin.Context) {
	var empleado models.Empleado

	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("---------->         :", empleado.EmpresaID)

	empl, err := u.empleadoService.Create(empleado)

	fmt.Println(empl)
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al crear registro",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": empl})

}

func (u empleadoController) ActualizarEmpleado(c *gin.Context) {
	id := c.Param("ID")
	var empl models.Empleado

	if err := c.ShouldBindJSON(&empl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empl, err := u.empleadoService.Update(id, empl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error alactualizar el registro",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": empl})

}

func (u empleadoController) ObtenerEmpleado(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)

	empl, err := u.empleadoService.GetOne(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": empl,
	})
}
