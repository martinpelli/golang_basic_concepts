package main

import (
	"fmt"
	"github.com/martinpelli/golang_platzi/src/mypackage"
	"log"
	"strconv"
	"strings"
	"sync"
	"net/http"
	"github.com/labstack/echo/v4"
)

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

	parse()

	keywords()

	arrays()

	slices()

	fmt.Println(isPalindromo("amor a roma"))
	fmt.Println(isPalindromo("ama"))

	maps()

	structs()

	pointers()

	stringers()

	interfaces()

	goroutines()

	anonymous()

	channels()

	echoModule()

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

	//for {
	//infinito
	//}

}

func parse() {
	s, errors := strconv.Atoi("535")
	b, errorb := strconv.ParseBool("true")
	f, errorf := strconv.ParseFloat("3.1415", 64)
	i, errori := strconv.ParseInt("-42", 10, 64)
	u, erroru := strconv.ParseUint("42", 10, 64)

	if errors != nil {
		log.Fatal(errors)
	}
	if errorb != nil {
		log.Fatal(errorb)
	}
	if errorf != nil {
		log.Fatal(errorf)
	}
	if errori != nil {
		log.Fatal(errori)
	}
	if erroru != nil {
		log.Fatal(erroru)
	}

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(u)
}

func keywords() {
	//Defer
	defer fmt.Println("últumo")
	fmt.Println("primero")

	//continue y break para loops
}

func arrays() {
	var array [4]int
	array[0] = 1
	fmt.Println(array, len(array), cap(array))
}

func slices() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	slice = append(slice, 7)

	newSlice := []int{8, 9}
	slice = append(slice, newSlice...)

	for i, value := range slice {
		fmt.Println(i)
		fmt.Println(value)
	}

}

func maps() {
	m := make(map[string]int)
	m["jose"] = 14
	m["jose"] = 12
	fmt.Println(m)

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	for i, value := range temperature {
		fmt.Println(i, value)
	}

	apellido, ok := m["pellicer"]
	if ok {
		fmt.Println(apellido)
	} else {
		fmt.Println("No existe ese apellido")
	}

}

func isPalindromo(text string) bool {
	return text == strings.ToLower(reverse(text))
}

func reverse(str string) (result string) {
	for _, v := range str {
		result += string(v)
	}
	return
}

func structs() {
	type car struct {
		model int
		brand string
	}

	duster := car{model: 2018, brand: "Renault"}

	sharan := car{}
	sharan.brand = "Volkswagen"
	sharan.model = 2012

	fmt.Println(duster, sharan)

	python := mypackage.Language{Name: "python", Type: "Backend"}
	fmt.Println(python)

}

func pointers() {
	a := 10
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	myPc := pc{ram: 8, rom: 500, processor: "i5"}

	myPc.ping()
	myPc.duplicateRam()
	fmt.Println(myPc)

}

type pc struct {
	ram       int
	rom       int
	processor string
}

func (myPc pc) ping() {
	fmt.Println(myPc.processor, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram *= 2
}

func stringers() {
	myPC := pc{ram: 8, rom: 16, processor: "amd"}
	fmt.Println(myPC)
}

func (myPc pc) String() string {
	return fmt.Sprintf("%d poquito", myPc.ram)
}

func interfaces() {
	cuadra := cuadrado{base: 10}
	recta := rectangulo{base: 10, altura: 5}
	calcularArea(cuadra)
	calcularArea(recta)

	myInterface := []interface{}{"hola", 1, 1.2}
	fmt.Println(myInterface...)

}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type figuras interface {
	area() float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcularArea(f figuras) {
	fmt.Println(f.area())
}

func goroutines() {
	var wg sync.WaitGroup

	fmt.Println("First")
	wg.Add(1)

	go say("Second", &wg)

	wg.Wait()

}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func anonymous() {
	go func(text string) {
		fmt.Println(text)
	}("mensaje")
}

func channels(){
	c := make(chan string, 1)

	fmt.Println("Hello")
	go  say2("Bye", c)
	fmt.Println(<-c)

	d := make(chan string, 2)
	d <- "Message1"
	d <- "Message2"

	fmt.Println(len(d),cap(d))

	close(d)

	for message := range d {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email12 := make(chan string)

	go message("message 1", email1)
	go message("message 2", email12)

	for i:= 0; i < 2; i++{
		select{
		case m1 := <-email1:
			fmt.Println("email recibido de email1", m1)
		case m2 := <-email12:
			fmt.Println("mail recibido de email2", m2)
		}
	}


}

func say2(text string, c chan<- string){
	c <- text
}

func message (text string, c chan string){
	c <- text
}

func echoModule(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1323"))
}