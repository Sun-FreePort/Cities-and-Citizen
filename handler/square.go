package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// SquareInfo 广场信息、城市信息
func (h *Handler) SquareInfo(c *fiber.Ctx) error {
	resp := &SquareInfoResp{}

	return c.Status(http.StatusOK).JSON(resp)
}

// 最近的演讲

// 发表演讲

// 表态
