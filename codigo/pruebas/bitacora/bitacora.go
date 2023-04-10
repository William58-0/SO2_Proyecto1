

package bitacora
 
import (
	"encoding/json"
	"io/ioutil"
    "fmt"
)

type Log struct {
    Tipo, Archivo, FechaHora string
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
        fmt.Println("Archivo ", data.Bitacora[i].Archivo)
        fmt.Println("FechaHora ", data.Bitacora[i].FechaHora,"\n")
	}

    return data
 
}

func AgregarBitacora(log Log){
    logs := LeerJSON();

    fmt.Println("logs iniciales")
    fmt.Println(logs)

    arregloLogs := logs.Bitacora;

    arregloLogs = append(arregloLogs, log);

    for i := 0; i < len(arregloLogs); i++ {
		fmt.Println("Tipo ", arregloLogs[i].Tipo)
        fmt.Println("Archivo ", arregloLogs[i].Archivo)
        fmt.Println("FechaHora ", arregloLogs[i].FechaHora,"\n")
	}

    EscribirJSON(arregloLogs);
}

