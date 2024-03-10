package main

import (
	"fmt"
	"sync"
	//"sync"
	//"time"
)

// func printnum(wg *sync.WaitGroup){
// 	wg.Done()
// 	for i:=1;i<=10;i++{
// 		fmt.Println(i)
// 		// time.Sleep(200*time.Millisecond)
// 	}
// }

// func printchar(wg *sync.WaitGroup){
// 	wg.Done()
// 	for i:=65;i<=75;i++{
// 		fmt.Println(string(i))
// 		// time.Sleep(100*time.Millisecond)
// 	}
// }

// func sum(a, b int,ch chan int) {
// 	s := 0
// 	for i := a; i <= b; i++ {
// 		s += i
// 	}
// 	ch <- s
// }


// func incre(a *int,mu *sync.Mutex,wg *sync.WaitGroup){
// 	// defer return
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	*a++
// }

// func decre(a *int,mu *sync.Mutex,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	*a--

// }
// func main(){
// 	a:=8
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go incre(&a, &mu, &wg)

// 	wg.Add(1)
// 	go decre(&a, &mu, &wg)
// 	fmt.Println(a)
// 	wg.Wait()







		// ch := make(chan int)
		// go sum(1,100,ch)
		// output := <-ch
		// fmt.Println(output)




	// var wg sync.WaitGroup

	// wg.Add(1)
	// go printnum(&wg)

	// wg.Add(1)
	// go printchar(&wg)
	// // time.Sleep(1500*time.Millisecond )


	// wg.Wait()
}





/*     Mutexes -   lock unlock mechanism
					make sure that only one go rouitie can access the shared resource at a time




*/