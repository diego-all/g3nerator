# JSON Case Design

~~~
class: Product
attributes: [
	"name": strint,
	"description": string,
	"price" int
]
~~~

Enfoque más flexible que permita incluir diferentes tipos de datos y estructuras.

[
  {
    "Clase": "Product",
    "atributos": {
      "nombre": "Camiseta",
      "descripcion": "Camiseta roja de algodón",
      "precio": 25000
    }
  },
  {
    "Clase": "Category",
    "atributos": {
      "propiedad1": "valor1",
      "propiedad2": "valor2"
    }
  }
]

Se utiliza un array JSON para almacenar múltiples objetos.
Cada objeto dentro del array representa una instancia de una clase.
Se incluye una propiedad "tipo" para identificar el tipo de clase a la que pertenece la instancia.
La propiedad "atributos" contiene un objeto que almacena los valores específicos de los atributos de la clase.
La estructura de los atributos dependerá del tipo de clase que se represente.


Beneficios de esta estructura:

Flexibilidad: Permite almacenar instancias de diferentes clases con sus respectivos atributos.
Genericidad: No está limitado a la clase Product, se puede adaptar a cualquier modelo de datos.
Escalabilidad: Se puede ampliar fácilmente para incluir nuevas clases y atributos.


~~~~
[
  {
    "tipo": "Product",
    "atributos": {
      "nombre": {
        "valor": "Camiseta",
        "tipoDato": "string"
      },
      "descripcion": {
        "valor": "Camiseta roja de algodón",
        "tipoDato": "string"
      },
      "precio": {
        "valor": 25000,
        "tipoDato": "integer"
      }
    }
  },
  {
    "tipo": "Category",
    "atributos": {
      "propiedad1": {
        "valor": "valor1",
        "tipoDato": "string"
      },
      "propiedad2": {
        "valor": "valor2",
        "tipoDato": "string"
      }
    }
  }
]

~~~~ 