package domain_media_handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	dto "github.com/felixa1996/go_next_be/app/domain/media/dto"
	error_wrapper "github.com/felixa1996/go_next_be/app/infra/error"
	"github.com/felixa1996/go_next_be/app/infra/response"
	"github.com/felixa1996/go_next_be/app/infra/validator"
)

// MediaCreate godoc
// @Summary      Create media
// @Description  Create media
// @Accept       multipart/form-data
// @Security	  JWT
// @Tags         Media
// @Produce      json
// @Param        files  formData file true "Support multiple files"
// @Param        uri  formData string true "Uri"
// @Success      200  {object}  response.JSONSuccessResult{data=domain_media.Media,code=int,message=string}
// @Failure      400  {object}  response.JSONBadRequest{code=int,message=string}
// @Failure      422  {object}  response.JSONUnprocessableEntity{code=int,message=string}
// @Failure      500  {object}  response.JSONInternalServerError{code=int,message=string}
// @Router       /v1/media [post]
func (h *MediaHandler) Create(c echo.Context) error {
	var ew error_wrapper.ErrorWrapper

	ctx := c.Request().Context()

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// manual bind required because multipart/form-data
	var dto dto.MediaDtoCreateInput
	dto.Uri = c.FormValue("uri")
	dto.Files = form.File["files"]

	err = h.validate.Struct(dto)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate create media", zap.Error(err))
		return response.FailResponse(c, http.StatusUnprocessableEntity, "Failed to process entity", localizedErr)
	}

	res, err := h.usecase.Create(ctx, dto)
	if err != nil && errors.As(err, &ew) {
		h.logger.Error("Failed to create media", zap.Error(err))
		return response.FailResponse(c, ew.Code, ew.Message, ew.Err.Error())
	}

	return response.SuccessReponse(c, res)
}
