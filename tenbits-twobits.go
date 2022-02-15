package main

import (
	"fmt"
	"strconv"
)
type Template struct {
	IntegerNu int
	CountNu int
	ss string
}
func (t *Template)convertToBin() string {
	fmt.Printf("Input= %v\n",t.IntegerNu)
	Number:=t.IntegerNu
	t.ss = ""

	if t.IntegerNu == 0 {
		return "0"
	}
	t.CountNu=0
	for ;t.IntegerNu > 0 ; t.IntegerNu /= 2 {

		lsb := t.IntegerNu % 2
		if lsb ==1 {
			t.CountNu+=1
		}
		t.ss = strconv.Itoa(lsb) + t.ss
	}

	fmt.Printf("Output Count =%v\n",t.CountNu)
	fmt.Printf("Details: %v =%v (%v set bits)\n",Number,string(t.ss),t.CountNu)
	return t.ss
}

func (t *Template) Count(s string){

}

func main(){
	t:=new(Template)
	fmt.Println( "Plesae input an integer Number:")
	fmt.Scanln(&t.IntegerNu)

	t.convertToBin()



}