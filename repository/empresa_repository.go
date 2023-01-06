package repository

import (
	"fmt"
	"poc-go/models"

	"gorm.io/gorm"
)

type empresaRepository struct {
	DB *gorm.DB
}

type EmpresaRepository interface {
	GetAll() ([]models.Empresa, error)
	Create(models.Empresa) (models.Empresa, error)
	GetOne(string) (models.Empresa, error)
	Delete(string) (models.Empresa, error)
}

func NewEmpresaRepository(db *gorm.DB) EmpresaRepository {
	return empresaRepository{
		DB: db,
	}
}

func (e empresaRepository) GetAll() (empresas []models.Empresa, err error) {
	//err = e.DB.Find(&empresas).Error
	err = e.DB.Model(&models.Empresa{}).Preload("Empleados").Find(&empresas).Error
	fmt.Println(empresas)

	return empresas, err
}

func (e empresaRepository) Create(empresa models.Empresa) (models.Empresa, error) {
	err := e.DB.Create(&empresa).Error
	return empresa, err
}

func (e empresaRepository) GetOne(s string) (models.Empresa, error) {
	var empr models.Empresa
	fmt.Println("String que llega por parámetro = ", s)
	err := e.DB.Where("nombre = ?", s).Preload("Empleados").First(&empr).Error

	return empr, err
}

func (e empresaRepository) Delete(s string) (models.Empresa, error) {
	var empr models.Empresa
	err := e.DB.Where("ID = ?", s).Preload("Empleados").First(&empr).Error
	if err != nil {
		fmt.Println(err)
	}

	//err = e.DB.Unscoped().Select("Empleados").Delete(&empr).Error //Este metodo elimina el registro y registros asociados en la tabla Empleados
	err = e.DB.Select("Empleados").Delete(&empr).Error //ESTE METODO APLICA SOFT DELETED (actualiza el campo deleted_at no elimina el registro) EN LA TABLA PRINCIPAL Y ELIMINA LOS REGISTROS "HIJOS".

	return empr, err
}
