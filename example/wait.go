package main

import (
	"context"
	"fmt"

	"github.com/theanarkh/safe_errgroup"
)

func main() {
	handler := safe_errgroup.WithHandler(func(_ context.Context, err *error) {
		if e := recover(); e != nil {
			if err != nil {
				fmt.Println("custom panic handler")
				panic(e)
			}
		}
	})
	eg := safe_errgroup.New(handler)
	eg.SafeGo(context.Background(), func() error {
		panic("panic should be recovered")
	})
	eg.SafeWait()
	fmt.Println("should run here")
}
