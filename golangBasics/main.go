package main

import (
	"errors"
	"fmt"
	"strconv"
)

func sum(num1 int, num2 int) int {
	// ==== that take two variable of int type and return a single int data ====
	return num1 + num2
}

func sum2(num1, num2 int, firstName, lastName string) int {
	// ==== that take two variable of int type and return a single int data ====
	// == we can describe same data type variable at once at last variable place
	return num1 + num2
}

func callbackFunc(f func(int, int) int, num1, num2 int) int {
	// ===== sum is function type, that take 2 int data and return a int data, (like callback in js)
	return f(num1, num2)
}

func incrementByValue(x int) {
	x++
}

func incrementByReference(x *int) {
	*x++
}

func switchCaseExample(number int) {
	// in switch case we do not have to mention break in each case
	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Please passs value from 1-7")
	}
}

// named return data
func namedReturn() (firstName string, lastName string) {
	firstName = "bipul"
	lastName = "dubey"
	return
}

// ====================== struct ====================
type Wheel struct {
	Radius   int
	Material string
}

type Car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	RearWheel  Wheel
}

// EMBEDDED VS NESTED
// Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate string
}

type Student struct {
	Course  string
	Class   int
	Section string
	User
}

// ===== STRUCT METHODS IN GO ======
type rect struct {
	width  int
	height int
}

// Receiver function = it bind method to struct
func (r rect) area() int {
	return r.width * r.height
}

// ============ Interface ============
// Interfaces are just collections of method signatures.
// we can add only those method which is common
type shape interface {
	Area()
	// example (value string) int
}

type circle struct {
	radius int
}

type rectangle struct {
	height int
	width  int
}

func (c circle) Area() {
	fmt.Println(3.14 * float64(c.radius) * float64(c.radius))
}

func (r rectangle) Area() {
	fmt.Println(r.height * r.width)
}

// ======== interface assertion ==========
type expenses interface {
}

type email struct {
	isSubscribed bool
	body         string
}

type sms struct {
	isSubscribed bool
	body         string
}

func getExpenses(e expenses) {
	em, ok := e.(sms)
	fmt.Println(em, ok)
}

// === VARIADIC function ===
func concatString(strs ...string) string {
	// strs is just a slice of strings
	final := ""
	for i := 0; i < len(strs); i++ {
		final += strs[i]
	}
	return final
}

func main() {
	// ====== printing ========
	fmt.Print("print in single line =========")
	fmt.Printf("print in single line, with string formatting %s ======= ", "this is string")
	fmt.Println("print and change line")

	// == Data type Conversion ==
	var floatNumber float64 = 2.4
	fmt.Println(int(floatNumber), int64(floatNumber))

	intNumber := 20
	fmt.Println(float64(intNumber))

	const pi float64 = 3.14 // const does not support short hand declaration
	fmt.Println(pi)

	msg := fmt.Sprintf("Hello this is value of pi: %.3f", pi) // sPrintf, sPrint, sPrintln return a (formatted) string
	fmt.Println(msg)

	// if statement
	// if len(msg) > 5 {
	// 	fmt.Println("long message")
	// }
	// declare variable in if statement and use in that scope only (scope level)
	if msglen := len(msg); msglen > 5 {
		fmt.Println("long message")
	}

	//  ==== function =====
	fmt.Println(callbackFunc(sum, 10, 23))

	fmt.Println("========== passing variable by value ===============")
	val := 5
	incrementByValue(val)
	fmt.Println(val)
	incrementByReference(&val)
	fmt.Println(val)

	// named return value from function
	firstName, lastNames := namedReturn()
	fmt.Println("named return", firstName, lastNames)

	// ======== struct ========
	// structs in Go to represent structured data. It's often convenient to group different types of variables together
	myCar := Car{Make: "BMW", Model: "RR1000", Height: 5, Width: 4, FrontWheel: Wheel{Radius: 3, Material: "Rubber"}, RearWheel: Wheel{Radius: 4, Material: "Rubber"}}
	fmt.Println(myCar)
	fmt.Println(myCar.FrontWheel.Material)

	// anonymous struct - just like normal struct but it is defined without name and therefore cannot use later
	// immediately using a second pair of brackets
	myCar2 := struct {
		Make  string
		Model string
	}{Make: "Tesla", Model: "model 3"}

	fmt.Println(myCar2)

	// in above example we can see the use of nested struct, and accessing nested struct using dot annotation again and again

	// ======= Embedded struct =======
	stud1 := Student{User: User{FirstName: "Bipul", LastName: "Dubey", Age: 24, BirthDate: "01/01/2000"}, Course: "MCA", Class: 2, Section: "B"}
	fmt.Println(stud1)
	fmt.Println(stud1.FirstName) // in embedded struct we can access nested value directly instead of using dot annotation again and again

	rect1 := rect{4, 5}
	fmt.Println(rect1)
	fmt.Println(rect1.area())

	// ========= Interface ===========
	shape := []shape{
		circle{
			radius: 7,
		},
		circle{
			radius: 5,
		},
		rectangle{
			height: 4,
			width:  6,
		},
		rectangle{
			height: 3,
			width:  5,
		},
	}

	for _, v := range shape {
		v.Area()
	}

	// ======= assertion ========
	sms1 := sms{
		isSubscribed: true, body: "hello",
	}
	fmt.Print("sms1", sms1)
	getExpenses(sms1)

	// ======== Errors handling ============
	/*
		Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:
		type error interface {
		    Error() string
		}
	*/
	// Because errors are just interfaces, you can build your own custom types that
	// implement the error interfaceBecause errors are just interfaces, you can build your
	// own custom types that implement the error interface
	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("couldn't convert:", err)
	}
	fmt.Println(i)

	var err1 error = errors.New("something went wrong")
	fmt.Println(err1)

	fmt.Println("======== Loop =========")
	/*
		for INITIAL; CONDITION; UPDATE {
		}

		// loop continue
		for INITIAL; ; UPDATE {
		}

		// go do not have d0-while, this for loop run until condition is true
		for CONDITION {
		}

	*/
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	// ======== Array and slice =========
	/* Array
	1. fixed in size
	*/
	var myInts [10]int
	fmt.Println(myInts, len(myInts))

	even := [5]int{0, 2, 4, 6, 8}
	fmt.Println(even)

	names := [3]string{"Ajay", "Aman", "Abhi"}
	fmt.Println(names)

	/* Slices
	1. Dynamic in size - A slice is a dynamically-sized, flexible view of the elements of an array.
	2. len = return the number of element present in current slice
	3. cap = return the number of element can be add before resigning that slice into new
	*/
	var odd = []int{}
	fmt.Println("size ", len(odd), " cap ", cap(odd))
	odd = append(odd, 13)
	odd = append(odd, 23)
	odd = append(odd, 3)
	odd = append(odd, 33)
	fmt.Println("size ", len(odd), " cap ", cap(odd))
	fmt.Println(odd)

	/* ========== Make ============
	MAKE : Most of the time we don't need to think about the underlying array of a slice.
	We can create a new slice using the make function */

	mySlice := make([]int, 5, 10) // int type, len = 5, cap = 10
	fmt.Println(mySlice, len(mySlice), cap(mySlice))

	// === VARIADIC function ===
	fmt.Println(concatString("hello", "testing", "concat"))

	// === SPREAD Operator === work only for SLICE not ARRAY
	names2 := []string{"Ajay", "Aman", "Abhi"}
	fmt.Println(concatString(names2...))
}
