package main

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TablaPrueba struct {
	gorm.Model
	Texto       string
	PruebaFecha datatypes.Date
}

func mainDos() {
	//Conexion BD
	dsn := "development:devPassword1!@tcp(127.0.0.1:3306)/trafkindb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//Se crea la tabla
	db.AutoMigrate(&TablaPrueba{})

	//Probando inserts
	err = insertFunc(db, "hola", datatypes.Date(time.Now()))
	if err != nil {
		fmt.Println("Error al hacer primer insert: " + err.Error())
	}
	err = insertFunc(db, "mundo", datatypes.Date(time.Date(2023, 1, 10, 0, 0, 0, 0, time.Local)))
	if err != nil {
		fmt.Println("Error al hacer segundo insert: " + err.Error())
	}

	//Probando consulta
	var pruebaConsulta []TablaPrueba
	resultConsulta := db.Find(&pruebaConsulta)
	fmt.Println("Find RowsAffected: " + strconv.FormatInt(resultConsulta.RowsAffected, 10))
	if resultConsulta.Error != nil {
		fmt.Println("Error: " + resultConsulta.Error.Error())
	}

	//Imprimiendo registros insertados
	for i, registroPrueba := range pruebaConsulta {
		fmt.Println(registroPrueba)
		fmt.Print("ID[" + strconv.Itoa(i) + "]: " + strconv.Itoa(int(registroPrueba.ID)))
		fmt.Print(", Texto[" + strconv.Itoa(i) + "]: " + registroPrueba.Texto)

		fecha, _ := registroPrueba.PruebaFecha.Value()
		fmt.Println("Prueba Fecha[" + strconv.Itoa(i) + "]: " + fecha.(time.Time).Format("2006-01-02"))
	}
}

func insertFunc(db *gorm.DB, texto string, fecha datatypes.Date) error {
	insert := TablaPrueba{Texto: texto, PruebaFecha: fecha}
	resultInsert := db.Create(&insert)
	fmt.Println("Create RowsAffected: " + strconv.FormatInt(resultInsert.RowsAffected, 10))
	return resultInsert.Error
}
