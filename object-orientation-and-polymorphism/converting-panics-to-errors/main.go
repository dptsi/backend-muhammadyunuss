package main

import (
	"errors"
	"fmt"
)

func main() {
	// result, err := divide1(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("result:", result)

	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result:", result)
}

func divide(l, r int) int {
	return l / r
}

func divide1(l, r int) (int, error)  {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return l / r, nil
}

func divide2(l, r int) (result int, err error)  {
	defer func ()  {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l/r, nil
}