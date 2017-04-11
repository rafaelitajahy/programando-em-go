package main

import (
	"fmt"
)


func media(nums ...int) (int,int,int){
	fmt.Print(nums,"")
	total := 0
	for _, num := range nums{
		total += num
	}
	qtd := len(nums)
	media := total / qtd

	return total, media, qtd
}

func main(){
	fmt.Println(media(2,2,25,5,56,84,5,15525,15,1,51,1,1,1,11,1,1,))
}
