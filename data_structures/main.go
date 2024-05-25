package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	name  string
	age   int
	email string
}

func (p *Person) sayHello() {
	fmt.Println("Hello", p.name)
}

type task struct {
	title       string
	description string
	done        bool
}

type taskList struct {
	tasks []task
}

// add task
func (t *taskList) addTask(tsk task) {
	t.tasks = append(t.tasks, tsk)
}

// remove task
func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

// done task
func (t *taskList) taskAsDone(index int) {
	t.tasks[index].done = true
}

// edit task
func (t *taskList) editTask(index int, tsk task) {
	t.tasks[index] = tsk
}

// print tasks
func (t *taskList) printTasks() {
	for i, t := range t.tasks {
		fmt.Println(i, t.title, t.description, t.done)
	}
}

func main() {

	taskListProject()

	var matrix [5]int
	fmt.Println(matrix) // [0 0 0 0 0]
	matrix[0] = 1
	matrix[1] = 4

	var m2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(m2)

	// unspecified length
	var m3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(m3)

	// access to element
	fmt.Println(m3[2])

	// length
	fmt.Println(len(m3))

	for i := 0; i < len(m3); i++ {
		fmt.Println(m3[i])
	}

	// range index and value
	for i, v := range m3 {
		fmt.Println(i, v)
	}

	// matrix bidimensional
	var m4 = [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(m4)

	// slice -> you can change the length of the array
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	var newSlice = s1[0:3]
	fmt.Println(newSlice)

	// length
	fmt.Println("Length: ", len(s1))

	// capacity
	fmt.Println("Capacity: ", cap(s1))

	// append
	s1 = append(s1, 6)
	fmt.Println(s1)

	// remove element
	s1 = append(s1[:2], s1[3:]...) // ... -> unpack the slice, index 2 is removed
	fmt.Println("Remove", s1)

	// slice from array
	var a1 = [5]int{1, 2, 3, 4, 5}
	var s2 = a1[1:3]
	fmt.Println(s2)

	// make -> create a slice with a length and capacity
	var s3 = make([]int, 3, 5) // length 3, capacity 5
	fmt.Println("Make ", s3)

	// copy -> copy the content of a slice to another
	var s4 = []int{1, 2, 3}
	var s5 = make([]int, 2)
	copy(s5, s4) // copy s4 to s5 -> s5 = [1, 2]
	fmt.Println("Copy ", s5)

	// map -> key value
	colors := map[string]string{ // key string, value string
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])

	// add element
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	value := colors["white"]
	fmt.Println(value)

	// delete element
	delete(colors, "white")
	fmt.Println(colors)

	// check if key exists
	value, exists := colors["white"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exist")

	}

	// range
	for key, value := range colors {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	// struct -> custom data type with fields
	var p1 Person
	p1.name = "John"
	p1.age = 25
	p1.email = "mail@gmail.com"
	fmt.Println(p1)
	// if you don't assign a value to a field, it will be the default value

	p2 := Person{name: "Jane", age: 30, email: "jij@gmail.com"}
	fmt.Println(p2)

	// access to field
	fmt.Println(p2.name)

	// modify field
	p2.age = 35
	fmt.Println(p2)

	// struct can be used for representing an object

	// pointer -> memory address
	// method -> function that belongs to a struct

	// pointer
	var x int = 10
	var p *int = &x // memory address
	fmt.Println(p)
	fmt.Println(*p) // value
	fmt.Println(&x) // memory address
	fmt.Println(x)  // value

	// change value
	changeValue(&x)
	fmt.Println(x)

	// create an instance of a struct
	p3 := Person{name: "Alice", age: 25, email: "mail@gmail.com"}
	p3.sayHello()
}

// function that receives a pointer
func changeValue(x *int) {
	*x = 20
}

// project task list
func taskListProject() {
	list := taskList{}

	// bufio -> read from the console, os.Stdin -> input
	read := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Select an option: \n",
			"1. Add task\n",
			"2. Remove task\n",
			"3. Edit task\n",
			"4. Mark task as done\n",
			"5. Print tasks\n",
			"6. Exit\n",
		)
		fmt.Println("Enter an option:")
		fmt.Scanln(&option)

		switch option {
		case 1: // add task
			var tsk task
			fmt.Println("Enter the title of the task:")
			title, _ := read.ReadString('\n')
			tsk.title = title
			fmt.Println("Enter the description of the task:")
			description, _ := read.ReadString('\n')
			tsk.description = description
			tsk.done = false
			list.addTask(tsk)
		case 2: // remove task
			list.printTasks()
			var index int
			fmt.Println("Enter the index of the task to remove:")
			fmt.Scanln(&index)
			list.removeTask(index)
		case 3: // edit task
			var index int
			fmt.Println("Enter the index of the task to edit:")
			fmt.Scanln(&index)
			var tsk task
			fmt.Println("Enter the title of the task:")
			title, _ := read.ReadString('\n')
			tsk.title = title
			fmt.Println("Enter the description of the task:")
			description, _ := read.ReadString('\n')
			tsk.description = description
			list.editTask(index, tsk)
		case 4: // mark task as done
			var index int
			fmt.Println("Enter the index of the task to mark as done:")
			fmt.Scanln(&index)
			list.taskAsDone(index)
		case 5:
			list.printTasks()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}

	// list of tasks
	for i, t := range list.tasks {
		fmt.Printf("%v: %v\n", i, t.title)
	}
}
