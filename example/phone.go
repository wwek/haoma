package main

import (
	"fmt"

	"github.com/wwek/haoma/libs/phone"
)

func main() {
	p := phone.New("053266114000")
	// p := phone.New("18870208731")
	// p := phone.New("051288178411")
	//053266114000
	//051288178411
	//18870208731
	//15710296540
	//18613846501
	//18887174942
	//15722051740
	//15710235224
	//13111871972
	//18622929857
	//15722235112
	//15710296549
	//15710235112
	//15760738024
	//18302562964
	//18630857593
	//15722235014
	//15710235216
	//15722291226
	//13263387512
	//18630879302

	result, _ := p.QueryAll()
	fmt.Println(result)
	for _, r1 := range result {
		//fmt.Println(n)
		fmt.Println(r1)
	}

	// result, _ := p.Query_sogouhaomatong()
	// fmt.Println(result)
	// result1, _ := p.Query_360shoujiweishi()
	// fmt.Println(result1)
	// result, _ = p.Query_baidushoujiweishi()
	// fmt.Println(result)

}
