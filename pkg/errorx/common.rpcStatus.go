package errorx

var (
	RpcCanceled = NewRpcStatus(498, "Canceled")
	RpcDeadline = NewRpcStatus(504, "deadline exceeded")
)
