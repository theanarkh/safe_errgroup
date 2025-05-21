package main

import (
	"context"
	"fmt"

	"github.com/theanarkh/safe_errgroup"
)

func main() {
	var eg safe_errgroup.ErrGroup
	eg.SafeGo(context.Background(), func() error {
		panic("oops")
	})
	err := eg.Wait()
	fmt.Println(err)
}
