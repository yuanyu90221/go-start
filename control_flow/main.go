package main

import "fmt"

func main(){
	// continueLoop()
	// breakExample()
	switchDemo()

}

func continueLoop() {
	for i := 0; i < 10; i++ {
		if i%2 ==0 {
			continue
		}
		fmt.Println(i)
	}
}

func breakExample() {
	flag := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		}
		fmt.Println(i)
	}
	fmt.Println(flag)
}

func switchDemo() {
	day := "Fri"
	switch day {
		case "Fri":
			fmt.Println("TGIF")
		case "Mon", "Tue", "Wed":
			fmt.Println("Boring")
		default: 
			fmt.Println("default")
	}
	switch {
		case  day == "Fri":
			fmt.Println("TGIF")
			break
	}
}