namespace go index

include "base.thrift"

struct PingReq {
    1: string ping (api.path="ping");
}

struct PingResp {
    1: base.Base base,
    2: map<string, string> data;
}

service PingService {
    PingResp PingServer(1: PingReq request) (api.get="/ping", api.group="ping");
}
