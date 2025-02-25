package main

const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// General errors
	CodeGeneralError = 50000
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

const (
	// Shared
	Pending = 0

	//event
	Scheduled    = 1
	PastComplete = -1
	Delayed      = -2
	Cancelled    = -99

	// location
	Active               = 1
	TemporaryUnavailable = -1
	Closed               = -99

	// order
	Complete       = 2
	PayingStage    = 1
	OrderCancelled = -99

	// payment
	Paid   = 1
	Failed = -99
)
