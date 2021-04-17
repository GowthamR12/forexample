package main
import "fmt"
const(
	 a=20+iota
	 b=20+iota
)
func main(){
	fmt.Printf("%d\\t%b\\n",a,a)
	f:=10
	c:=f<<1
	d:=f>>1
	fmt.Printf("%d\t%b\n",c,c)
	fmt.Printf("%d\t%b\n",d,d)
	fmt.Printf("%d\t%b",b,b)
}