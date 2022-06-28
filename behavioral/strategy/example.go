package strategy

import "fmt"

func Example() {
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(Context)

	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(data1)

	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(data2)

	fmt.Println("sorted slices:")
	fmt.Println(data1)
	fmt.Println(data2)
}
