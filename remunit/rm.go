package remunit

import (
	"fmt"
	"strconv"
)

func RemToPx(remValue string, rem string){
	floatValue, err := strconv.ParseFloat(remValue, 64) 
	if err != nil{
		fmt.Println(err)
	}
	sum :=  floatValue * 16.0000
	if sum > 99.0 {
		fmt.Printf("%srem is equal to %.0fpx \n", rem, sum)
	}else{
		fmt.Printf("%srem is equal to %.0fpx \n", rem, sum)	
	}
	
}