package main
import ("fmt")


func main(){

  var (
    a int = 434
    c = "hello"
    d float32
  )

  fmt.Print(a, " ", "Hello World");

  fmt.Println(a)
  // fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  fmt.Printf("The value of v is %v and the type is %T \n", a,a);
}