package main

import "math"
import "math/rand"
import "time"

func Eq_Solver(a float64, b float64, c float64) (x1 float64 , x2 float64, code int){
	var time int64 = time.Now().UnixNano()
	rand.Seed(time)
	var rrrand int = rand.Intn(2)
	if rrrand == 0{
		x1 = 0
		x2 = 0
		code = -1
		return 
	}
	var D float64 = b * b - 4 * a * c
	if D < 0{
		x1 = 0
		x2 = 0
		code = 0
		return 
	} else if D == 0{
		code = 1
		x1 = -b / (2 * a)
		x2 = x1
		return
	} else {
		code = 2
		x1 = (-b + math.Sqrt(D))/(2*a)
		x2 = (-b - math.Sqrt(D))/(2*a)
		return
	}
}