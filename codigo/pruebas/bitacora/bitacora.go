

package bitacora
 
import (
	"encoding/json"
	"io/ioutil"
    "fmt"
)

type Log struct {
    Tipo, Nombre, FechaHora string
}

type Logs struct {
    Bitacora []Log
}
 
func EscribirJSON() {
    log := Log{Tipo: "Hacia USB", Nombre:"archivo1", FechaHora: "hoy"}
    var logs []Log

    logs = append(logs, log)
    logs = append(logs, log)
    logs = append(logs, log)

	data := Logs{Bitacora: logs}
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func LeerJSON() {
	file, _ := ioutil.ReadFile("test.json")
 
	data := Logs{}
 
	_ = json.Unmarshal([]byte(file), &data)
 
	for i := 0; i < len(data.Bitacora); i++ {
		fmt.Println("Tipo ", data.Bitacora[i].Tipo)
        fmt.Println("Nombre ", data.Bitacora[i].Nombre)
        fmt.Println("FechaHora ", data.Bitacora[i].FechaHora,"\n")
	}
 
}

