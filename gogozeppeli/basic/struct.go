package basic

import "fmt"

type rect struct {
	width  int
	height int
}

// receiver
func (r rect) area() int {
	return r.width * r.height
}

func Struct() {
	// norm
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"Alice", 25}
	fmt.Println("person.Name: " + person.Name)
	fmt.Println("person.Age: " + fmt.Sprint(person.Age))

	// Anonymous struct
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "Model 3",
	}
	fmt.Println(myCar)

	// Anonymous struct in a struct
	type car struct {
		Make  string
		Model string
		Wheel struct {
			Radius   int
			Material string
		}
	}

	type bus struct {
		make  string
		model string
	}

	// Embedded struct
	// truct contains all props in bus
	type truck struct {
		bus
		bedSize int
	}

	lanesTruck := truck{
		bedSize: 10,
		bus: bus{
			make:  "toyota",
			model: "camry",
		},
	}
	// truck.make to call "make" prop, instead of truck.bus.make
	fmt.Println(lanesTruck.model, lanesTruck.bedSize)

	r := rect{
		width:  5,
		height: 19,
	}
	fmt.Println(r.area())
}
