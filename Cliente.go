package main

import (
	"fmt"
	"net/rpc"
)

//Aux ...
type Aux struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar Calificación a alumno")
		fmt.Println("2) Obtener el promedio de un alumno")
		fmt.Println("3) Obtner el promedio general de las materias")
		fmt.Println("4) Obtner el promedio la materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {

		case 1:
			var aux Aux
			fmt.Print("Nombre Alumno: ")
			fmt.Scanln(&aux.Alumno)
			fmt.Println("Nombre Materia:")
			fmt.Scanln(&aux.Materia)
			fmt.Println("Calificación:")
			fmt.Scanln(&aux.Calificacion)
			var result string
			err = c.Call("Server.AgregarCalificacion", aux, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Agregado con exito")
			}
		case 2:
			var nombre string
			fmt.Print("Nombre Alumno: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Server.ObtenerPromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("promedio:", result)
			}
		case 3:
			var ignore int
			var result float64
			err = c.Call("Server.PromedioGeneral", ignore, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio Genral=", result)
			}
		case 4:
			var nomMateria string
			var result float64
			fmt.Println("Nombre de la materia: ")
			fmt.Scanln(&nomMateria)
			err = c.Call("Server.PromedioMateria", nomMateria, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de ", nomMateria, " :", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
