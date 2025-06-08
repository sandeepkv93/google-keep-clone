package services

import (
    "errors"
    "time"
    "github.com/google/uuid"
    "google-keep-clone/backend/internal/models"
    "google-keep-clone/backend/internal/repositories"
    "google-keep-clone/backend/internal/validators"
)

type NoteService struct {
    noteRepo *repositories.NoteRepository
    userRepo *repositories.UserRepository
}

func NewNoteService(noteRepo *repositories.NoteRepository, userRepo *repositories.UserRepository) *NoteService {
    return &NoteService{
        noteRepo: noteRepo,
        userRepo: userRepo,
    }
}

func (s *NoteService) CreateNote(userID uuid.UUID, req *validators.CreateNoteRequest) (*models.Note, error) {
    // Verify user exists
    _, err := s.userRepo.GetByID(userID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    note := &models.Note{
        UserID:   userID,
        Title:    req.Title,
        Content:  req.Content,
        Color:    "#ffffff", // default color
        IsPinned: false,
        IsArchived: false,
        IsDeleted: false,
        Position: 0,
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

    return s.noteRepo.Delete(id, userID)
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

func (s *NoteService) GetPinnedNotes(userID uuid.UUID) ([]models.Note, error) {
    return s.noteRepo.GetPinnedNotes(userID)
}

func (s *NoteService) GetArchivedNotes(userID uuid.UUID) ([]models.Note, error) {
    return s.noteRepo.GetArchivedNotes(userID)
}