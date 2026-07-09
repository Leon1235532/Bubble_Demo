package schemas

type IDsPara struct {
	IDs []uint64 `json:"ids" binding:"required,min=1,dive,min=1"`
}
