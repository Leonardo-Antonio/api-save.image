package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Leonardo-Antonio/api-save.image/src/helper"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/gofiber/fiber/v2"
)

type ImageForm struct{}

func (i *ImageForm) SaveImage(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(helper.Response{
				MessageType: helper.ERROR,
				Message:     err.Error(),
				Error:       true,
				Data:        nil,
			})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(helper.Response{
				MessageType: helper.ERROR,
				Message:     err.Error(),
				Error:       true,
				Data:        nil,
			})
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(helper.Response{
				MessageType: helper.ERROR,
				Message:     err.Error(),
				Error:       true,
				Data:        nil,
			})
	}

	uuid, err := nanoid.New()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(helper.Response{
				MessageType: helper.ERROR,
				Message:     err.Error(),
				Error:       true,
				Data:        nil,
			})
	}

	fileImage := strings.Split(fileHeader.Filename, " ")
	uri := uuid + "-" + strings.Join(fileImage, "-")
	if err := ioutil.WriteFile(
		"./public/images/"+uri,
		fileBytes,
		os.ModePerm,
	); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(helper.Response{
				MessageType: helper.ERROR,
				Message:     err.Error(),
				Error:       true,
				Data:        nil,
			})
	}

	return ctx.Status(http.StatusCreated).
		JSON(helper.Response{
			MessageType: helper.MESSAGE,
			Message:     "the image was saved successfully",
			Error:       false,
			Data: map[string]string{
				"url": ctx.BaseURL() + "/api/v1/images/" + uri,
			},
		})
}
