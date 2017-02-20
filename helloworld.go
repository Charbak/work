package main
import "fmt"

func main() {
//var name string = "Charbak"
//var temp float64= 35
const pi float64=3.14
for i:=1;i<5;i++{
fmt.Println(i)
}
todaysTemp:=29
if todaysTemp < 30{
fmt.Println("Weather is pleasant")
}else{
fmt.Println("warm Today")
}
}