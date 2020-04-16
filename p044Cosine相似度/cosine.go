package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Cosine(a []float64, b []float64) float64 {
	var (
		aLen  = len(a)
		bLen  = len(b)
		s     = 0.0
		sa    = 0.0
		sb    = 0.0
		count = 0
	)

	if aLen > bLen {
		count = aLen
	} else {
		count = bLen
	}

	for i := 0; i < count; i++ {
		if i >= bLen {
			sa += math.Pow(a[i], 2)
			continue
		}
		if i >= aLen {
			sb += math.Pow(b[i], 2)
			continue
		}
		s += a[i] * b[i]
		sa += math.Pow(a[i], 2)
		sb += math.Pow(b[i], 2)
	}
	return s / ((math.Sqrt(sa) * math.Sqrt(sb)) + 1e-6)
}

func durationTime(start int64, t string) int64 {
	end := time.Now().UnixNano()
	if t == "ms" {
		return (end - start) / int64(time.Millisecond)
	}
	return (end - start) / int64(time.Second)
}

func RandFloat64Slice(n int) []float64 {
	vec := make([]float64, n)
	for i := 0; i < n; i++ {
		vec[i] = rand.Float64()
	}
	return vec
}

func main() {
	a := []float64{0.6415715, -0.099079, 0.285402, 0.396016, 0.165099, 0.10195151}
	b := []float64{0.23921451, 0.217118, 0.030731, 0.426671, 0.18734, 0.28663248}
	c := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	score1 := Cosine(a, b)
	score2 := Cosine(a, c)
	fmt.Println(score1, "\n", score2)

	rand.Seed(time.Now().UnixNano())
	dim := 200
	v1 := RandFloat64Slice(dim)
	v2 := RandFloat64Slice(dim)
	//fmt.Println(v1)
	//fmt.Println(v2)

	start := time.Now().UnixNano()
	n := 1000000
	socreSlice := make([]float64, n)
	for i := 0; i < n; i++ {
		socreSlice[i] = Cosine(v1, v2)
	}

	fmt.Printf("Cost time %d s.\n", durationTime(start, "s"))
	fmt.Println(socreSlice[n-1])
}
