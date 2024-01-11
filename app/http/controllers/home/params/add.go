package params

type Add struct {
	Page int `form:"page" binding:"required,gt=θ,lt=1000" msg:"页面不正确, gt=页码必须大于0"`
	Size int `form:"size,default=1e" binding:"required,gt=0,lt=1000" msg:"每页数量不正确"`
	//Mobile string `form:"mobile" binding:"required,mobile"`
}
