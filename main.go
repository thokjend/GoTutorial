package main

import (
	"fmt"
)

/* type elecrticEngine struct{
	mpkwh uint8
	kwh uint8
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}
func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg
}

func (e elecrticEngine) milesLeft() uint8{
	return e.kwh*e.mpkwh
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
} */

/* var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{} */

/* func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}
func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}
func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
} */

/* func process(c chan int){
	defer close(c)
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Exiting process")
} */

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

func main() {
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))
	/* var c = make(chan int, 5)
	go process(c)
	for i:= range c{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	} */

	/* t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results) */



	/* var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	p = &i */

	/* var gasEngine gasEngine = gasEngine{25, 15}
	var elecrticEngine elecrticEngine = elecrticEngine{10, 15}
	canMakeIt(gasEngine, 1)
	canMakeIt(elecrticEngine, 1) */

	/* var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i,v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s","u","b","s","c","r","i","b","e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr) */

	/* var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n)) */

	/* intArr := [...]int32{1, 2, 3}
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
	for i:=0; i<10; i++{
		fmt.Println(i)
	} */

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

/* func timeLoop(slice []int, n int) time.Duration{
	var t0= time.Now()
	for len(slice) < n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
} */

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