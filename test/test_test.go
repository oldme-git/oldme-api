package test

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	_ "oldme-api/internal/logic/file"
	_ "oldme-api/internal/logic/rich"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	var (
		ctx = gctx.New()
	)
	//gtimer.AddTimes(ctx, time.Second, 10, func(ctx context.Context) {
	//	fmt.Println(gtime.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
	//	now = time.Now()
	//})
	fmt.Println(time.Now())
	gtimer.AddOnce(ctx, 20*time.Second, func(ctx context.Context) {
		fmt.Println(time.Now())
	})
	select {}
}
