package golang_united_school_homework

import (
	"fmt"
	"strconv"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	cap := strconv.Itoa(b.shapesCapacity)
	if len(b.shapes) > b.shapesCapacity {
		return fmt.Errorf("the expected length is "+cap+", but actual %d", len(b.shapes))
	} else if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("max cap reached")
	}
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes)-1 < i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes)-1 < i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}
	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes)-1 < i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}
	old := b.shapes[i]
	b.shapes[i] = shape
	return old, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	panic("implement me")

}
