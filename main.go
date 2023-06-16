package main

import "fmt"


func binarySearch(array []int, checkVal int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == checkVal {
			return mid
		} else if array[mid] < checkVal {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main(){
	slice:=[]int {1,2,3,4,5,6,7,9,10,11,15,18,20}
	checkVal:=18

	position:=binarySearch(slice,checkVal)

	if position != -1{
		fmt.Printf("%v found at position %v ",checkVal,position+1)
	}else{
		fmt.Println("value not found")
	}
	
	
}
