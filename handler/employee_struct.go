package handler

type EmployeeMarketInfoReq struct {
}

type EmployeeMarketInfoResp struct {
}

type EmployeeWantAdListReq struct {
}

type EmployeeWantAdListResp struct {
}

type EmployeeWantAdSetReq struct {
}

type EmployeeWantAdSetResp struct {
}

type EmployeeWantAdCancelReq struct {
}

type EmployeeWantAdCancelResp struct {
}

type EmployeeJoinReq struct {
}

type EmployeeJoinResp struct {
}

type EmployeeFireReq struct {
	EmployeeId int64 `json:"employee_id"`
	Block      bool  `json:"block"` // 拉黑
}

type EmployeeFireResp struct {
}
