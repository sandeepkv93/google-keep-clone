package services

import (
	"errors"
	"github.com/google/uuid"
	"google-keep-clone/internal/models"
	"google-keep-clone/internal/repositories"
	"google-keep-clone/internal/validators"
	"google-keep-clone/internal/websocket"
	"time"
)

type NoteService struct {
	noteRepo *repositories.NoteRepository
	userRepo *repositories.UserRepository
	hub      *websocket.Hub
}

func NewNoteService(noteRepo *repositories.NoteRepository, userRepo *repositories.UserRepository, hub *websocket.Hub) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
		userRepo: userRepo,
		hub:      hub,
	}
}

func (s *NoteService) CreateNote(userID uuid.UUID, req *validators.CreateNoteRequest) (*models.Note, error) {
	// Verify user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	note := &models.Note{
		UserID:     userID,
		Title:      req.Title,
		Content:    req.Content,
		Color:      "#ffffff", // default color
		IsPinned:   false,
		IsArchived: false,
		IsDeleted:  false,
		Position:   0,
	}

	if req.Color != "" {
		note.Color = req.Color
	}

	if req.IsPinned != nil {
		note.IsPinned = *req.IsPinned
	}

	if err := s.noteRepo.Create(note); err != nil {
		return nil, errors.New("failed to create note")
	}

	// Broadcast note creation to WebSocket clients
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "note_created", note)
	}

	return note, nil
}

func (s *NoteService) GetNotesByUserID(userID uuid.UUID, includeArchived, includeDeleted bool) ([]models.Note, error) {
	return s.noteRepo.GetByUserID(userID, includeArchived, includeDeleted)
}

func (s *NoteService) GetNoteByID(id, userID uuid.UUID) (*models.Note, error) {
	note, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("note not found")
	}
	return note, nil
}

func (s *NoteService) UpdateNote(id, userID uuid.UUID, req *validators.UpdateNoteRequest) (*models.Note, error) {
	// Get existing note
	note, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("note not found")
	}

	// Update fields if provided
	if req.Title != nil {
		note.Title = *req.Title
	}

	if req.Content != nil {
		note.Content = *req.Content
	}

	if req.Color != nil {
		note.Color = *req.Color
	}

	if req.IsPinned != nil {
		note.IsPinned = *req.IsPinned
	}

	if req.IsArchived != nil {
		note.IsArchived = *req.IsArchived
	}

	if req.Position != nil {
		note.Position = *req.Position
	}

	note.UpdatedAt = time.Now()

	if err := s.noteRepo.Update(note); err != nil {
		return nil, errors.New("failed to update note")
	}

	// Broadcast note update to WebSocket clients
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "note_updated", note)
	}

	return note, nil
}

func (s *NoteService) DeleteNote(id, userID uuid.UUID, soft bool) error {
	// Check if note exists and belongs to user
	_, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return errors.New("note not found")
	}

	if soft {
		return s.noteRepo.SoftDelete(id, userID)
	}

	err = s.noteRepo.Delete(id, userID)

	// Broadcast note deletion to WebSocket clients
	if err == nil && s.hub != nil {
		s.hub.BroadcastToUser(userID, "note_deleted", map[string]string{"id": id.String()})
	}

	return err
}

func (s *NoteService) TogglePin(id, userID uuid.UUID) (*models.Note, error) {
	// Check if note exists and belongs to user
	_, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("note not found")
	}

	if err := s.noteRepo.TogglePin(id, userID); err != nil {
		return nil, errors.New("failed to toggle pin")
	}

	// Return updated note
	return s.noteRepo.GetByID(id, userID)
}

func (s *NoteService) ToggleArchive(id, userID uuid.UUID) (*models.Note, error) {
	// Check if note exists and belongs to user
	_, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("note not found")
	}

	if err := s.noteRepo.ToggleArchive(id, userID); err != nil {
		return nil, errors.New("failed to toggle archive")
	}

	// Return updated note
	return s.noteRepo.GetByID(id, userID)
}

func (s *NoteService) UpdateColor(id, userID uuid.UUID, color string) (*models.Note, error) {
	// Check if note exists and belongs to user
	_, err := s.noteRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("note not found")
	}

	if err := s.noteRepo.UpdateColor(id, userID, color); err != nil {
		return nil, errors.New("failed to update color")
	}

	// Return updated note
	return s.noteRepo.GetByID(id, userID)
}

func (s *NoteService) SearchNotes(userID uuid.UUID, query string) ([]models.Note, error) {
	if query == "" {
		return s.noteRepo.GetByUserID(userID, false, false)
	}

	return s.noteRepo.Search(userID, query)
}

func (s *NoteService) SearchNotesAdvanced(userID uuid.UUID, query string, labelIDs []uuid.UUID, color string, includeArchived bool) ([]models.Note, error) {
	// If searching by color specifically
	if color != "" && query == "" && len(labelIDs) == 0 {
		return s.noteRepo.SearchByColor(userID, color, includeArchived)
	}

	// If using advanced search with labels or other filters
	if query != "" || len(labelIDs) > 0 {
		return s.noteRepo.SearchWithLabels(userID, query, labelIDs, includeArchived)
	}

	// Default to getting all notes
	return s.noteRepo.GetByUserID(userID, includeArchived, false)
}

func (s *NoteService) GetPinnedNotes(userID uuid.UUID) ([]models.Note, error) {
	return s.noteRepo.GetPinnedNotes(userID)
}

func (s *NoteService) GetArchivedNotes(userID uuid.UUID) ([]models.Note, error) {
	return s.noteRepo.GetArchivedNotes(userID)
}
