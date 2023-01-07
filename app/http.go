package app

type BaseController struct{}

type Request struct {
	QueryString string
	Body        map[string]any
}

// 框架 -> 实现(业务代码)
