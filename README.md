# go_learning
Repositorio para guardar aprendizaje básico de go en Meli


# Primer acercamiento a go


## Ejercicio 1 - Imprimí tu nombre
Tendrás que:

Crear una aplicación donde tengas como variable tu nombre y dirección.
Imprimir en consola el valor de cada una de las variables.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/PrimerAcervamientoGo/imprimirNombre/main.go)



## Ejercicio 2 - Clima


Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.

Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/PrimerAcervamientoGo/Clima/main.go)


## Ejercicio 3 - Declaración de variables


Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.

Necesita ayuda para:

Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.

```go
var 1firstName string
var lastName string
var int age
1lastName := 6
var driver_license = true
var person height int
childsNumber := 2
```

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/PrimerAcervamientoGo/DeclaracionVariables/main.go)



## Ejercicio 4 - Tipos de datos

Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:

Verificar su código y realizar las correcciones necesarias.

```go
var lastName string = "Smith"
var age int = "35"
boolean := "false"
var salary string = 45857.90
var firstName string = "Mary"
```
### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/PrimerAcervamientoGo/TiposDeDato/main.go)


--------------
# Estructuras de control


## Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:

Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimir cada una de las letras.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasDeControl/LetrasDeUnaPalabra/main.go)



## Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.

Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasDeControl/Prestamo/main.go)


## Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes.

Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.

Nota: Validar que el número del mes, sea correcto.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasDeControl/QueMesCorresponde/main.go)



## Ejercicio 4 - Qué edad tiene...

Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

```go
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
```

Por otro lado también es necesario:

Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasDeControl/QueEdadTiene/main.go)

---

# Funciones en Go

## Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/FuncionesEnGo/ImpuestosDeSalario/main.go)



## Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/FuncionesEnGo/CalcularPromedio/main.go)



## Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/FuncionesEnGo/CalcularSalario/main.go)





## Ejercicio 4 - Calcular estadísticas


Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

Ejemplo:

```go
const (|    123
minimum = "minimum"
average = "average"
maximum = "maximum"
)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...


minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

```

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/FuncionesEnGo/CalcularEsradisticas/main.go)



## Ejercicio 5 - Calcular cantidad de alimento


Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:

Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:

```go
const (
dog    = "dog"
cat    = "cat"
)

...

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

...

var amount float64
amount += animalDog(10)
amount += animalCat(10)
```

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/FuncionesEnGo/CalcularCantidadAlimento/main.go)

