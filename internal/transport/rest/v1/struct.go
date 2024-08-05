package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lapkomo2018/goServiceExample/internal/core"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) initStructApi(api *echo.Group) {
	api.POST("/add", h.structAdd)
	api.GET("/info/:id", h.structInfo)
}

// @Summary      Add Struct
// @Description  add struct with msg
// @Tags         Struct
// @Accept       json
// @Produce      json
// @Param input body structAddReq true "struct obj"
// @Success      200  {object}  structAddResp
// @Failure      default  {object}  echo.HTTPError
// @Router       /add [post]
func (h *Handler) structAdd(c echo.Context) error {
	var body structAddReq
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Message(body.Message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	str := &core.Struct{
		Message: body.Message,
	}

	if err := h.structService.Add(str); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, structAddResp{
		ID: str.ID,
	})
}

type (
	structAddReq struct {
		Message string `json:"message"`
	}
	structAddResp struct {
		ID uint `json:"id"`
	}
)

// @Summary      Info
// @Description  Get struct info
// @Tags         Struct
// @Produce      json
// @Param        id   path     int  true  "ID of the struct"
// @Success      200  {object} structInfoResp
// @Failure      default  {object}  echo.HTTPError
// @Router       /info/{id} [get]
func (h *Handler) structInfo(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("bad param id"))
	}

	str, err := h.structService.FindByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, structInfoResp{
		ID:        str.ID,
		Message:   str.Message,
		CreatedAt: str.CreatedAt,
		UpdatedAt: str.UpdatedAt,
	})
}

type (
	structInfoResp struct {
		ID        uint      `json:"id"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
