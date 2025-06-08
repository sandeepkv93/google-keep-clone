package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google-keep-clone/internal/services"
	"google-keep-clone/internal/validators"
)

type LabelHandler struct {
	labelService *services.LabelService
}

func NewLabelHandler(labelService *services.LabelService) *LabelHandler {
	return &LabelHandler{labelService: labelService}
}

// @Summary Get all labels
// @Description Get all labels for the authenticated user
// @Tags labels
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Label
// @Router /labels [get]
func (h *LabelHandler) GetLabels(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	labels, err := h.labelService.GetLabelsByUserID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(labels)
}

// @Summary Create label
// @Description Create a new label
// @Tags labels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body validators.CreateLabelRequest true "Label data"
// @Success 201 {object} models.Label
// @Router /labels [post]
func (h *LabelHandler) CreateLabel(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	var req validators.CreateLabelRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateCreateLabelRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	label, err := h.labelService.CreateLabel(userID, &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(label)
}

// @Summary Get label by ID
// @Description Get a specific label by ID
// @Tags labels
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Label ID"
// @Success 200 {object} models.Label
// @Router /labels/{id} [get]
func (h *LabelHandler) GetLabelByID(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	labelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	label, err := h.labelService.GetLabelByID(labelID, userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(label)
}

// @Summary Update label
// @Description Update an existing label
// @Tags labels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Label ID"
// @Param request body validators.UpdateLabelRequest true "Label data"
// @Success 200 {object} models.Label
// @Router /labels/{id} [put]
func (h *LabelHandler) UpdateLabel(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	labelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	var req validators.UpdateLabelRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateUpdateLabelRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	label, err := h.labelService.UpdateLabel(labelID, userID, &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(label)
}

// @Summary Delete label
// @Description Delete a label
// @Tags labels
// @Security ApiKeyAuth
// @Param id path string true "Label ID"
// @Success 204
// @Router /labels/{id} [delete]
func (h *LabelHandler) DeleteLabel(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	labelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	if err := h.labelService.DeleteLabel(labelID, userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}

// @Summary Attach label to note
// @Description Attach a label to a note
// @Tags labels
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param note_id path string true "Note ID"
// @Param request body validators.AttachLabelRequest true "Label attachment data"
// @Success 200 {object} map[string]string
// @Router /notes/{note_id}/labels [post]
func (h *LabelHandler) AttachLabelToNote(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("note_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var req validators.AttachLabelRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	labelID, err := uuid.Parse(req.LabelID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	if err := h.labelService.AttachLabelToNote(noteID, labelID, userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Label attached to note successfully"})
}

// @Summary Detach label from note
// @Description Detach a label from a note
// @Tags labels
// @Security ApiKeyAuth
// @Param note_id path string true "Note ID"
// @Param label_id path string true "Label ID"
// @Success 200 {object} map[string]string
// @Router /notes/{note_id}/labels/{label_id} [delete]
func (h *LabelHandler) DetachLabelFromNote(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("note_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	labelID, err := uuid.Parse(c.Params("label_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	if err := h.labelService.DetachLabelFromNote(noteID, labelID, userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Label detached from note successfully"})
}

// @Summary Get notes by label
// @Description Get all notes with a specific label
// @Tags labels
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Label ID"
// @Success 200 {array} models.Note
// @Router /labels/{id}/notes [get]
func (h *LabelHandler) GetNotesByLabel(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	labelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID"})
	}

	notes, err := h.labelService.GetNotesByLabel(labelID, userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}