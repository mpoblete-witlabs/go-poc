package service

import (
	"poc-go/models"
	"poc-go/repository"
)

type EmpleadoService interface {
	Create(models.Empleado) (models.Empleado, error)
	GetAll() ([]models.Empleado, error)
	Update(string, models.Empleado) (models.Empleado, error)
	GetOne(string) (models.Empleado, error)
}

type empleadoService struct {
	empleadoRepository repository.EmpleadoRepository
}

func NewEmpleadoService(r repository.EmpleadoRepository) EmpleadoService {
	return empleadoService{
		empleadoRepository: r,
	}
}

func (u empleadoService) Create(empleado models.Empleado) (models.Empleado, error) {
	return u.empleadoRepository.Create(empleado)
}

func (u empleadoService) GetAll() ([]models.Empleado, error) {
	return u.empleadoRepository.GetAll()
}

func (u empleadoService) Update(s string, empl models.Empleado) (models.Empleado, error) {
	return u.empleadoRepository.Update(s, empl)
}

func (u empleadoService) GetOne(s string) (models.Empleado, error) {
	return u.empleadoRepository.GetOne(s)
}
