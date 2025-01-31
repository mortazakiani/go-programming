package main

import (
	"fmt"
	"sync"
)

//mutex  made for  stabel  and lock a part  of code for go routines to  avoide miss access and conflict between  go routines
var x =0
func sum(wg *sync.WaitGroup,m *sync.Mutex) {
    m.Lock()//here we lock  the x value and let one go routtines come 
    x+=1
    m.Unlock()//heer  we unlock  x value  then next go routine come
    wg.Done()
}


func main() {

	var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 10000; i++ {
        w.Add(1)
        go sum(&w,&m)
    }
    w.Wait()
    fmt.Println("final", x)

}
