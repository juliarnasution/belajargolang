package src

import "fmt"

//Array tes switch case
func Array(){
	/* local variable definition */
	var angka [5] int32//array
	angka[0] = 2
	angka[1] = 5
	fmt.Println(angka)
	var arr = [] float32{1,2,3,4}//slice
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(angka[0])
	//len adalah fungsi menghitung panjang array
	for j := 0; j < len(arr); j++ {
		fmt.Printf("Element[%d] = %f\n", j, arr[j] )
	 }

	//  multidimension array
	var nomor [5][2] int32
	nomor[0][0] = 10
	nomor[0][1] = 11
	fmt.Println(nomor[0][0])
	fmt.Println(nomor[0][1])

	a := [3][4] int{  
		{0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
		{4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
		{8, 9, 10, 11}}   /*  initializers for row indexed by 2 */
	
	fmt.Println(a)
	param := [10]int {1,2,3,4,5,6,7,8,9,10}
	myFunction1(param)
	param2 := []int {1,2,3}
	myFunction2(param2)
}

// function parameter with array
func myFunction1(param [10]int){
	fmt.Println(param)
}

func myFunction2(param []int){
	fmt.Println(param)
}