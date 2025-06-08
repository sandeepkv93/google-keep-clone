package services

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"google-keep-clone/internal/models"
	"google-keep-clone/internal/repositories"
	"google-keep-clone/internal/validators"
	"google-keep-clone/internal/websocket"
)

type LabelService struct {
	labelRepo *repositories.LabelRepository
	noteRepo  *repositories.NoteRepository
	userRepo  *repositories.UserRepository
	hub       *websocket.Hub
}

func NewLabelService(labelRepo *repositories.LabelRepository, noteRepo *repositories.NoteRepository, userRepo *repositories.UserRepository, hub *websocket.Hub) *LabelService {
	return &LabelService{
		labelRepo: labelRepo,
		noteRepo:  noteRepo,
		userRepo:  userRepo,
		hub:       hub,
	}
}

func (s *LabelService) CreateLabel(userID uuid.UUID, req *validators.CreateLabelRequest) (*models.Label, error) {
	// Verify user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Check if label with same name already exists
	if _, err := s.labelRepo.GetByName(userID, req.Name); err == nil {
		return nil, errors.New("label with this name already exists")
	}

	label := &models.Label{
		UserID: userID,
		Name:   strings.TrimSpace(req.Name),
		Color:  "#ffffff", // default color
	}

	if req.Color != "" {
		label.Color = req.Color
	}

	if err := s.labelRepo.Create(label); err != nil {
		return nil, errors.New("failed to create label")
	}

	// Broadcast label creation to WebSocket clients
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "label_created", label)
	}

	return label, nil
}

func (s *LabelService) GetLabelsByUserID(userID uuid.UUID) ([]models.Label, error) {
	return s.labelRepo.GetByUserID(userID)
}

func (s *LabelService) GetLabelByID(id, userID uuid.UUID) (*models.Label, error) {
	label, err := s.labelRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("label not found")
	}
	return label, nil
}

func (s *LabelService) UpdateLabel(id, userID uuid.UUID, req *validators.UpdateLabelRequest) (*models.Label, error) {
	// Get existing label
	label, err := s.labelRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("label not found")
	}

	// Update fields if provided
	if req.Name != nil {
		newName := strings.TrimSpace(*req.Name)
		if newName != label.Name {
			// Check if another label with this name exists
			if existingLabel, err := s.labelRepo.GetByName(userID, newName); err == nil && existingLabel.ID != id {
				return nil, errors.New("label with this name already exists")
			}
			label.Name = newName
		}
	}

	if req.Color != nil {
		label.Color = *req.Color
	}

	if err := s.labelRepo.Update(label); err != nil {
		return nil, errors.New("failed to update label")
	}

	// Broadcast label update to WebSocket clients
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "label_updated", label)
	}

	return label, nil
}

func (s *LabelService) DeleteLabel(id, userID uuid.UUID) error {
	// Check if label exists and belongs to user
	_, err := s.labelRepo.GetByID(id, userID)
	if err != nil {
		return errors.New("label not found")
	}

	if err := s.labelRepo.Delete(id, userID); err != nil {
		return errors.New("failed to delete label")
	}

	// Broadcast label deletion to WebSocket clients
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "label_deleted", map[string]string{"id": id.String()})
	}

	return nil
}

func (s *LabelService) AttachLabelToNote(noteID, labelID, userID uuid.UUID) error {
	// Verify note exists and belongs to user
	_, err := s.noteRepo.GetByID(noteID, userID)
	if err != nil {
		return errors.New("note not found")
	}

	// Verify label exists and belongs to user
	_, err = s.labelRepo.GetByID(labelID, userID)
	if err != nil {
		return errors.New("label not found")
	}

	if err := s.labelRepo.AttachToNote(noteID, labelID); err != nil {
		return errors.New("failed to attach label to note")
	}

	// Get updated note for broadcasting
	note, _ := s.noteRepo.GetByID(noteID, userID)
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "note_updated", note)
	}

	return nil
}

func (s *LabelService) DetachLabelFromNote(noteID, labelID, userID uuid.UUID) error {
	// Verify note exists and belongs to user
	_, err := s.noteRepo.GetByID(noteID, userID)
	if err != nil {
		return errors.New("note not found")
	}

	// Verify label exists and belongs to user
	_, err = s.labelRepo.GetByID(labelID, userID)
	if err != nil {
		return errors.New("label not found")
	}

	if err := s.labelRepo.DetachFromNote(noteID, labelID); err != nil {
		return errors.New("failed to detach label from note")
	}

	// Get updated note for broadcasting
	note, _ := s.noteRepo.GetByID(noteID, userID)
	if s.hub != nil {
		s.hub.BroadcastToUser(userID, "note_updated", note)
	}

	return nil
}

func (s *LabelService) GetNotesByLabel(labelID, userID uuid.UUID) ([]models.Note, error) {
	// Verify label exists and belongs to user
	_, err := s.labelRepo.GetByID(labelID, userID)
	if err != nil {
		return nil, errors.New("label not found")
	}

	return s.labelRepo.GetNotesByLabel(labelID, userID)
}
