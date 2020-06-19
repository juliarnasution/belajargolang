package src
import "fmt"

// ForLoop contoh looping
func ForLoop(){
	for i := 0; i < 10; i++ {
		fmt.Println("for biasa ",i)
	}
	angka :=0
	for angka < 6 {
		fmt.Println("for syarat ",angka)
		angka ++
	}

	//range
	nums :=[]int {1,2,3,4,5}// array
	sum :=0
	// firts var is index, second var is value
	for i, num := range nums {//_ cannot use as value
		fmt.Println("index ",i)
		sum += num
	}
	
	fmt.Println("sum ",sum)
}