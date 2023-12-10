package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ritesh-15/mapup-assessment/utils"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

func ProcessSingle(c *fiber.Ctx) error {
	var payload RequestPayload
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "input is not in specified format",
		})
	}

	startTime := time.Now()

	sortedArrays := utils.SequentialSort(payload.ToSort)

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime).Nanoseconds()

	return utils.SendResponse(c, sortedArrays, timeTaken)
}

func ProcessConcurrent(c *fiber.Ctx) error {
	var payload RequestPayload
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "input is not in specified format",
		})
	}

	startTime := time.Now()

	sortedArrays := utils.ConcurrentSort(payload.ToSort)

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime).Nanoseconds()

	return utils.SendResponse(c, sortedArrays, timeTaken)
}
