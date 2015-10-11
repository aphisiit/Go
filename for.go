package main 
import "fmt"
func main() {
	var j int
	fmt.Print("Enter a number : ")
	fmt.Scan(&j)
	var t string
	fmt.Print("Enter a string : ")
	fmt.Scan(&t)	
	fmt.Printf("Number in 'j' = %d\n",j)
	for i := 0 ; i <= j ; i++{
		fmt.Println(t)
	}
}