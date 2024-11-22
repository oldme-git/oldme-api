// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
)

type IReadingV1 interface {
	Reading(ctx context.Context, req *v1.ReadingReq) (res *v1.ReadingRes, err error)
}
