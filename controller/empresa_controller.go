package controller

import (
	"errors"
	"fmt"
	"net/http"
	"poc-go/models"
	"poc-go/service"
	"poc-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmpresaController interface {
	AgregarEmpresa(*gin.Context)
	ObtenerEmpresas(*gin.Context)
	ObtenerEmpresa(*gin.Context)
	EliminarEmpresa(*gin.Context)
	DummyValidatorFunc(*gin.Context)
	CrearEmpleadoEnEmpresa(*gin.Context)
	ListaEmpresasPaginable(c *gin.Context)
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
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": emprs})

}

func (e empresaController) ObtenerEmpresa(c *gin.Context) {
	name := c.Param("name")
	empr, err := e.empresaService.GetOne(name)

	if err != nil {
		fmt.Println("Error al consultar registro... = ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
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

func (e empresaController) DummyValidatorFunc(c *gin.Context) {

	type Dummy struct {
		Nombre     string `validate:"required"`
		Email      string `validate:"required,email"`
		EstaActiva bool   `validate:"required"`
		Direccion  string
		Age        int `validate:"required,gte=0,lte=100"`
	}

	var dumm Dummy

	if err := c.ShouldBindJSON(&dumm); err != nil {
		fmt.Println("Error al Bindear JSON")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(dumm); err != nil {
		fmt.Println("No pasa validaciones")
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			//fmt.Println("Len -> ", len(ve))
			out := make([]models.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = models.ApiError{
					Campo:   fe.Field(),
					Mensaje: utils.MsgForTag(fe),
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": out,
			})
		}
	}

}

func (e empresaController) CrearEmpleadoEnEmpresa(c *gin.Context) {
	var empl models.Empleado
	id := c.Param("ID")

	if err := c.ShouldBindJSON(&empl); err != nil {
		fmt.Println("Error al unmarshaling")
	}
	fmt.Println(empl)

	res, err := e.empresaService.CreateEmpleadoByEmpresa(id, empl)

	fmt.Println(res)
	fmt.Println(err)
}

func (e empresaController) ListaEmpresasPaginable(c *gin.Context) {
	pageParam := c.Param("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintln("No es posible formatear el valor: ", pageParam),
		})

		return
	}

	empresas, err := e.empresaService.ListEmpresasPagination(page)

	if err != nil {
		fmt.Println("Error al consultar lista de empresas")
	}

	c.JSON(http.StatusAccepted, gin.H{
		"empresas": empresas,
	})

}
