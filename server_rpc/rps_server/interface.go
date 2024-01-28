package rpcserver

type RPCServer interface {
	StartServer(port string, rcvr ...any) error
}

type Address struct {
    GeoLat string `json:"lat"`
    GeoLon string `json:"lon"`
    Result string `json:"result"`
}