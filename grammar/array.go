package main

import "fmt"

func main() {
	var strings [5]string
	strings[0] = "Apple"
	strings[1] = "Banana"
	strings[2] = "Cat"
	strings[3] = "Dog"
	strings[4] = "Egg"
	fmt.Println("遍历strings：")
	for index,value := range strings{
		fmt.Println(index,value)
	}

	var five [5]int

	four := [4]int{1,2,3,4}

	fmt.Println("不同的数组：")
	fmt.Println(five)
	fmt.Println(four)

	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}
	for index,value := range six{
		fmt.Printf("Value[%s]\t Address[%p] IndexAddr[%p]\n",value,index,&six[index])
	}

}
