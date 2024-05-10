package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* this is main function,
which is called when application starta*/

// functions in GO.
func multiply(num1, num2 float64) float64 {
	var result float64
	result = num1 * num2
	return result
}

func addMultiply(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}

// variadic functions in Go.
func addition(numbers ...int) {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println(sum)
}

func checksFileExist() {
	if _, err := os.Stat("file.go"); err == nil {
		fmt.Println("file exist")
	} else {
		fmt.Println("File does not exist")
	}
}

func createAndWriteToFile() {
	filePointer, err := os.Create("file.go")
	if err != nil {
		fmt.Println("Error while creating file")
	}
	var cities = []string{"Hyderabad", "Mumbai", "Chennai"}

	for _, val := range cities {
		filePointer.WriteString(val)
		filePointer.WriteString(" ")
	}
}

func readFile() {
	data, err := os.ReadFile("file.go")
	if err != nil {
		fmt.Println("error while reading file")
	}
	fmt.Println(data)         // byte data.
	fmt.Println(string(data)) // convert byte data to string
}

func readLineByLineFile() {
	filePointer, err := os.Open("file.go")
	if err != nil {
		fmt.Println("error while opening file")
	}

	var fileData []string //created a array with 0 len,0 cap.
	fmt.Println(len(fileData), cap(fileData))
	var scanner = bufio.NewScanner(filePointer)
	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}

	fmt.Println(fileData)

}

func main() {
	numbers := [11]uint{0: 1, 2: 1}
	names := [2]string{"rkesh", "rohit"}
	for counter := 0; counter < len(numbers)-1; counter++ {
		fmt.Print(numbers[counter])
	}

	//range in Go.
	for _, values := range names {
		fmt.Println(values)
	}
	fmt.Println(numbers)
	fmt.Println(names)

	if numbers[0] == 1 {
		fmt.Println("number is 1")
	}

	//while loop in Go.
	i := 0
	for i < 3 {
		println(i);
		i++
	}

	var fullName = strings.Join([]string{"Rakesh", "Maddineni", "Venkata"}, "@")

	fmt.Println(fullName)

	//structs in Go

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	var person1 Person
	person1.firstName = "Rakesh"
	person1.lastName = "Maddineni"
	person1.age = 23

	fmt.Println(person1)

	//Maps in Go
	var countryAndCapitalCities = make(map[string]string) // creates an empty map.
	countryAndCapitalCities["India"] = "NewDelhi"
	countryAndCapitalCities["USA"] = "Washington DC"
	countryAndCapitalCities["UK"] = "London"

	var miPlayers = map[string]string{"Rohith": "India", "SKY": "India"} // intialize map with data.
	miPlayers["Bumrah"] = "India"                                        // add new key value pair.
	miPlayers["David"] = "Australia"

	fmt.Println(miPlayers["David"])

	miPlayers["David"] = "Singapore"

	fmt.Println(miPlayers["David"])

	//pointers in Go
	var age = 20
	fmt.Println((age))

	var pointer *int = &age // stores memory address of age variable.
	fmt.Println(pointer)    // prints memory address of age variable.

	fmt.Println((*pointer)) // accessing age variable data.

	*pointer = 30

	fmt.Println(age)

	//Slices in Go

	var number = make([]int, 4, 10) // slice of length 4, capacity 10
	number[0] = 1
	fmt.Println(len(number), cap(number))

	//defer statement in Go. Used to call function,once all code is executed in current parent function.
	defer fmt.Println((number))

	//call multiply function in Go
	fmt.Println((multiply(2, -2)))
	var add, mul int = addMultiply(10, 5)
	fmt.Println(add, mul)

	//call variadic function.
	addition(10, 20, 30, 40, 50)

	//Files
	createAndWriteToFile()
	checksFileExist()
	readFile()
	readLineByLineFile()

}
