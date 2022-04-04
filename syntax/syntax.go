package syntax

import "fmt"

// Các kiểu khai báo
func Variable() int {
	var a int  // or a := 15
	var b bool // or b := true
	var (
		c int
		d bool
	)
	const (
		e = iota
		f
	)
	fmt.Sprintf("%d %v %d %v %d %d", a, b, c, d, e, f)
	return 0

}
