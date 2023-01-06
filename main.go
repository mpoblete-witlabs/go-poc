package main

import (
	"poc-go/controller"
	"poc-go/database"
	"poc-go/models"
	"poc-go/repository"
	"poc-go/service"
	"poc-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.GetConfig()
	mysqlDb, _ := database.InitDatabase(config.MysqlConfig)

	mysqlDb.AutoMigrate(&models.Empresa{}, &models.Empleado{})

	empleadoRepository := repository.NewEmpleadoRepository(mysqlDb)
	empleadoService := service.NewEmpleadoService(empleadoRepository)
	empleadoController := controller.NewEmpleadoController(empleadoService)

	empresaRepository := repository.NewEmpresaRepository(mysqlDb)
	empresaService := service.NewEmpresaService(empresaRepository)
	empresaController := controller.NewEmpresaController(empresaService)

	httpRouter := gin.Default()

	httpRouter.POST("/empleados", empleadoController.AgregarEmpleado)
	httpRouter.PUT("/empleados/:ID", empleadoController.ActualizarEmpleado)

	httpRouter.POST("/empresas", empresaController.AgregarEmpresa)
	httpRouter.GET("/empresas", empresaController.ObtenerEmpresas)
	httpRouter.GET("/empresas/:name", empresaController.ObtenerEmpresa)
	httpRouter.DELETE("/empresas/:ID", empresaController.EliminarEmpresa)
	httpRouter.Run()

}
