package models

// with https://transform.tools/json-to-go
// type Product []struct {
// 	Tipo      string `json:"tipo"`
// 	Atributos struct {
// 		Nombre struct {
// 			Valor    string `json:"valor"`
// 			TipoDato string `json:"tipoDato"`
// 		} `json:"nombre"`
// 		Descripcion struct {
// 			Valor    string `json:"valor"`
// 			TipoDato string `json:"tipoDato"`
// 		} `json:"descripcion"`
// 		Precio struct {
// 			Valor    int    `json:"valor"`
// 			TipoDato string `json:"tipoDato"`
// 		} `json:"precio"`
// 	} `json:"atributos"`
// }

// type Product struct {
// 	Tipo      string                 `json:"tipo"`
// 	Atributos map[string]interface{} `json:"atributos"`
// }

// with https://json2struct.mervine.net/  (BOTA EL TIPO)
type MyJsonName []struct {
	Atributos struct {
		Descripcion struct {
			TipoDato string `json:"tipoDato"`
			Valor    string `json:"valor"`
		} `json:"descripcion"`
		Nombre struct {
			TipoDato string `json:"tipoDato"`
			Valor    string `json:"valor"`
		} `json:"nombre"`
		Precio struct {
			TipoDato string `json:"tipoDato"`
			Valor    int64  `json:"valor"`
		} `json:"precio"`
	} `json:"atributos"`
	Tipo string `json:"tipo"`
}

type Product struct {
	Tipo      string                 `json:"tipo"`
	Atributos map[string]interface{} `json:"atributos"`
}

// Definición del struct Atributo
type Atributo struct {
	Valor    interface{} `json:"valor"`
	TipoDato string      `json:"tipoDato"`
}
