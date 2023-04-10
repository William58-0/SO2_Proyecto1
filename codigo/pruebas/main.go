package main

import (
    "fmt"
    "os/exec"
	"os"
	"strings"
	"io/ioutil"

	"USB/bitacora"

	"time"
)

type Archivo struct {
	USB string
	Ruta string
}

var archivosActuales []Archivo
var tienePermisos bool

func compararTiempo(archivo1 string, archivo2 string) (bool, string) {
	out, err := exec.Command("stat", archivo1, archivo2).Output()
	if err != nil { fmt.Printf("%s", err) }

	output := string(out[:])

	acceso1 := strings.Split(string(output), "\n")[7][8:27]
	acceso2 := strings.Split(string(output), "\n")[15][8:27]

	time1, error := time.Parse("2006-01-02 15:04:05", acceso1)
	if error != nil {
		fmt.Println(error)
	}
	time2, error1 := time.Parse("2006-01-02 15:04:05", acceso2)
	if error1 != nil {
		fmt.Println(error1)
	}

	var devolver string
	res:= time1.Before(time2)
	if(res){
		devolver = time2.Format("2006-01-02 15:04:05")
		// time.Now().Format("2006-01-02 15:04:05")
	}else{
		devolver = time1.Format("2006-01-02 15:04:05")
	}

	return res, devolver
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

// --------------------------------------------------------
func getNombreArchivo(ruta string) string{
	division := strings.Split(ruta, "/")

	return division[len(division) - 1]
}

func comprararArchivos(ruta1 string, ruta2 string) bool{
	out, err := exec.Command("diff", "-sqr", ruta1, ruta2).Output()
	if err != nil { 
		//fmt.Printf("%s", err) 
	}

	output := string(out[:])

	return strings.Contains(string(output), "identical")
}

func verificarArchivosCopiados(){
	if(!tienePermisos){
		return
	}
	archivosActuales = obtenerArchivosUSB();

	// para los archivos en las usb
	for i:=0; i<len(archivosActuales); i++{

		nombreArchivo := getNombreArchivo(archivosActuales[i].Ruta);

		// buscar archivo en la computadora
		out, err := exec.Command("find", "/home", "-name", nombreArchivo).Output()
		if err != nil { 
			// fmt.Printf("%s", err) 
		}

		output := string(out[:])

		// si encuentra los archivos en la computadora
		if(output != ""){
			archivosEnCompu := strings.Split(string(output), "\n");

			for j:=0; j<len(archivosEnCompu); j++{
				// ignorar los que estén en la papelera
				if(strings.Contains(archivosEnCompu[j], "share/Trash/")){
					continue
				}

				// verificar si son iguales
				if(comprararArchivos(archivosActuales[i].Ruta, archivosEnCompu[j])){
					menor, tiempo := compararTiempo(archivosActuales[i].Ruta, archivosEnCompu[j])
					if(menor){
						// desde usb
						log := bitacora.Log{
							Tipo: "Desde USB",
							Origen: archivosActuales[i].Ruta,
							Destino: archivosEnCompu[j],
							FechaHora: tiempo}
						bitacora.AgregarBitacora(log)
					} else {
						// hacia usb
						log := bitacora.Log{
							Tipo: "Hacia USB",
							Origen: archivosEnCompu[j],
							Destino: archivosActuales[i].Ruta,
							FechaHora: tiempo}
						bitacora.AgregarBitacora(log)
					}
				}
			}
		}
	}
}

func AnalizarArchivos(){
	if(!tienePermisos){
		return
	}
	archivosActuales = obtenerArchivosUSB();

	for {
		verificarArchivosCopiados()

		time.Sleep(2 * time.Second)
		
	}
}

func MostrarMenu(){
	// verificar si tiene permisos de usb
	_, err := ioutil.ReadDir("/media/")
    if err != nil {
		fmt.Println(err)
		tienePermisos = false
    } else {
		tienePermisos = true
	}

	mensajePermisos := ""
	quitarConceder := "Quitar"
	rutaBitacora := "/tmp/bitacoraUSB.txt"

	if (tienePermisos){
		mensajePermisos = "Los puertos USB están DESBLOQUEADOS"
		quitarConceder = "Bloquear"
	} else {
		mensajePermisos = "Los puertos USB están BLOQUEADOS"
		quitarConceder = "Desbloquear"
	}

	menu:= "\n---- William Alejandro Borrayo Alarcón - 201909103 ----\n\n"+
	mensajePermisos+"\n"+
	"Path de bitácora: " + rutaBitacora + "\n\n"+
	"Seleccione una opción:\n"+
	"1. "+quitarConceder+" puertos USB \n"+
	"2. Salir\n"


	for{
		// para analizar archivos copiados en segundo plano
		if(tienePermisos){
			go AnalizarArchivos()
		}
		
		fmt.Print(menu)

		// pedir numero
		var eleccion int
		fmt.Scanln(&eleccion)

		switch eleccion {
		case 1:
			if(Autenticar()){
				if(tienePermisos){
					_, err := exec.Command("sudo", "chmod", "000", "/media/").Output()
					if err != nil {
						fmt.Printf("%s", err)
					}
	
					tienePermisos = false
					fmt.Println("\n----- Se han bloqueado los puertos USB -----\n")		
				} else {
					_, err := exec.Command("sudo", "chmod", "777", "/media/").Output()
					if err != nil {
						fmt.Printf("%s", err)
					}
					tienePermisos = true
					fmt.Println("\n----- Se han desbloqueado los puertos USB -----\n")
				}
			}
			
			MostrarMenu()
		default:
			fmt.Println("\nAdiós!\n")
			os.Exit(0)
		}
	}
}

func Autenticar() bool{
	// pedir usuario y contraseña
	var usuario string
	var contrasenia string

	fmt.Println("\n-- Ingrese los datos de autenticación--\n")
	fmt.Println("Usuario:")
	fmt.Scanln(&usuario)

	fmt.Println("Contraseña:")
	fmt.Scanln(&contrasenia)

	if(usuario == "william" && contrasenia == "201909103"){
		return true
	}

	fmt.Println("\nUsuario y contraseña incorrectos\n")
	return false
}

// sudo chmod 777 /media/
func main() {
	for !Autenticar(){

	}
	MostrarMenu()
}
