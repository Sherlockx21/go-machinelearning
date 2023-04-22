package vector

import "fmt"

func Example() {
	//Initialize
	var myvec []float64

	myvec = append(myvec, 11.0)
	myvec = append(myvec, 5.3)

	fmt.Println(myvec)
}
