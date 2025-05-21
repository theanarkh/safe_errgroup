package safe_errgroup

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type ErrGroup struct {
	errgroup.Group
	cancel  func(error)
	handler func(context.Context, *error)
}

type Option func(*ErrGroup)

func WithHandler(handler func(context.Context, *error)) func(e *ErrGroup) {
	return func(e *ErrGroup) {
		e.handler = handler
	}
}

func New(options ...Option) *ErrGroup {
	eg := &ErrGroup{}
	for _, option := range options {
		option(eg)
	}
	return eg
}

func handler(_ context.Context, err *error) {
	if e := recover(); e != nil {
		if err != nil {
			*err = fmt.Errorf("panic happen: %v", e)
		}
	}
}

func WithContext(ctx context.Context) (*ErrGroup, context.Context) {
	ctx, cancel := context.WithCancelCause(ctx)
	return &ErrGroup{cancel: cancel}, ctx
}

func (e *ErrGroup) SafeGo(ctx context.Context, f func() error) {
	e.Go(func() (err error) {
		if e.handler != nil {
			defer e.handler(ctx, &err)
		} else {
			defer handler(ctx, &err)
		}
		return f()
	})
}

func (e *ErrGroup) SafeTryGo(ctx context.Context, f func() error) bool {
	return e.TryGo(func() (err error) {
		if e.handler != nil {
			defer e.handler(ctx, &err)
		} else {
			defer handler(ctx, &err)
		}
		return f()
	})
}

func (e *ErrGroup) SafeWait() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic happen: %v", e)
		}
	}()
	err = e.Wait()
	if e.cancel != nil {
		e.cancel(err)
	}
	return err
}
