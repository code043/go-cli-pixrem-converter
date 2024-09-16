package pixelunit

import (
	"fmt"
	"strconv"
)
func PixelToRem(px string){
	floatValue, err := strconv.ParseFloat(px, 64) 
	if err != nil{
		fmt.Println(err)
	}
	sum :=  floatValue / 16.000
	if sum >= 1.0 {
		fmt.Printf("%spx is equal to %.2frem \n", px, sum)
	}
	if sum < 1.0{
		fmt.Printf("%spx is equal to %.4frem \n", px, sum)
	}
}