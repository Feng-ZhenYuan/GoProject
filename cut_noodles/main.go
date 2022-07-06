package main
import "fmt"
func cut_noodles (times int) int {
	result := 2
	t:=1
	for i:=0;i<times;i++ {
		result += t
		t = t*2
	}
	return result
}

func main(){
	var result int
	times := 10
	result = cut_noodles(times)
	fmt.Printf("对折%d次，从中间切一刀得到的面条数是：%d\n",times,result)
}
