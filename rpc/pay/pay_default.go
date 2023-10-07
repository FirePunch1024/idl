package pay

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	pay "pay/kitex_gen/pay"
)

func Notify(ctx context.Context, request *pay.NotifyReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	resp, err = defaultClient.Notify(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Notify call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
