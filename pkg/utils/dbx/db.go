package dbx

import (
	"fmt"
)

//func Set(ctx context.Context, cachePrefix string) string {
//	return fmt.Sprintf(cachePrefix, ctxutil.GetTenant(ctx))
//}

func SchInit(format, schema string) string {
	return fmt.Sprintf(format, schema)
}
