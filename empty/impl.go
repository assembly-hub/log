package empty

import (
	"context"

	"github.com/assembly-hub/log"
)

type empty struct {
}

func (e empty) Debug(ctx context.Context, msg string) {

}

func (e empty) Trace(ctx context.Context, msg string) {

}

func (e empty) Notice(ctx context.Context, msg string) {

}

func (e empty) Warning(ctx context.Context, msg string) {

}

func (e empty) Error(ctx context.Context, msg string) {

}

func (e empty) Fatal(ctx context.Context, msg string) {

}

var NoLog log.Log = empty{}
