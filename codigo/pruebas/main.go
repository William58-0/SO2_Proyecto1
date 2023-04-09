package main

import (
    "fmt"
    "os/exec"
	"strings"
)

// sudo chmod 777 /media/
func main() {
	/*
	// una forma de habilitar 777 y desabilitar 000
	out, err := exec.Command("sudo", "chmod", "777", "/media/").Output()

    if err != nil {
        fmt.Printf("%s", err)
    }

    output := string(out[:])

	fmt.Println(output)
	*/

	// listar archivos en USB
	// ls /media/ -R
	archivosUSB, err := exec.Command("ls", "/media/", "-R").Output()

	if err != nil{
		return
	}

	conDivisionCarpeta := strings.Replace(string(archivosUSB), "\n\n", "\n<<DIVISION_CARPETA>>\n", -1)

	fmt.Println(string(conDivisionCarpeta))
	fmt.Println("\n")

	arregloArchivosUSB := strings.Split(string(conDivisionCarpeta), "\n")

	fmt.Println(arregloArchivosUSB)

	// recortar arreglo, porque las primeras no contienen archivos de USB
	recortado := arregloArchivosUSB[2:len(arregloArchivosUSB)]

	fmt.Println(recortado)

	// crear direcciones absolutas
	//var concatenar = false
	var carpeta = ""
	var direccionesAbs []string

	for i := 0; i < len(recortado); i++ {
		palabra := recortado[i]

		if palabra == "<<DIVISION_CARPETA>>"{
			// hay que concatenar
			// concatenar = true

			// el siguiente elemento del arreglo es la carpeta
			carpeta = strings.Replace(recortado[i+1], ":", "/",-1)
			i = i+1;
			fmt.Println("aqui tuvo que haber saltado", carpeta)
			continue
		}
	
		if(palabra != "" && strings.Count(carpeta, "/") >3){
			concatenado := carpeta + palabra
			direccionesAbs = append(direccionesAbs, concatenado)
		}
		
	}

	fmt.Println("CONCATENADO")
	fmt.Println(direccionesAbs)
	
	/*
	// sinSaltos := strings.Replace(string(archivosUSB), "\n\n", "\n", -1)
	sinSaltos1 := strings.Replace(string(archivosUSB), ":\n", "", -1)

	// listaArchivosUSB := strings.Split(string(archivosUSB), "\n\n")
	fmt.Println(sinSaltos1)
	*/

}