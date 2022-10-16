package main
import (
	"fmt"
	eq_solv "github.com/KA-7A/web_technologies/tree/task_4/calc/quadratic_equation"
)

func main() {
	var x1, x2 float64
	var code int
	
	var a, b, c float64
	a = 1
	b = 2
	c = 1
	x1, x2, code = eq_solv.Eq_Solver(a, b, c)
	fmt.Println("Сегодня мы будем решать уравнение\n ", a, "x^2 + ", b, "x + ", c, "= 0")
	if code == -1 {
		fmt.Println("Извините, вам не повезло, потому что я не собираюсь это решать. Попробуйте позже.")
	} else if code == 0{
		fmt.Println("Действительных корней нет")
	} else if code == 1{
		fmt.Println("Корень: x = ", x1)
	} else {
		fmt.Println("Корни: \nx1 = ", x1, "\nx2 = ", x2)
	}

}