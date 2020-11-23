package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

//Aux ...
type Aux struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

//Server ...
type Server struct {
}

//Materias ...
var Materias = make(map[string]map[string]float64)

//Alumno ...
var Alumno = make(map[string]map[string]float64)

//AgregarCalificacion ...
func (*Server) AgregarCalificacion(alumno Aux, reply *string) error {
	_, flagMateria := Materias[alumno.Materia]               //Bool si la materia existe
	_, flagAlumno := Materias[alumno.Materia][alumno.Alumno] //bool si el alumno tiene calificación en la meteria
	_, flagAlumnoClases := Alumno[alumno.Alumno]             //si el alumno existe
	if flagMateria {
		if flagAlumno {
			return errors.New("El alumno ya tiene calificación para la materia")
		}
		Materias[alumno.Materia][alumno.Alumno] = alumno.Calificacion
		if flagAlumnoClases {
			Alumno[alumno.Alumno][alumno.Materia] = alumno.Calificacion
		} else {
			var materiaAux = make(map[string]float64)
			materiaAux[alumno.Materia] = alumno.Calificacion
			Alumno[alumno.Alumno] = materiaAux
		}
		return nil
	}
	//Crear materia
	var alumnos = make(map[string]float64)
	alumnos[alumno.Alumno] = alumno.Calificacion
	Materias[alumno.Materia] = alumnos
	if flagAlumnoClases {
		Alumno[alumno.Alumno][alumno.Materia] = alumno.Calificacion
	} else {
		var materiaAux = make(map[string]float64)
		materiaAux[alumno.Materia] = alumno.Calificacion
		Alumno[alumno.Alumno] = materiaAux
	}
	return nil
}

//ObtenerPromedioAlumno ...
func (*Server) ObtenerPromedioAlumno(nombre string, reply *float64) error {
	var promedio float64
	promedio = 0
	numMaterias := 0.0
	_, flagAlumno := Alumno[nombre]
	if flagAlumno {
		for _, calificación := range Alumno[nombre] {
			promedio += calificación
			numMaterias++
		}
		promedio = promedio / numMaterias
		*reply = promedio
		return nil
	}
	return errors.New("El alumno no existe")
}

//PromedioGeneral ...
func (*Server) PromedioGeneral(_ int, reply *float64) error {
	alumnos := len(Alumno)
	var PromedioGeneral float64
	var numMaterias float64
	if alumnos > 0 {
		for materia := range Materias {
			numMaterias++
			promedioMateria := 0.0
			numAlumnos := 0.0
			for _, cal := range Materias[materia] {
				promedioMateria += cal
				numAlumnos++
			}
			PromedioGeneral += promedioMateria / numAlumnos
		}
		*reply = PromedioGeneral / numMaterias
		return nil
	}
	return errors.New("Ningun alumno registrado")
}

// PromedioMateria ...
func (*Server) PromedioMateria(materia string, reply *float64) error {
	_, ok := Materias[materia]
	if ok {
		promedio := 0.0
		numAlumnos := 0.0
		for _, cal := range Materias[materia] {
			promedio += cal
			numAlumnos++
		}
		*reply = promedio / numAlumnos
		return nil
	}
	return errors.New("No existe la materia")
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}
