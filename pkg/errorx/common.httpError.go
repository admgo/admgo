package errorx

var (
	OK                        = NewHttpErr(0, "ok")
	NoLogin                   = NewHttpErrAndSetHttpCode(401, 10001, "no login")
	RequestErr                = NewHttpErr(40001, "INVALID_ARGUMENT")
	Unauthorized              = NewHttpErr(40002, "UNAUTHENTICATED")
	AccessDenied              = NewHttpErr(40003, "PERMISSION_DENIED")
	NotFound                  = NewHttpErr(40004, "not found")
	MethodNotAllowed          = NewHttpErr(40005, "METHOD_NOT_ALLOWED")
	ServerErr                 = NewHttpErrAndSetHttpCode(500, 50001, "Internal error")
	ServerErrForRedis         = NewHttpErrAndSetHttpCode(500, 50002, "Internal error")
	ServerErrForMysql         = NewHttpErrAndSetHttpCode(500, 50003, "Internal error")
	ServerErrForElasticsearch = NewHttpErrAndSetHttpCode(500, 50004, "Internal error")
	ServiceUnavailable        = NewHttpErrAndSetHttpCode(500, 50020, "Service Unavailable")
	HttpDeadline              = NewHttpErr(50021, "deadline exceeded")
	LimitExceed               = NewHttpErr(50022, "RESOURCE_EXHAUSTED")
)
