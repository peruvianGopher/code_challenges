// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var ints = map[byte]int{
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}

func main() {
	a := "1+2"
	//b := ints[a[0]] + ints[a[2]]
    b := (a[0] - '0') + (a[2] - '0')
	fmt.Println(b)
    Atoi()
}
