package main

import (
    "fmt"
    "os/exec"
	"strings"

	"USB/bitacora"

	"time"
)

type Archivo struct {
	USB string
	Ruta string
}

var archivosActuales []Archivo
var usbsActuales []string

func obtenerArchivosUSB() []Archivo{
	// listar archivos en USB
	// ls /media/ -R
	archivosUSB, err := exec.Command("ls", "/media/", "-R").Output()
	if err != nil{	
	}

	conDivisionCarpeta := strings.Replace(string(archivosUSB), "\n\n", "\n<<DIVISION_CARPETA>>\n", -1)

	arregloArchivosUSB := strings.Split(string(conDivisionCarpeta), "\n")
	
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

	return ListaArchivos
}

func ListarUSBS() []string{
	var usbs []string
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
		} else {
			usbs = append(usbs, arregloArchivosUSB[u])
		}
	}

	return usbs
}

func verificarCopiadosHaciaUSB(){
	// nueva lectura de archivos
	nuevaLecturaArchivosUSB := obtenerArchivosUSB();
	
	if (len(nuevaLecturaArchivosUSB) > len(archivosActuales)){
		for i:=0; i<len(nuevaLecturaArchivosUSB); i++ {
			contiene := USBContains(archivosActuales, nuevaLecturaArchivosUSB[i])

			// si el archivo nuevo de usb no está en los que leyó inicialmente
			if(!contiene){
				// si es de usb existente se agrega a bitacora, de lo contrario se omite
				if(ExisteUSB(usbsActuales, nuevaLecturaArchivosUSB[i].USB)){
					fmt.Println("Se agrega a bitácora")
					log := bitacora.Log{
						Tipo: "Hacia USB",
						Archivo: nuevaLecturaArchivosUSB[i].Ruta,
						FechaHora: time.Now().Format("2006-01-02 15:04:05")}
					bitacora.AgregarBitacora(log)
				}
			}
		}
	}

	archivosActuales = nuevaLecturaArchivosUSB
}

func USBContains(arreglo []Archivo, estruct Archivo) bool{
	for i:=0; i<len(arreglo); i++{
		if (arreglo[i] == estruct){
			return true
		}
	}
	return false
}	

func ExisteUSB(usbs []string, nombre string) bool{
	for i:=0; i<len(usbs); i++{
		if (usbs[i] == nombre){
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


	archivosActuales = obtenerArchivosUSB();
	usbsActuales = ListarUSBS();

	for {
		verificarCopiadosHaciaUSB();
		time.Sleep(1 * time.Second)
	}


	// bitacora.EscribirJSON();

	// bitacora.LeerJSON();

}