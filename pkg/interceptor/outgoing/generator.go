package outgoing

import (
	"context"
	"strings"

	"github.com/v3nooonn/trytry/pkg/constant/types"
	"github.com/v3nooonn/trytry/pkg/utils/contextx"
	"google.golang.org/grpc/metadata"
)

type MDGenerator struct {
	context.Context
	MD metadata.MD
}

func defaultMD(ctx context.Context) metadata.MD {
	return metadata.New(map[string]string{
		strings.ToLower(types.CtxKeyTenantSchema.Key()): contextx.GetSchema(ctx),
	})
}

func EmptyGenerator(ctx context.Context) *MDGenerator {
	return &MDGenerator{
		Context: ctx,
		MD:      metadata.New(map[string]string{}),
	}
}

func NewMDGenerator(ctx context.Context) *MDGenerator {
	return &MDGenerator{
		Context: ctx,
		MD:      defaultMD(ctx),
	}
}

func (g *MDGenerator) set(md metadata.MD) *MDGenerator {
	g.MD = md
	return g
}

func (g *MDGenerator) context() context.Context {
	return g.Context
}

func (g *MDGenerator) Metadata() metadata.MD {
	return g.MD
}

func (g *MDGenerator) WithTenant() *MDGenerator {
	md := g.Metadata()

	md[strings.ToLower(types.CtxKeyTenantSchema.Key())] = append(
		md[strings.ToLower(types.CtxKeyTenantSchema.Key())],
		contextx.GetSchema(g.context()),
	)
	g.set(md)

	return g
}
