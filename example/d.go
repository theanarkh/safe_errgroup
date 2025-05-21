package main

import "golang.org/x/sync/errgroup"

func demo() {
	defer func() {
		if err := recover(); err != nil {
			println("recover")
		}
	}()
	var g errgroup.Group
	g.Go(func() error {
		panic("hello")
	})
	//g.Wait()
}
func main() {
	demo()
	// var once sync.Once
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	once.Do(func() {
	// 		wg.Done()
	// 		println("hello")
	// 	})
	// }()
	// wg.Wait()
}
