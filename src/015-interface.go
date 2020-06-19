package src

import (
	"fmt" 
	"math"
)

/*Shape define an interface */
type Shape interface {
   area() float64
}

/*Circle define a circle */
type Circle struct {
   x,y,radius float64
}

/*Rectangle define a rectangle */
type Rectangle struct {
   width, height float64
}
// Kubus adalah struc
type Kubus struct {
	width, height float64
 }

func (kubus Kubus) area() float64{
	return kubus.width*kubus.height
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
   return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}
// InterfaceExamp contohnya
func InterfaceExamp() {
   	circle := Circle{x:0,y:0,radius:5}
   	rectangle := Rectangle {width:10, height:5}
	kubus := Kubus{width:2,height:4}
   	fmt.Printf("Circle area: %f\n",getArea(circle))
   	fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
   	fmt.Printf("Kubus area: %f\n",getArea(kubus))
}