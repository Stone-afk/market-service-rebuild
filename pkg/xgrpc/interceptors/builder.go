package interceptors

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

type Builder struct {
}

func (b *Builder) PeerName(ctx context.Context) string {
	return b.grpcHeaderValue(ctx, "app")
}

func (b *Builder) PeerIP(ctx context.Context) string {
	clientIP := b.grpcHeaderValue(ctx, "client-ip")
	if clientIP != "" {
		return clientIP
	}
	pr, ok2 := peer.FromContext(ctx)
	if !ok2 {
		return ""
	}
	if pr.Addr == net.Addr(nil) {
		return ""
	}
	addSlice := strings.Split(pr.Addr.String(), ":")
	if len(addSlice) > 1 {
		return addSlice[0]
	}
	return ""
}

func (b *Builder) grpcHeaderValue(ctx context.Context, key string) string {
	if key == "" {
		return ""
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	return strings.Join(md.Get(key), ";")
}
