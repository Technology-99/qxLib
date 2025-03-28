package qxMiddleware

import (
	"context"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxSony"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"regexp"
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
		start := time.Now().UnixMilli()
		ctx := context.WithValue(r.Context(), CtxFullMethod, r.URL.Path)
		ctx = context.WithValue(ctx, CtxRequestURI, r.RequestURI)
		ctx = context.WithValue(ctx, CtxStartTime, start)
		fullAddr := httpx.GetRemoteAddr(r)
		// 定义正则表达式
		regex := `^\[?([a-fA-F0-9:.%]+)\]?(?::([0-9]+))?$`
		re := regexp.MustCompile(regex)
		// 去除首尾空格
		fullAddr = strings.TrimSpace(fullAddr)
		// 匹配输入字符串
		matches := re.FindStringSubmatch(fullAddr)
		// 匹配输入字符串
		if len(matches) != 3 {
			fullAddrAndPort := strings.Split(fullAddr, ":")
			if len(fullAddrAndPort) == 1 {
				logx.Infof("client ip: %s", fullAddrAndPort[0])
				ctx = context.WithValue(ctx, CtxClientIp, fullAddrAndPort[0])
				ctx = context.WithValue(ctx, CtxClientPort, "")
			} else if len(fullAddrAndPort) == 2 {
				logx.Infof("client ip: %s", fullAddrAndPort[0])
				logx.Infof("client port: %s", fullAddrAndPort[1])
				ctx = context.WithValue(ctx, CtxClientIp, fullAddrAndPort[0])
				ctx = context.WithValue(ctx, CtxClientPort, fullAddrAndPort[1])
			} else {
				logx.Errorf("fullAddr: %s, fullAddrAndPort: %v", fullAddr, fullAddrAndPort)
			}
		} else {
			clientIp := matches[1]
			clientPort := matches[2]
			ctx = context.WithValue(ctx, CtxClientPort, clientPort)
			if matches[1] == "::1" {
				clientIp = "127.0.0.1"
			}
			ctx = context.WithValue(ctx, CtxClientIp, clientIp)
			logx.Infof("client ip: %s", clientIp)
			logx.Infof("client port: %s", clientPort)
		}

		requestID := ""
		xRequestIDFor := r.Header.Get(qxCommonHeader.HeaderXRequestIDFor)
		if xRequestIDFor == "" {
			requestID = qxSony.NextId()
			ctx = context.WithValue(ctx, CtxRequestID, requestID)
		} else {
			requestID = xRequestIDFor
		}
		ctx = context.WithValue(ctx, CtxRequestID, requestID)

		// 获取 User-Agent
		userAgent := r.Header.Get(CtxUserAgent)
		ctx = context.WithValue(ctx, CtxUserAgent, userAgent)

		xAccessKeyFor := r.Header.Get(qxCommonHeader.HeaderXAccessKeyFor)
		xAuthMethodFor := r.Header.Get(qxCommonHeader.HeaderXAuthMethodFor)
		xAccountFor := r.Header.Get(qxCommonHeader.HeaderXAccessKeyFor)
		ctx = context.WithValue(ctx, CtxXAccessKeyFor, xAccessKeyFor)
		ctx = context.WithValue(ctx, CtxXAuthMethodFor, xAuthMethodFor)
		ctx = context.WithValue(ctx, CtxXAccountFor, xAccountFor)

		//ctx = context.WithValue(ctx, "clientPort", fullAddrAndPort[1])
		r = r.WithContext(ctx)
		next(w, r)
	}
}
