package qxMiddleware

import (
	"context"
	"errors"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxSony"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// note: 基于grpc的中间件，实现读取metadata中的信息映射到context中

func StreamHeaderParseInterceptor() grpc.StreamServerInterceptor {
	return func(svr any, stream grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		ctx := stream.Context()

		result := &Resp{
			Code: qxCodes.QxEngineStatusOK,
			Msg:  qxCodes.StatusText(qxCodes.QxEngineStatusOK),
			Path: info.FullMethod,
		}

		ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)

		// note: metadata中尝试获取requestId, 如果不存在就生成一个
		tempMD, isExist := metadata.FromIncomingContext(ctx)
		if !isExist {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return errors.New(qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata))
		}

		requestId := tempMD.Get(qxCommonHeader.HeaderXRequestIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxRequestID, requestId[0])
			result.RequestID = requestId[0]
		} else {
			tempRequestId := qxSony.NextId()
			ctx = context.WithValue(ctx, CtxRequestID, tempRequestId)
			result.RequestID = tempRequestId
		}

		//note: 读取metadata中的信息
		xAccessKeyFor := tempMD.Get(qxCommonHeader.HeaderXAccessKeyFor)
		if len(xAccessKeyFor) > 0 {
			ctx = context.WithValue(ctx, CtxXAccessKeyFor, xAccessKeyFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return errors.New(qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata))
		}

		//note: 读取metadata中的信息
		xTenantIDFor := tempMD.Get(qxCommonHeader.HeaderXTenantIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxTenantId, xTenantIDFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return errors.New(qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata))
		}

		xDomainIdFor := tempMD.Get(qxCommonHeader.HeaderXDomainIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxDomainId, xDomainIdFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return errors.New(qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata))
		}

		return handler(svr, stream)
	}
}

func UnaryHeaderParseInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {

		result := &Resp{
			Code: qxCodes.QxEngineStatusOK,
			Msg:  qxCodes.StatusText(qxCodes.QxEngineStatusOK),
			Path: info.FullMethod,
		}

		ctx = context.WithValue(ctx, CtxFullMethod, info.FullMethod)

		// note: metadata中尝试获取requestId, 如果不存在就生成一个
		tempMD, isExist := metadata.FromIncomingContext(ctx)
		if !isExist {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return result, nil
		}

		requestId := tempMD.Get(qxCommonHeader.HeaderXRequestIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxRequestID, requestId[0])
			result.RequestID = requestId[0]
		} else {
			tempRequestId := qxSony.NextId()
			ctx = context.WithValue(ctx, CtxRequestID, tempRequestId)
			result.RequestID = tempRequestId
		}

		//note: 读取metadata中的信息
		xAccessKeyFor := tempMD.Get(qxCommonHeader.HeaderXAccessKeyFor)
		if len(xAccessKeyFor) > 0 {
			ctx = context.WithValue(ctx, CtxXAccessKeyFor, xAccessKeyFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return result, nil
		}

		xTenantIDFor := tempMD.Get(qxCommonHeader.HeaderXTenantIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxTenantId, xTenantIDFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return result, nil
		}

		xDomainIdFor := tempMD.Get(qxCommonHeader.HeaderXDomainIDFor)
		if len(requestId) > 0 {
			ctx = context.WithValue(ctx, CtxDomainId, xDomainIdFor[0])
		} else {
			result.Code = qxCodes.QxEngineStatusNotFoundMetadata
			result.Msg = qxCodes.StatusText(qxCodes.QxEngineStatusNotFoundMetadata)
			return result, nil
		}

		return handler(ctx, req)
	}
}
