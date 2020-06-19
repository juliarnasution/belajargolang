package src
import "fmt"

// IfElse contoh looping
func IfElse(){
	//range
	nums :=[]int {1,2,3,4,5,6,7,8,9,10}// array
	sum :=0
	// firts var is index, second var is value
	for _, num := range nums {//_ cannot use as value, just write _ underscore if not used
		if num % 2 ==0 {
			sum += num
			fmt.Println("add value ", num)
		}else if num == 5{
			continue
		}else{
			fmt.Println("skipping value ", num)
		}
	}
	
	fmt.Println("sum ",sum)
}