package schemas

type IDsPara struct {
	IDs []uint64 `binding:"required,min=1,dive,min=1" json:"ids"`
}

type Pagination struct {
	Page     uint64 `binding:"omitempty,min=1" json:"page"`
	PageSize uint64 `binding:"omitempty,min=3" json:"pagesize"`
}
