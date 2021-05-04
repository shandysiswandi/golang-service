package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shandysiswandi/echo-service/internal/domain"
)

func (h *Handler) FetchTodos(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.tdu.FetchTodos(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) GetTodoById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	data, err := h.tdu.GetTodoByID(ctx, id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.TodoCreatePayload{}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// call usecase or business logic
	payload.ID = h.generator.Nanoid()
	payload.Timestamp = h.clock.Now()
	if err := h.tdu.CreateTodo(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "OK")
}

func (h *Handler) UpdateTodoById(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.TodoUpdatePayload{ID: c.Param("id")}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// call usecase or business logic
	if err := h.tdu.UpdateTodoByID(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "OK")
}

func (h *Handler) ReplaceTodoById(c echo.Context) error {
	ctx := c.Request().Context()
	payload := domain.TodoReplacePayload{ID: c.Param("id")}

	// biding payload to struct
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// call usecase or business logic
	if err := h.tdu.ReplaceTodoByID(ctx, payload); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteTodoById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	// call usecase or business logic
	if err := h.tdu.DeleteTodoByID(ctx, id); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "OK")
}
