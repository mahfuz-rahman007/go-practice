package main
import ("fmt")


func main(){

var a = [2]int{1,2};

var b = [3]string{"a", "b", "c"};

b[2] = "bullshit"; // element update

d := [...]int{1,2,3,4,5,6}; // no length added

g := [4]int{}; // not initialized

l := [4]int{1,3}; // partially initialized

m := [3]string{"a"} // partially initialized

x := [5]string{1:"hello", 4:"ma"};

fmt.Println(a);
fmt.Println(b[1]);
fmt.Println(b);
fmt.Println(d);

fmt.Println(g);
fmt.Println(l);

fmt.Println(m);

fmt.Println(len(x)); // length of array



}