package main

func main() {
	age := 17

	if age >= 18 {
		println("person is adult")
	} else if age >= 12 {
		println("person is teenager")
	} else {
		println(("Person is a kid"))
	}

	//or condition
	// age >= 12 || age <=10

	//and condition  age>12 && age ==12

	// declare varibale in condition

	if age := 15; age > 12 {
		print("something")
	}

	// go does not have ternary , we will have to use normal if else
}
