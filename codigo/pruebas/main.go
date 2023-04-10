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

func compararTiempo(archivo1 string, archivo2 string) bool {
	out, err := exec.Command("stat", archivo1, archivo2).Output()
	if err != nil { fmt.Printf("%s", err) }

	output := string(out[:])

	fmt.Println(output)

	acceso1 := strings.Split(string(output), "\n")[7][8:27]
	acceso2 := strings.Split(string(output), "\n")[15][8:27]

	fmt.Println(acceso1)
	fmt.Println(acceso2)

	time1, error := time.Parse("2006-01-02 15:04:05", acceso1)
	if error != nil {
		fmt.Println(error)
	}
	time2, error1 := time.Parse("2006-01-02 15:04:05", acceso2)
	if error1 != nil {
		fmt.Println(error1)
	}

	res:= time1.Before(time2)

	return res
}

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

// -------------------------------------------------------- HACIA USB
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
					bitacora.ExisteEnBitacora(log)
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

// -------------------------------------------------------- DESDE USB
func getNombreArchivo(ruta string) string{
	division := strings.Split(ruta, "/")

	return division[len(division) - 1]
}

func comprararArchivos(ruta1 string, ruta2 string) bool{
	out, err := exec.Command("diff", "-sqr", ruta1, ruta2).Output()
	if err != nil { fmt.Printf("%s", err) }

	output := string(out[:])

	return strings.Contains(string(output), "identical")
}

func verificarCopiadosDesdeUSB(){
	// para los archivos en las usb
	for i:=0; i<len(archivosActuales); i++{
		nombreArchivo := getNombreArchivo(archivosActuales[i].Ruta);

		// buscar archivo en la computadora
		out, err := exec.Command("find", "/home", "-name", nombreArchivo).Output()
		if err != nil { fmt.Printf("%s", err) }

		output := string(out[:])

		// si encuentra los archivos en la computadora
		if(output != ""){
			archivosEnCompu := strings.Split(string(output), "\n");

			for j:=0; j<len(archivosEnCompu); j++{
				// verificar si son iguales
				if(comprararArchivos(archivosActuales[i].Ruta, archivosEnCompu[j])){
					fmt.Println("agregar a bitacora")
					if(compararTiempo(archivosActuales[i].Ruta, archivosEnCompu[j])){
						// comprobar que no exista ya en la bitacora
						fmt.Println("Se agrega a bitácora")
						log := bitacora.Log{
							Tipo: "Desde USB",
							Archivo: archivosEnCompu[j],
							FechaHora: time.Now().Format("2006-01-02 15:04:05")}
						bitacora.ExisteEnBitacora(log)
					}
				}
			}
		}
	}
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
		verificarCopiadosDesdeUSB();
		time.Sleep(3 * time.Second)
	}

}
