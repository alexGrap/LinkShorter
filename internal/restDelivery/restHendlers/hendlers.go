package restHendlers

import (
	"github.com/gofiber/fiber/v2"
	"ozon/internal/models"
	"ozon/internal/usecase"
)

func Post(ctx *fiber.Ctx) error {
	var (
		body models.RestResult
		err  models.OwnError
	)
	key := ctx.Query("startLink")
	if models.DB == "redis" {
		body.ResultLink, err = usecase.CreationRedis(key)
	} else {
		body.ResultLink, err = usecase.CreateShortLink(key)
	}
	if err.Err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(body)
}

func Get(ctx *fiber.Ctx) error {
	var (
		body models.RestResult
		err  models.OwnError
	)
	key := ctx.Query("startLink")
	if models.DB == "redis" {
		body.ResultLink, err = usecase.GetterRedis(key)
	} else {
		body.ResultLink, err = usecase.GetFullLink(key)
	}
	if err.Err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(body)
}
