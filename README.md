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


# Testing

## Ejercicio 1 - Testear el impuesto del salario
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:

Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/UnitTestGo/FuncionesEnGoTest/ImpuestosDeSalario/ImpuestosDeSalario_test.go)

## Ejercicio 2 - Calcular promedio
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos corresponde realizar los test correspondientes:

Calcular el promedio de las notas de los alumnos.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/UnitTestGo/FuncionesEnGoTest/CalcularPromedio/CalcularPromedio_test.go)

## Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:

Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/UnitTestGo/FuncionesEnGoTest/CalcularSalario/CalcularSalario_test.go)

## Ejercicio 4 - Testear el cálculo de estadísticas
Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:

Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.


### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/UnitTestGo/FuncionesEnGoTest/CalcularEstadisticas/CalcularEstadisticas_test.go)

## Ejercicio 5 - Calcular cantidad de alimento

El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios para verificar que todo funcione correctamente:

Verificar el cálculo de la cantidad de alimento para los perros.
Verificar el cálculo de la cantidad de alimento para los gatos.
Verificar el cálculo de la cantidad de alimento para los hamster.
Verificar el cálculo de la cantidad de alimento para las tarántulas.

Ejemplo:

```go
func TestDog(t *testing.T)

func TestCat(t *testing.T)
```

### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/UnitTestGo/FuncionesEnGoTest/CalcularCantidadAlimento/CalcularCantidadAlimento_test.go)



# Composición 

## Ejercicio 1 Product

Crear un programa que cumpla los siguiente puntos:

1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.   
2. Tener un slice global de Product llamado Products instanciado con valores.
3. métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
4. Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
5. Ejecutar al menos una vez cada método y función definido desde main().


### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasYMetodos/Productos/main.go)



## Ejercicio 2 - Employee

Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:

1. Crear una estructura Person con los campos ID, Name, DateOfBirth.
2. Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
3. Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
4. Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().

5. Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.


### [Solucion](https://github.com/cdmorau023/go_learning/tree/main/EstructurasYMetodos/Employee/main.go)


