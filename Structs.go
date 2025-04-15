package array

import (
	"fmt"
)

type Days struct {
	Monday    int
	Tuesday   int
	Wednesday string
}

func Run() {
	var Day = Days{1, 2, "three"}
	fmt.Println(Day)

}
