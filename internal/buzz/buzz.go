package buzz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreater(ctx context.Context, greeter *Greeter) error
	UpdateGreeter(ctx context.Context, greeter *Greeter) error
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger *log.Helper) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: logger}
}
