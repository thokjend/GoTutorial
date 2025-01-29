package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":32, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", i,v)
	}





	/* var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator,denominator)
	if err != nil{
		fmt.Println(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v", result)
	}else{
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	} */
}

/* func printMe(printValue string) {
	fmt.Println(printValue)
} */

/* func intDivision(numerator int, denominator int) (int,int, error){
	var err error
	if denominator == 0{
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
} */