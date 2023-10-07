namespace go pay

//购买商品
struct BuyGoodsReq{
    1: string Subject
}

struct CallBackReq{}

struct NotifyReq{}

struct NilResponse {}

service payService {

    NilResponse Notify(1: NotifyReq request) (api.post="/alipay/notify")
}

