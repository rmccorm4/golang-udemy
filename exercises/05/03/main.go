package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type suv struct {
	v       vehicle
	compact bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "orange",
		},
		luxury: false,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.doors)
	fmt.Println(s1.color)

	suv1 := suv{
		v: vehicle{
			doors: 4,
			color: "black",
		},
		compact: true,
	}

	fmt.Println(suv1)
	fmt.Println(suv1.v)
	fmt.Println(suv1.v.doors)

	fmt.Println(t1)
	fmt.Println(t1.vehicle)
	fmt.Println(t1.vehicle.doors)
}
