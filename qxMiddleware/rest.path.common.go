package qxMiddleware

import (
	"context"
	"fmt"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net"
	"net/http"
	"strings"
	"time"
)

type PathHttpInterceptorMiddleware struct {
}

func NewPathHttpInterceptorMiddleware() *PathHttpInterceptorMiddleware {
	return &PathHttpInterceptorMiddleware{}
}

func (m *PathHttpInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ctx := context.WithValue(r.Context(), CtxFullMethod, r.URL.Path)
		ctx = context.WithValue(ctx, CtxRequestURI, r.RequestURI)
		ctx = context.WithValue(ctx, CtxStartTime, startTime.UnixMilli())
		fullAddr := httpx.GetRemoteAddr(r)
		ips := strings.Split(fullAddr, ",")
		realAddr := ips[0]
		ip := ""
		port := ""
		var err error
		if ipType(realAddr) == "IPv4" {
			if strings.Contains(realAddr, ":") {
				// note: 带端口号的ip
				ip, port, err = net.SplitHostPort(realAddr)
			} else {
				// note: 不带端口号的ip
				ip = realAddr
				port = ""
			}
			ctx = context.WithValue(ctx, CtxClientIp, ip)
			ctx = context.WithValue(ctx, CtxClientPort, port)
		} else {
			ip = fmt.Sprintf("%s", net.ParseIP(realAddr))
			ctx = context.WithValue(ctx, CtxClientIp, ip)
			ctx = context.WithValue(ctx, CtxClientPort, "")
		}
		logc.Infof(ctx, "IP: %s, Port: %s", ip, port)
		if err != nil {
			logx.Infof("解析ip报错: %s", err)
			http.Error(w, "不支持的ip类型", http.StatusNotImplemented)
			return
		}

		requestID := uuid.NewString()
		ctx = context.WithValue(ctx, CtxRequestID, requestID)
		w.Header().Set(qxCommonHeader.HeaderXRequestIDFor, requestID)

		// 获取 User-Agent
		userAgent := r.Header.Get(CtxUserAgent)
		ctx = context.WithValue(ctx, CtxUserAgent, userAgent)

		endTime := time.Now()
		logc.Infof(ctx, "路径ip处理中间件耗时: %v", endTime.Sub(startTime).Milliseconds())

		r = r.WithContext(ctx)
		next(w, r)
	}
}

func ipType(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "invalid"
	}
	if ip.To4() != nil {
		return "IPv4"
	}
	return "IPv6"
}
