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
				*err = fmt.Errorf("custom panic handler: %v", e)
			}
		}
	})
	eg := safe_errgroup.New(handler)
	eg.SafeGo(context.Background(), func() error {
		panic("panic should be recovered")
	})
	err := eg.SafeWait()
	fmt.Println(err)
}
