

package bitacora
 
import (
	"encoding/json"
	"io/ioutil"
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
 
	//_ = ioutil.WriteFile("bitacora.json", file, 0644)
    _ = ioutil.WriteFile("/tmp/bitacoraUSB.txt", file, 0644)
}

func LeerJSON() Logs{
	file, _ := ioutil.ReadFile("/tmp/bitacoraUSB.txt")
 
	data := Logs{}
 
	_ = json.Unmarshal([]byte(file), &data)

    return data
 
}

func AgregarBitacora(log Log){
    if(ExisteEnBitacora(log)){
        return
    }
    // fmt.Println("Agregando a bitacora")
    logs := LeerJSON();

    arregloLogs := logs.Bitacora;

    arregloLogs = append(arregloLogs, log);

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

    return existe
    
}

