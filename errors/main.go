package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func divide2(a, b int) (int, error) {
	// this function will recover from the panic and continue the execution, and control panic
	defer func() { // defer will be executed at the end of the function
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()
	validateZero(b)
	return a / b, nil
}

func validateZero(a int) error {
	if a == 0 {
		panic("Cannot divide by zero")
	}
	return nil
}

func main() {

	/* srt := "123"
	num, err := strconv.Atoi(srt)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("No error", num)
	*/

	/*r, e := divide(10, 0)
	if e != nil {
		fmt.Println("Error: ", e)
		return
	}
	fmt.Println("No error", r) */
	/*
		// with error
		srt = "123a"
		num, err = strconv.Atoi(srt)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	*/

	// defer --------- defer will be executed at the end of the function
	/* defer fmt.Println("End of the function")
	fmt.Println("Start of the function")
	fmt.Println("End of the function")*/

	// if all have defer then it will be executed in reverse order, with a stack
	/*
		file, e := os.Create("test.txt")
		if e != nil {
			fmt.Println("Error: ", e)
			return
		}

		defer file.Close() // close the flow at the end of the function

		_, e = file.Write([]byte("Hello World"))
		if e != nil {
			fmt.Println("Error: ", e)
			return
		}
	*/
	// panic and recover -> panic is used to stop the execution of the program and recover is used to recover from the panic

	// panic interrupts the flow of the program
	/*divide2(10, 10)
	divide2(10, 0)
	divide2(10, 5)
	*/

	// register errors
	log.SetPrefix("main: ") // set the prefix for the log
	//log.Print("Hello World") // print the log
	//log.Fatal("Error: ")     // fatal and stop the execution
	//log.Panic("Error: ")     // panic and stop the execution
	//log.Panicln("Error: ")   // panic and stop the execution

	// register errors in a file
	file, error := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if error != nil {
		log.Fatal("Error: ", error)
		return
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print("I am a log")

}
