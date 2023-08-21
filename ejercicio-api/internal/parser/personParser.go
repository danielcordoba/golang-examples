package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/CodelyTV/golang-examples/ejercicio-api/internal/api"
	"os"
	"reflect"
)

const fileName = "../../data/people.csv"

type PersonParser interface {
	SaveJson(person []api.Person) (err error)
}

func SaveJson(people []api.Person) (err error) {

	// Crear o abrir el archivo
	file, err := os.Create(fileName)

	if err != nil {
		return
	}

	// Crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado del CSV
	err = writeCSVHeader(writer, api.Person{})

	if err != nil {
		return
	}

	// Escribir los datos en el CSV
	for _, person := range people {
		err = writeCSVRow(writer, person)
		if err != nil {
			return
		}
	}

	if err == nil {
		fmt.Println("Archivo CSV guardado exitosamente.")
	}

	return
}

func writeCSVHeader(writer *csv.Writer, data interface{}) error {

	val := reflect.ValueOf(data)
	typeOfData := val.Type()

	headers := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		headers[i] = typeOfData.Field(i).Tag.Get("json")
	}

	return writer.Write(headers)
}

func writeCSVRow(writer *csv.Writer, data interface{}) error {
	val := reflect.ValueOf(data)

	row := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		row[i] = fmt.Sprint(field.Interface())
	}

	return writer.Write(row)
}
