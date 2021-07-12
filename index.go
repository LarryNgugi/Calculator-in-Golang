
package main

import "fmt"

 func main(){
	 
	var x,y float32
	var operator string
	
    fmt.Println("Enter First value :")
	fmt.Scan(&x)
	fmt.Println("Enter Second value : ")
	fmt.Scan(&y)
    fmt.Println("Enter Operator(+ - * /)")
	fmt.Scan(&operator)


	switch operator{

	case "+" :

		fmt.Printf("%f,%s,%f=%f",x,operator,y,x+y)
		
	case "-" :

		fmt.Printf("%f,%s,%f=%f",x,operator,y,x-y)

	case "*" :

		fmt.Printf("%f,%s,%f=%f",x,operator,y,x*y)

	case "/" :
		if y == 0.0 {
			
			fmt.Println("Divide by Zero situation")
		}else {
			fmt.Printf("%f,%s,%f=%f",x,operator,y,x/y)
		}

	default:
		fmt.Println("Inavalid operator")


	}
 }
