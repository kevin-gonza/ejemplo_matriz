package main 

import (
	"fmt"
)
// creamos la matriz

type Nums [][]int 

// sumamos todos los valores 
 func (nums Nums) Sumar() int {
     
	total:=0
    
	for i := 0; i < len(nums); i++ {
		for j := 0; j< len(nums[i];  j++)
		{
			total += nums[i][j]
		}
		
	}
    return total

 }
func main () {
    
	nums := Nums{
		[]int{2, 3, 4}
		[]int{6, 5, 7}
		[]int{7, 8, 2}
	}
	 fmt.println(nums)
	 fmt.Println("la suma de todos los valores es: " nums.Sumar())
  	
 }