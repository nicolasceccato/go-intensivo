package main

//
//func main() {
//	println("Hello World")
//}

type Car struct {
	Model string
	Color string
}

func main() {
	car := Car{
		Model: "Lamborghini",
		Color: "yellow",
	}
	println("Model: " + car.Model + " Color: " + car.Color)
}
