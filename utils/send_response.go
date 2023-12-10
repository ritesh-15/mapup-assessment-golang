package utils

import "github.com/gofiber/fiber/v2"

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func SendResponse(c *fiber.Ctx, arrays [][]int, timeTaken int64) error {
	response := ResponsePayload{
		TimeNs:       timeTaken,
		SortedArrays: arrays,
	}

	return c.JSON(response)
}
