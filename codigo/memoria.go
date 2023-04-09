package main

import (
    "fmt"
    "os/exec"
	"strings"
	"strconv"
)

// Memoria struct
type Memoria struct {
	Total string `json:"Total"`
	Usado string `json:"Usado"`
	Porcentaje string `json:"Porcentaje"`
}

// https://www.educative.io/answers/how-to-execute-linux-commands-in-golang
func EjecutarComando(cmd string, args string) string{
    out, err := exec.Command(cmd, args).Output()

    if err != nil {
        fmt.Printf("%s", err)
    }

    output := string(out[:])

	return output
}


func QuitarEspacios(texto string) []string{
    arreglo := strings.Split(texto, " ");
	var nuevo []string
	
	for i := 0; i < len(arreglo); i++ {
		nueva := strings.Replace(arreglo[i], " ", "",-1)
		if(nueva != ""){
			nuevo = append(nuevo, nueva)
		}
		
	}

	return nuevo
}

func getMemoriaInfo() Memoria{
	factorGb := 1024 * 1024 * 1024;
	memoria := Memoria{}

	resultado := EjecutarComando("free", "-b")
	
	// quitar lineas innecesarias
    sinSaltos := strings.Split(resultado,"\n")[1]

	// total de memoria
	total := QuitarEspacios(sinSaltos)[1]
	total1, err := strconv.Atoi(total)
	if err != nil {
		return memoria
	}
	total2 := fmt.Sprintf("%.2f", (float64(total1) / float64(factorGb)))


	// memoria utilizada
	usado := QuitarEspacios(sinSaltos)[2]
	usado1, err := strconv.Atoi(usado)
	if err != nil {
		return memoria
	}
	usado2 := fmt.Sprintf("%.2f", (float64(usado1) / float64(factorGb)))


	// porcentaje
	porcentaje := (float64(usado1) / float64(total1)) * 100
	porcentaje1 := fmt.Sprintf("%.2f", porcentaje)

	memoria = Memoria{
		Total: total2,
		Usado: usado2,
		Porcentaje: porcentaje1,
	}

	return memoria
}