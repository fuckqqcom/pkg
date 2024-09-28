package metadata

import (
	"context"
)

func GetValFromCtx(ctx context.Context, key string) any {

	return ctx.Value(key)
}
