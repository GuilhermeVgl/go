package main

type Car struct {
	Model string
	Color string
}

func (c Car) getModelo() {
	println("Modelo do carro: " + c.Model)
}

func (c *Car) changeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

func soma(x, y int) int {
	return x + y
}

func main() {

	car := Car{
		Model: "Ferrari",
		Color: "Red",
	}

	car.getModelo()

	car.changeColor("Blue")
	println(car.Color)

}
