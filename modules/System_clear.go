package modules
import "fmt"

func Clear(){
	fmt.Println("\033[H\033[2J")
}