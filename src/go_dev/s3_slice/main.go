package main

import (
	"go_dev/s3_slice/statInit"
)

func main() {
	statInit.MakeStatSlice()


	statInit.WordStatSlice()

	statInit.DiffArraySlice()
}