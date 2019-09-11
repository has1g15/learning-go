package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	x := 10
	y := 3
	total := x * y

	fmt.Println(total)

	const pi = 3.14159265
	fmt.Printf("%.3f \n", pi)

	if x > 5 {
		fmt.Println("More than 5")
	}

	var a [5]int
	var b [5]string

	a[2] = 1
	b[3] = "not empty"

	c := [5]int{1,2,3,4,5}

	//slices
	d := []int{1,2,3,4,5}
	d = append(d,6)
	e := d[3:6]

	f := make([]int, 5, 10)
	copy(f, d)

	fmt.Println(a, b, c, d, e, f)

	//maps
	aMap := make(map[string]int)

	aMap["Entry 1"] = 1
	aMap["Entry 2"] = 2
	aMap["Entry 3"] = 3

	delete(aMap, "Entry 1")

	fmt.Println(aMap)

	//loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	//equivalent of while
	for i < 5 {
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	for key, value := range aMap {
		fmt.Println("key:", key, "value:", value)
	}

	//functions
	fmt.Println(sum(7,8))

	result, err := sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	num1, num2 := next2values(5)
	fmt.Println(num1, num2)

	fmt.Println(subtracting(1,2,3,4,5))

	//structs
	p := person{name: "Hannah", age:22}
	fmt.Println(p.name)

	//assign methods to structs
	fmt.Println(p.stateMyAge())

	//pointers - pass memory address, not pass by value
	j := 3
	fmt.Println(&j)

	increment(&j)
	fmt.Println(j)

	//closures
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum())

	//recursion
	fmt.Println(factorial(3))

	//defer
	defer printTwo()
	printOne()

	//recover
	fmt.Println(safeDivision(3,0))

	//interfaces
	rect := Rectangle{20,50}
	circ := Circle{4}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area=", getArea(circ))

	//string manipulation
	sampleString := "Hello World"

	fmt.Println(strings.Contains(sampleString, "Hello"))
	fmt.Println(strings.Index(sampleString, "lo"))
	fmt.Println(strings.Count(sampleString, "l"))
	fmt.Println(strings.Replace(sampleString, "World", "Hannah", 1))
	fmt.Println(strings.Split(sampleString, " "))

	letters := []string{"c", "a", "b"}
	sort.Strings(letters)
	fmt.Println("Letters:", letters)
	letterList := strings.Join(letters, ", ")
	fmt.Println(letterList)

	//file io
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("First Line")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")

	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)

	//casting
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

	//create web server
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/world", handler2)
	//http.ListenAndServe(":8080", nil)

	//go routines
	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)

	//channels
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}

type person struct {
	name string
	age int
}

func (person *person) stateMyAge() string {
	return person.name + " is " + strconv.Itoa(person.age)
}

func sum(x int, y int)int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func increment(x *int) {
	*x++
}

func next2values(number int) (int, int) {
	return number+1, number+2
}

func subtracting(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

func safeDivision(num1 int, num2 int) int {
	defer func() {
		fmt.Println(recover())
	} ()

	solution := num1/num2
	return solution
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Second World")

}

func count(id int) {

	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 1000)
	}
}

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make dough, send for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 1000)
}

func addSauce(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Add sauce and send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 1000)
}

func addToppings(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Add toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 1000)
}