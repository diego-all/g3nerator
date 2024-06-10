package main

import (
	"encoding/json"
	"fmt"
	"g3nerator/models"
	"io/ioutil"
)

// func decodificarJSON(jsonData []byte) (map[string]interface{}, error) {
func decodificarJSON(jsonData []byte) (string, map[string]string, error) {
	//var productos []Product
	var productos []models.Product
	err := json.Unmarshal(jsonData, &productos)
	if err != nil {
		return "", nil, err
	}

	producto := productos[0] // Suponiendo que solo hay un producto en el JSON

	fmt.Println("La clase es: ", producto.Tipo)

	// Acceder a un map dentro de una interface

	typeInformation := make(map[string]string)

	for nombreAtributo, valor := range producto.Atributos {

		fmt.Printf("%s: %v \n", nombreAtributo, valor)

		fmt.Println("NombreAtributo es:", nombreAtributo)

		// nombresito := nombreAtributo
		// fmt.Println(nombresito)

		// Verificar si el tipo subyacente es un mapa
		if datos, ok := valor.(map[string]interface{}); ok {
			// Acceder a los valores dentro del mapa
			if tipoDato, ok := datos["tipoDato"].(string); ok {
				fmt.Println("Tipo de dato:", tipoDato)
				typeInformation[nombreAtributo] = tipoDato
			} else {
				fmt.Println("Error al acceder al tipo de dato")
			}

			// Manejar el valor según su tipo
			switch valor := datos["valor"].(type) {
			case string:
				fmt.Println("Valor (string):", valor)
			case float64:
				fmt.Println("Valor (float64):", valor)
			default:
				fmt.Println("Valor desconocido")
			}
		} else {
			fmt.Println("No se pudo convertir a un mapa")
		}

		// valorsito := valor
		// fmt.Println(valorsito)

		//atributosMap[nombreAtributo] = valor["tipoDato"]

		fmt.Println("\n")

	}

	// return producto.Tipo, nil, nil
	return producto.Tipo, typeInformation, nil
}

func main() {
	// ./ indica la carpeta actual (raiz del proyecto)
	archivoJSON := "./inputs/product.json"
	// Leer el contenido del archivo JSON
	jsonData, err := ioutil.ReadFile(archivoJSON)
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	// Decodificar JSON en la estructura Product
	clase := ""
	trin := make(map[string]string)

	clase, trin, err = decodificarJSON(jsonData)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
	} else {
		fmt.Println("Atributos del producto:")

	}

	fmt.Println("clase: ", clase)
	fmt.Println("map con atributos y tipo: ", trin)

}

func generateScaffolding() {

}

func createModels() {

}
