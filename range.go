package main
import(
"fmt"
"os"
)
func main()
{

s,sep:=","
for_, arg := range os.Args[1:]
{s += sep + arg
sep = ""
}
fmt.Println(s)
}
