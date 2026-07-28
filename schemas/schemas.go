package schemas

type IDsPara struct {
	IDs []uint64 `binding:"required,min=1,dive,min=1" json:"ids"`
}

type Pagination struct {
	Page     uint64 `binding:"omitempty,min=1" json:"page"`
	PageSize uint64 `binding:"omitempty,min=3" json:"pagesize"`
}

type RegisterRequest struct {
	Username string `binding:"required,min=3,max=20" json:"username"`
	Password string `binding:"required,min=6,max=20" json:"password"`
}

type RegisterRespond struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type LoginRespond struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type PwdChange struct {
	OldPwd string `binding:"required" json:"oldpwd"`
	NewPwd string `binding:"required,min=6,max=20" json:"newpwd"`
}

type LogOutRequest struct {
	Pwd string `json:"pwd"`
}
