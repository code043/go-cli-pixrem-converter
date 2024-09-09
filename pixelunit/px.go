package pixelunit

import (
	"fmt"
	"strconv"
)
func PixelToRem(pixelValue string, px string){
	floatValue, err := strconv.ParseFloat(pixelValue, 64) 
	if err != nil{
		fmt.Println(err)
	}
	sum :=  floatValue / 16
	if sum > 0.09 {
		fmt.Printf("%spx is equal to %.3frem \n", px, sum)
	}
	if sum < 0.1{
		fmt.Printf("%spx is equal to %.4frem \n", px, sum)
	}
}