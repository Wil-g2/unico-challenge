package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wil-g2/unico-challenge/domain/fair"
	"github.com/wil-g2/unico-challenge/infra/log"
)

type FairHandler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindByName(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type fairHandler struct {
	service fair.Service
}

func NewUserHandler(service fair.Service) FairHandler {
	return &fairHandler{service}
}

// List godoc
// @Summary      List fairs
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Success      200  {array} fair.Fair
// @Router       /fairs [get]
func (h *fairHandler) List(ctx *fiber.Ctx) error {
	fairs, err := h.service.List()
	log.Info("list method run on handler file", nil)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fairs)
}

// Find godoc
// @Summary      Find fair
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Param        id		path	int		true	"ID"
// @Success      200  {object} fair.Fair
// @Router       /fairs/{id} [get]
func (h *fairHandler) Find(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	fair, err := h.service.Find(id)
	log.Info("find method run on handler file", nil)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fair)
}

// Create godoc
// @Summary      Create fairs
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Param		 raw	body	fair.CreateDTO		true	"body raw"
// @Success      200  {object} fair.Fair
// @Router       /fairs [post]
func (h *fairHandler) Create(ctx *fiber.Ctx) error {
	fairDto := new(fair.CreateDTO)
	if err := ctx.BodyParser(fairDto); err != nil {
		return err
	}

	err := h.service.Create(fairDto)
	log.Info("create method run on handler file", nil)
	if err != nil {
		log.Error("list method run on handler file", err, nil)
		return err
	}
	return nil
	// return ctx.Status(fiber.StatusCreated).JSON(user)
}

// Update godoc
// @Summary      Update fair
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Param        id		path	int		true	"ID"
// @Param		 raw	body	fair.UpdateDTO	true	"body raw"
// @Success      200  {object} fair.Fair
// @Router       /fairs/{id} [put]
func (h *fairHandler) Update(ctx *fiber.Ctx) error {
	fairDto := new(fair.UpdateDTO)
	if err := ctx.BodyParser(fairDto); err != nil {
		log.Error("error body parser", err, nil)
		return err
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Error("error trying to convert value ", err, nil)
		return err
	}
	result, err := h.service.Update(id, fairDto)
	log.Info("update method run on handler file", nil)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

// Delete godoc
// @Summary      Delete fair
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Param        id		path	int		true	"ID"
// @Success      200
// @Router       /fairs/{id} [delete]
func (h *fairHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	if err := h.service.Delete(id); err != nil {
		log.Error("error trying to delete fair", err, nil)
		return err
	}
	log.Info("delete method run on handler file", nil)
	return ctx.SendStatus(fiber.StatusNoContent)
}

// List godoc
// @Summary      List fairs by Name
// @Tags         fairs
// @Accept       json
// @Produce      json
// @Success      200  {array} fair.Fair
// @Router       /fairs/{name} [get]
// FindByName implements FairHandler
func (h *fairHandler) FindByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	fairs, err := h.service.FindByName(name)
	log.Info("list by name method run on handler file", nil)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fairs)
}
