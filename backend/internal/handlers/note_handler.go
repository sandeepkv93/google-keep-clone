package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google-keep-clone/internal/services"
	"google-keep-clone/internal/validators"
	"strconv"
)

type NoteHandler struct {
	noteService *services.NoteService
}

func NewNoteHandler(noteService *services.NoteService) *NoteHandler {
	return &NoteHandler{noteService: noteService}
}

// @Summary Get all notes
// @Description Get all notes for the authenticated user
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param archived query bool false "Include archived notes"
// @Param deleted query bool false "Include deleted notes"
// @Success 200 {array} models.Note
// @Router /notes [get]
func (h *NoteHandler) GetNotes(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	includeArchived := c.QueryBool("archived", false)
	includeDeleted := c.QueryBool("deleted", false)

	notes, err := h.noteService.GetNotesByUserID(userID, includeArchived, includeDeleted)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}

// @Summary Create note
// @Description Create a new note
// @Tags notes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body validators.CreateNoteRequest true "Note data"
// @Success 201 {object} models.Note
// @Router /notes [post]
func (h *NoteHandler) CreateNote(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	var req validators.CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateCreateNoteRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	note, err := h.noteService.CreateNote(userID, &req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(note)
}

// @Summary Get note by ID
// @Description Get a specific note by ID
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note
// @Router /notes/{id} [get]
func (h *NoteHandler) GetNoteByID(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	note, err := h.noteService.GetNoteByID(noteID, userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}

// @Summary Update note
// @Description Update an existing note
// @Tags notes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Param request body validators.UpdateNoteRequest true "Note data"
// @Success 200 {object} models.Note
// @Router /notes/{id} [put]
func (h *NoteHandler) UpdateNote(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var req validators.UpdateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateUpdateNoteRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	note, err := h.noteService.UpdateNote(noteID, userID, &req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}

// @Summary Delete note
// @Description Delete a note (soft delete by default)
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Param permanent query bool false "Permanent delete"
// @Success 204
// @Router /notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	permanent := c.QueryBool("permanent", false)
	soft := !permanent

	if err := h.noteService.DeleteNote(noteID, userID, soft); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}

// @Summary Toggle pin
// @Description Toggle pin status of a note
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note
// @Router /notes/{id}/pin [patch]
func (h *NoteHandler) TogglePin(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	note, err := h.noteService.TogglePin(noteID, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}

// @Summary Toggle archive
// @Description Toggle archive status of a note
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note
// @Router /notes/{id}/archive [patch]
func (h *NoteHandler) ToggleArchive(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	note, err := h.noteService.ToggleArchive(noteID, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}

// @Summary Update note color
// @Description Update the color of a note
// @Tags notes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Note ID"
// @Param request body validators.ColorUpdateRequest true "Color data"
// @Success 200 {object} models.Note
// @Router /notes/{id}/color [patch]
func (h *NoteHandler) UpdateColor(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))
	noteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var req validators.ColorUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateColorUpdateRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	note, err := h.noteService.UpdateColor(noteID, userID, req.Color)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}

// @Summary Search notes
// @Description Search notes by title and content
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Param q query string true "Search query"
// @Param limit query int false "Limit results" default(20)
// @Param page query int false "Page number" default(0)
// @Success 200 {array} models.Note
// @Router /notes/search [get]
func (h *NoteHandler) SearchNotes(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	query := c.Query("q", "")
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	page, _ := strconv.Atoi(c.Query("page", "0"))

	req := validators.SearchRequest{
		Query: query,
		Limit: limit,
		Page:  page,
	}

	if err := validators.ValidateSearchRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	notes, err := h.noteService.SearchNotes(userID, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}

// @Summary Advanced search notes
// @Description Search notes with advanced filters including labels and color
// @Tags notes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body validators.AdvancedSearchRequest true "Advanced search parameters"
// @Success 200 {array} models.Note
// @Router /notes/search/advanced [post]
func (h *NoteHandler) SearchNotesAdvanced(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	var req validators.AdvancedSearchRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := validators.ValidateAdvancedSearchRequest(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Convert string label IDs to UUID slice
	var labelIDs []uuid.UUID
	for _, labelIDStr := range req.LabelIDs {
		if labelID, err := uuid.Parse(labelIDStr); err == nil {
			labelIDs = append(labelIDs, labelID)
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid label ID: " + labelIDStr})
		}
	}

	notes, err := h.noteService.SearchNotesAdvanced(userID, req.Query, labelIDs, req.Color, req.IncludeArchived)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}

// @Summary Get pinned notes
// @Description Get all pinned notes for the authenticated user
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Note
// @Router /notes/pinned [get]
func (h *NoteHandler) GetPinnedNotes(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	notes, err := h.noteService.GetPinnedNotes(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}

// @Summary Get archived notes
// @Description Get all archived notes for the authenticated user
// @Tags notes
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Note
// @Router /notes/archived [get]
func (h *NoteHandler) GetArchivedNotes(c *fiber.Ctx) error {
	userID, _ := uuid.Parse(c.Locals("userID").(string))

	notes, err := h.noteService.GetArchivedNotes(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}
