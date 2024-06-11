package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Estructura para los atributos
type Atributo struct {
	TipoDato string `json:"tipoDato"`
}

// Estructura para el tipo
type Tipo struct {
	Tipo      string              `json:"tipo"`
	Atributos map[string]Atributo `json:"atributos"`
}

// Función para decodificar el JSON
func decodificarJSON(jsonData []byte) ([]Tipo, error) {
	var tipos []Tipo
	err := json.Unmarshal(jsonData, &tipos)
	if err != nil {
		return nil, err
	}
	return tipos, nil
}

func main() {
	archivoJSON := "./inputs/classes.json"

	// Leer el contenido del archivo JSON
	jsonData, err := ioutil.ReadFile(archivoJSON)
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		os.Exit(1)
	}

	// Decodificar JSON en la estructura Tipo
	tipos, err := decodificarJSON(jsonData)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		os.Exit(1)
	}

	// PROVISIONAL [Solo 1 Tipo del JSON]
	mapAtributos := make(map[string]string)

	// Iterar sobre cada tipo y sus atributos
	for _, tipo := range tipos {
		fmt.Println("Clase:", tipo.Tipo)
		fmt.Println("Atributos:")
		for nombreAtributo, atributo := range tipo.Atributos {

			fmt.Printf(" - %s: %s\n", nombreAtributo, atributo.TipoDato)

			// PROVISIONAL [Solo 1 Tipo del JSON]
			mapAtributos[nombreAtributo] = atributo.TipoDato
		}

		// PROVISIONAL [Solo 1 Tipo del JSON]
		oneType := true
		if oneType == true {
			break
		}
	}

	// PROVISIONAL [Solo 1 Tipo del JSON]
	fmt.Println("mapAtributos es: ", mapAtributos)

}

func createFolderStructure() {

}

func createModels() {

}
