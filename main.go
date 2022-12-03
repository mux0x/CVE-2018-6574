package main
// #cgo CFLAGS: -fplugin=./attack.so
// typedef int (*intFunc) ();
// int _func(intFunc f)
// {
//      return f();
// }
// int one()
// {
//      return 1;
// }
import "C"
import "fmt"

func main() {
    f := C.intFunc(C.one)
    fmt.Println(int(C._func(f)));
}
