package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// SquareInfo godoc
//
// @Summary     获取信息(广场/城市)
// @Tags        Square/广场
// @Accept      json
// @Produce     json
// @Param       request body SquareInfoResp true "基本参数"
// @Success     200 {string} string "JSON 格式的存档数据"
// @Failure     403 {string} string "错误码"
// @Router      /square/info [post]
func (h *Handler) SquareInfo(c *fiber.Ctx) error {
	resp := &SquareInfoResp{}

	return c.Status(http.StatusOK).JSON(resp)
}

// 最近的演讲

// SquarePublish godoc
//
// @Summary     发表演讲
// @Tags        Square/广场
// @Accept      json
// @Produce     json
// @Param       request body PublishSpeechReq true "基本参数"
// @Success     200 {string} string "JSON 格式的存档数据"
// @Failure     403 {string} string "错误码"
// @Router      /square/publish [post]
func (h *Handler) SquarePublish(c *fiber.Ctx) error {
	resp := &PublishSpeechReq{}

	return c.Status(http.StatusOK).JSON(resp)
}

// 表态
