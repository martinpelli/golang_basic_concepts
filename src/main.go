package main

import "fmt"

func main() {

	//Declaración de constantes

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables

	numero := 12
	var altura int = 14
	var area int

	fmt.Println(numero, altura, area)

	//Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println(areaCuadrado)

	//Operadores aritméticos

	x := 10
	y := 50

	result := x + y

	fmt.Println(result)

	const radio float64 = 5
	const h float64 = 4.43
	const z float64 = 3
	const w float64 = 5

	areaCirculo := pi * radio * h
	areaRectangulo := radio * h
	areaTrapecio := h * ((z + w) / 2)

	fmt.Println(areaCirculo, areaRectangulo, areaTrapecio)

	//Incremental
	x++
	//Decremental
	x--

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	fmtPackage()

	functions()

	loops()

}

func fmtPackage() {
	helloMessage := "hello"
	broMessage := "bro"

	fmt.Println(helloMessage, broMessage)
	fmt.Println(helloMessage, broMessage)

	club := "CSD"
	years := 100

	fmt.Printf("%s tiene más de %d años\n", club, years)
	fmt.Printf("%v tiene más de %v años\n", club, years)

	message := fmt.Sprintf("%v tiene más de %v años", club, years)

	fmt.Println(message)

	//tipo de datos
	fmt.Printf("el tipo de dato es %T\n", club)
}

func functions() {
	paramFunction("que tal")
	paramamsFunction(1, 2, "hi")
	returnedValue := returnFunction(2)
	returnedValue1, returnedValue2 := returnsFunction(2)
	returnedValue3, _ := returnsFunction(2)
	fmt.Println(returnedValue)
	fmt.Println(returnedValue1, returnedValue2)
	fmt.Println(returnedValue1, returnedValue3)
}

func paramFunction(message string) {
	fmt.Println(message)
}

func paramamsFunction(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnFunction(a int) int {
	return a * 2
}

func returnsFunction(a int) (b, c int) {
	return a * 2, a
}

func loops() {

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")

	counter := 0
	for counter < 10 {
		counter++
		fmt.Print(counter)
	}
	fmt.Print("\n")

	listaNumerosPares := []int{2, 4, 6, 8}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	for {
		//infinito
	}

}
