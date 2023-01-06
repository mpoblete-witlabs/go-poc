package repository

import (
	"fmt"
	"poc-go/models"

	"gorm.io/gorm"
)

type empleadoRepository struct {
	DB *gorm.DB
}

type EmpleadoRepository interface {
	Create(models.Empleado) (models.Empleado, error)
	GetAll() ([]models.Empleado, error)
	GetOne(string) (models.Empleado, error)
	Update(string, models.Empleado) (models.Empleado, error)
}

func NewEmpleadoRepository(db *gorm.DB) EmpleadoRepository {
	return empleadoRepository{
		DB: db,
	}
}

func (u empleadoRepository) Create(empleado models.Empleado) (models.Empleado, error) {
	err := u.DB.Create(&empleado).Error
	return empleado, err
}

func (u empleadoRepository) GetAll() (empleados []models.Empleado, err error) {
	err = u.DB.Find(&empleados).Error

	return empleados, err
}

func (u empleadoRepository) GetOne(s string) (models.Empleado, error) {
	var empl models.Empleado
	err := u.DB.Where("nombre = ?", s).First(&empl).Error

	if err != nil {
		fmt.Println("No encontrado : ", err)
	}

	return empl, err
}

func (u empleadoRepository) Update(s string, empl models.Empleado) (models.Empleado, error) {
	fmt.Println("EMPLEADO OBTENIDO--------->          :", empl)
	emplAux := empl
	err := u.DB.Where("ID = ?", s).First(&emplAux).Error

	fmt.Println("EMPLEADO OBTENIDO--------->          :", emplAux)

	if err != nil {
		fmt.Println("Error --->     :", err)
	}

	err = u.DB.Model(&emplAux).Updates(&empl).Error

	fmt.Println(emplAux, empl)

	return emplAux, err

}
