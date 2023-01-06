package controller

import (
	"fmt"
	"net/http"
	"poc-go/models"
	"poc-go/service"

	"github.com/gin-gonic/gin"
)

type EmpresaController interface {
	AgregarEmpresa(*gin.Context)
	ObtenerEmpresas(*gin.Context)
	ObtenerEmpresa(*gin.Context)
	EliminarEmpresa(*gin.Context)
}

type empresaController struct {
	empresaService service.EmpresaService
}

func NewEmpresaController(s service.EmpresaService) EmpresaController {
	return empresaController{
		empresaService: s,
	}
}

func (e empresaController) AgregarEmpresa(c *gin.Context) {

	var empresa models.Empresa

	if err := c.ShouldBindJSON(&empresa); err != nil {
		fmt.Println("Error al parsear Request... = {}", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintln(err.Error()),
		})
	}

	empr, err := e.empresaService.Create(empresa)

	if err != nil {
		fmt.Println("Error al guardar Registro... = {}", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": empr})

}

func (e empresaController) ObtenerEmpresas(c *gin.Context) {

	emprs, err := e.empresaService.GetAll()

	if err != nil {
		fmt.Println("Error al consultar registro... = {}", err)
		c.JSON(http.StatusBadRequest, gin.H{})

	}
	c.JSON(http.StatusOK, gin.H{"data": emprs})

}

func (e empresaController) ObtenerEmpresa(c *gin.Context) {
	name := c.Param("name")
	empr, err := e.empresaService.GetOne(name)

	if err != nil {
		fmt.Println("Error al consultar registro... = ", err)
	}

	fmt.Println(empr)
	c.JSON(http.StatusOK, gin.H{"data": empr})
}

// Metodo controlador que invoca al servicio que elimina un registro de la entidad
// 'Empresa'.
func (e empresaController) EliminarEmpresa(c *gin.Context) {
	id := c.Param("ID")

	empr, err := e.empresaService.Delete(id)

	if err != nil {
		fmt.Printf("Error al eliminar registro... %s Error ={ %s }", id, err)
	}
	fmt.Println(empr)
	c.JSON(http.StatusAccepted, gin.H{"msg": "Registro Eliminado", "data": empr})
}
