package proto

import "go_project_template/application/appA/repo"

type GetUserReq struct {
	Uid int `json:"uid"`
}

// 引用repo中的model struct
type GetUserRsp struct {
	Core  *repo.UserCore  `json:"base"`
	Other *repo.UserOther `json:"other"`
}
