package reply

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"testing"
)

func BenchmarkSReply_ListForAid(b *testing.B) {
	r := &sReply{}
	ctx := context.Background()
	r.ListForAid(ctx, 1)
}
