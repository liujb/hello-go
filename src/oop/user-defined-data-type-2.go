package main

import "fmt"

// 定义枚举
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// 定义三个类型
type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box

// 为Box类型定义Volume方法
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 为Box指针定义SetColor方法
// Box和*Box有啥区别？
// Box是传递过来的对象的copy
// *Box是真正的指向的那个对象
func (b *Box) SetColor(c Color) {
	b.color = c
}

// 为BoxList类型BiggestsColor方法
func (boxList BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range boxList {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

// 为BoxList类型定义PaintItBlack方法
func (boxList BoxList) PaintItBlack() {
	for i, _ := range boxList {
		boxList[i].SetColor(BLACK)
	}
}

// 为Color类型定义string方法
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{3, 4, 9, RED},
		Box{1, 4, 1, BLUE},
		Box{4, 2, 4, BLACK},
		Box{4, 5, 4, YELLOW},
		Box{8, 4, 4, WHITE},
	}

	fmt.Println(boxes[0].Volume())
	fmt.Println(boxes[0].color.String())        // RED
	fmt.Println(boxes.BiggestsColor().String()) // WHITE

	boxes.PaintItBlack()
	fmt.Println(boxes[0].color.String())        // BLACK
	fmt.Println(boxes.BiggestsColor().String()) // BLACK

	box := Box{1, 2, 2, WHITE}
	//box.SetColor(3)
	box.SetColor(YELLOW)
	fmt.Println(box.color.String()) // YELLOW
}
