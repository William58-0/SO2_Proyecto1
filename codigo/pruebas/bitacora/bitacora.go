

package bitacora
 
import (
	"encoding/json"
	"io/ioutil"
    "fmt"
)

type Log struct {
    Tipo, Origen, Destino, FechaHora string
}

type Logs struct {
    Bitacora []Log
}
 
func EscribirJSON(arregloLogs []Log) {
	data := Logs{Bitacora: arregloLogs}
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func LeerJSON() Logs{
	file, _ := ioutil.ReadFile("test.json")
 
	data := Logs{}
 
	_ = json.Unmarshal([]byte(file), &data)
 
	for i := 0; i < len(data.Bitacora); i++ {
		fmt.Println("Tipo ", data.Bitacora[i].Tipo)
        fmt.Println("Origen ", data.Bitacora[i].Origen)
        fmt.Println("Destino ", data.Bitacora[i].Destino)
        fmt.Println("FechaHora ", data.Bitacora[i].FechaHora,"\n")
	}

    return data
 
}

func AgregarBitacora(log Log){
    if(ExisteEnBitacora(log)){
        return
    }
    fmt.Println("entraa")
    logs := LeerJSON();

    fmt.Println("logs iniciales")
    fmt.Println(logs)

    arregloLogs := logs.Bitacora;

    arregloLogs = append(arregloLogs, log);

    for i := 0; i < len(arregloLogs); i++ {
		fmt.Println("Tipo ", arregloLogs[i].Tipo)
        fmt.Println("Origen ", arregloLogs[i].Origen)
        fmt.Println("Destino ", arregloLogs[i].Destino)
        fmt.Println("FechaHora ", arregloLogs[i].FechaHora,"\n")
	}

    EscribirJSON(arregloLogs);
}

func ExisteEnBitacora(log Log) bool{
    logs := LeerJSON();

    arregloLogs := logs.Bitacora;

    existe := false

    for i := 0; i < len(arregloLogs); i++ {
        if(arregloLogs[i].Tipo == log.Tipo && arregloLogs[i].Destino == log.Destino ){
            existe = true;
            break;
        }
	} 

    fmt.Println("llega aquii")
    fmt.Println(existe)

    return existe
}

