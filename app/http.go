package app

type BaseController struct{}

type Request struct {
	QueryString string
	Body        map[string]any
}
