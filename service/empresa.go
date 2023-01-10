package service

import (
	"fmt"
	"poc-go/models"
	"poc-go/repository"
	"strings"
)

type EmpresaService interface {
	GetAll() ([]models.Empresa, error)
	Create(models.Empresa) (models.Empresa, error)
	GetOne(string) (models.Empresa, error)
	Delete(string) (models.Empresa, error)
	CreateEmpleadoByEmpresa(string, models.Empleado) (models.Empresa, error)
	ListEmpresasPagination(int) ([]models.Empresa, error)
}

type empresaService struct {
	empresaRepository repository.EmpresaRepository
}

func NewEmpresaService(r repository.EmpresaRepository) EmpresaService {
	return empresaService{
		empresaRepository: r,
	}
}

func (u empresaService) Create(empresa models.Empresa) (models.Empresa, error) {
	return u.empresaRepository.Create(empresa)
}

func (u empresaService) GetAll() ([]models.Empresa, error) {
	return u.empresaRepository.GetAll()
}

func (u empresaService) GetOne(s string) (models.Empresa, error) {
	fmt.Println(s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	fmt.Println(s)
	return u.empresaRepository.GetOne(s)
}

func (u empresaService) Delete(s string) (models.Empresa, error) {
	return u.empresaRepository.Delete(s)
}

func (u empresaService) CreateEmpleadoByEmpresa(s string, empl models.Empleado) (models.Empresa, error) {
	return u.empresaRepository.CreateEmpleadoByEmpresa(s, empl)

}

func (u empresaService) ListEmpresasPagination(page int) ([]models.Empresa, error) {
	return u.empresaRepository.ListEmpresasPageable(page)
}
