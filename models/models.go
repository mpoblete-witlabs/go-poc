package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Empleado struct {
	ID                int
	Nombre            string
	Cargo             string
	FechaContratacion Fecha `json:"fecha_contratacion" db:"fecha_contratacion"`
	EmpresaID         uint  `json:"empresa_id"`
}

type Empresa struct {
	gorm.Model
	Nombre            string
	EstaActiva        bool
	Direccion         string
	FechaContratacion Fecha      `json:"fecha_contratacion" db:"fecha_contratacion"`
	Empleados         []Empleado `gorm:"constraint:OnDelete:CASCADE;"`
}

type Fecha time.Time

var formatDate string = "2006-01-02"

func (f *Fecha) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(formatDate, s)

	*f = Fecha(nt)
	return
}

func (t Fecha) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(formatDate)), nil
}

func (f Fecha) MarshalJSON() ([]byte, error) {
	return []byte(f.String()), nil
}

func (f *Fecha) String() string {
	t := time.Time(*f)
	return fmt.Sprintf("%q", t.Format(formatDate))
}

func (f *Fecha) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*f = Fecha(v)
	case []byte:
		err := f.UnmarshalJSON(v)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("no se puede escanear el tipo %T", v)
	}
	return nil

}
