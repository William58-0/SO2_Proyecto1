package main

import (
    "fmt"
    "os/exec"
	"strings"

	"USB/bitacora"

	//"time"
)

type Archivo struct {
	USB string
	Ruta string
}

var archivosActuales []Archivo
var usbs []string

func obtenerArchivosUSB() []Archivo{
	// listar archivos en USB
	// ls /media/ -R
	archivosUSB, err := exec.Command("ls", "/media/", "-R").Output()
	if err != nil{	
	}

	conDivisionCarpeta := strings.Replace(string(archivosUSB), "\n\n", "\n<<DIVISION_CARPETA>>\n", -1)

	arregloArchivosUSB := strings.Split(string(conDivisionCarpeta), "\n")

	// listar usb's
	for u:= 4; u<len(arregloArchivosUSB); u++ {
		if(arregloArchivosUSB[u] == "<<DIVISION_CARPETA>>"){
			break
		}else{
			usbs = append(usbs, arregloArchivosUSB[u])
		}
	}

	fmt.Println(usbs)
	
	// recortar arreglo, porque las primeras no contienen archivos de USB
	recortado := arregloArchivosUSB[2:len(arregloArchivosUSB)]

	// crear direcciones absolutas
	var carpeta = ""
	var ListaArchivos []Archivo

	for i := 0; i < len(recortado); i++ {
		palabra := recortado[i]

		if palabra == "<<DIVISION_CARPETA>>"{
			// el siguiente elemento del arreglo es la carpeta
			carpeta = strings.Replace(recortado[i+1], ":", "/",-1)
			i = i+1;
			continue
		}
	
		if(palabra != "" && strings.Count(carpeta, "/") > 3){
			archivo := Archivo{USB: strings.Split(carpeta, "/")[3], Ruta: carpeta + palabra}
			ListaArchivos = append(ListaArchivos, archivo)
		}
		
	}

	//fmt.Println("CONCATENADO")
	//fmt.Println(ListaArchivos)

	return ListaArchivos
}

func verificarCopiadosHaciaUSB(){
	// nueva lectura de archivos
	nuevaLecturaArchivosUSB := obtenerArchivosUSB();

	
	fmt.Println("original")
	fmt.Println(archivosActuales)

	fmt.Println("nuevo")
	fmt.Println(nuevaLecturaArchivosUSB)
	

	if (len(nuevaLecturaArchivosUSB) > len(archivosActuales)){
		fmt.Println("si entra aquiiii")
		for i:=0; i<len(nuevaLecturaArchivosUSB); i++ {
			contiene := ArrayContains(archivosActuales, nuevaLecturaArchivosUSB[i])
			fmt.Println("contiene", contiene)
			// si el archivo nuevo de usb no está en los que leyó inicialmente
			if(!contiene){
				fmt.Println("entraaa qeui tambien")
				// si es de usb existente se agrega a bitacora, de lo contrario se omite
				if(ExisteUSB(nuevaLecturaArchivosUSB[i].USB)){
					fmt.Println("Se agrega a bitácora")
				}
			}
		}

	}

	
}

func ArrayContains(arreglo []Archivo, estruct Archivo) bool{
	// fmt.Println("COMAPRANDO")
	for i:=0; i<len(arreglo); i++{
		//fmt.Println(arreglo[i])
		//fmt.Println(estruct)
		if (arreglo[i] == estruct){
			//fmt.Println("Sí lo contiene")
			return true
		}
	}
	return false
}	

func ExisteUSB(nombre string) bool{
	fmt.Println("COMPROBANDO USBS")
	fmt.Println(usbs)
	fmt.Println(nombre)
	for i:=0; i<len(usbs); i++{
		//fmt.Println(usbs[i])
		//fmt.Println(nombre)
		if (usbs[i] == nombre){
			//fmt.Println("Sí existe")
			return true
		}
	}
	return false
}	

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

	/*
	archivosActuales = obtenerArchivosUSB();
	for {
		verificarCopiadosHaciaUSB();
		time.Sleep(1 * time.Second)
	}
	*/

	//bitacora.EscribirJSON();

	bitacora.LeerJSON();

}