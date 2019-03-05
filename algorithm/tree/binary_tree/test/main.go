package main

import (
	"algorithm/tree/binary_tree"
	"fmt"
)

func main() {
	var err error

	err = bt.TransferLeftMiddleRight(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = bt.TransferLeftMiddleRightBad(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = bt.TransferMiddleLeftRightBad(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	err = bt.TransferMiddleLeftRight(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = bt.TransferLeftRightMiddleBad(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	err = bt.TransferLeftRightMiddle(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = bt.TransferWidthFirst(&commonBT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
