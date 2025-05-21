package qxMiddleware

import (
	"context"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxErrors"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net"
	"net/http"
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
		ip, port, err := net.SplitHostPort(fullAddr)
		if err != nil {
			qxRes.JsonBaseResponseCtx(r.Context(), w, r, nil, qxErrors.Quick(qxCodes.StatusNotImplemented, qxCodes.LangZhCN))
			return
		}
		logx.Infof("ip: %s, port: %s", ip, port)

		ctx = context.WithValue(ctx, CtxClientIp, ip)
		ctx = context.WithValue(ctx, CtxClientPort, port)

		//// 定义正则表达式
		//regex := `^\[?([a-fA-F0-9:.%]+)\]?(?::([0-9]+))?$`
		//re := regexp.MustCompile(regex)
		//// 去除首尾空格
		//fullAddr = strings.TrimSpace(fullAddr)
		//// 匹配输入字符串
		//matches := re.FindStringSubmatch(fullAddr)
		//// 匹配输入字符串
		//if len(matches) != 3 {
		//	fullAddrAndPort := strings.Split(fullAddr, ":")
		//	if len(fullAddrAndPort) == 1 {
		//		ctx = context.WithValue(ctx, CtxClientIp, fullAddrAndPort[0])
		//		ctx = context.WithValue(ctx, CtxClientPort, "")
		//	} else if len(fullAddrAndPort) == 2 {
		//		ctx = context.WithValue(ctx, CtxClientIp, fullAddrAndPort[0])
		//		ctx = context.WithValue(ctx, CtxClientPort, fullAddrAndPort[1])
		//	} else {
		//		logx.Errorf("fullAddr: %s, fullAddrAndPort: %v", fullAddr, fullAddrAndPort)
		//	}
		//} else {
		//	clientIp := matches[1]
		//	clientPort := matches[2]
		//	ctx = context.WithValue(ctx, CtxClientPort, clientPort)
		//	if matches[1] == "::1" {
		//		clientIp = "127.0.0.1"
		//	}
		//	ctx = context.WithValue(ctx, CtxClientIp, clientIp)
		//}

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
