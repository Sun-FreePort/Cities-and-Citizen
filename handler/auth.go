package handler

type Handler struct {
	UserId uint
}

// NewHandler 注册全局变量
func NewHandler() *Handler {
	return &Handler{}
}

// 鉴权

// 加入城市
