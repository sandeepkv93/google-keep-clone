package repositories

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "google-keep-clone/internal/models"
)

type NoteRepository struct {
    db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
    return &NoteRepository{db: db}
}

func (r *NoteRepository) Create(note *models.Note) error {
    return r.db.Create(note).Error
}

func (r *NoteRepository) GetByUserID(userID uuid.UUID, includeArchived, includeDeleted bool) ([]models.Note, error) {
    var notes []models.Note
    query := r.db.Where("user_id = ?", userID)

    if !includeArchived {
        query = query.Where("is_archived = ?", false)
    }
    if !includeDeleted {
        query = query.Where("is_deleted = ?", false)
    }

    err := query.Preload("Labels").Order("is_pinned DESC, position ASC, updated_at DESC").Find(&notes).Error
    return notes, err
}

func (r *NoteRepository) GetByID(id, userID uuid.UUID) (*models.Note, error) {
    var note models.Note
    err := r.db.Where("id = ? AND user_id = ?", id, userID).Preload("Labels").Preload("Attachments").First(&note).Error
    return &note, err
}

func (r *NoteRepository) Update(note *models.Note) error {
    return r.db.Save(note).Error
}

func (r *NoteRepository) Delete(id, userID uuid.UUID) error {
    return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{}).Error
}

func (r *NoteRepository) SoftDelete(id, userID uuid.UUID) error {
    return r.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", id, userID).Update("is_deleted", true).Error
}

func (r *NoteRepository) TogglePin(id, userID uuid.UUID) error {
    return r.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", id, userID).Update("is_pinned", gorm.Expr("NOT is_pinned")).Error
}

func (r *NoteRepository) ToggleArchive(id, userID uuid.UUID) error {
    return r.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", id, userID).Update("is_archived", gorm.Expr("NOT is_archived")).Error
}

func (r *NoteRepository) UpdateColor(id, userID uuid.UUID, color string) error {
    return r.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", id, userID).Update("color", color).Error
}

func (r *NoteRepository) UpdatePosition(id, userID uuid.UUID, position int) error {
    return r.db.Model(&models.Note{}).Where("id = ? AND user_id = ?", id, userID).Update("position", position).Error
}

func (r *NoteRepository) Search(userID uuid.UUID, query string) ([]models.Note, error) {
    var notes []models.Note
    err := r.db.Where("user_id = ? AND is_deleted = ? AND (title ILIKE ? OR content ILIKE ?)",
        userID, false, "%"+query+"%", "%"+query+"%").
        Preload("Labels").
        Order("updated_at DESC").
        Find(&notes).Error
    return notes, err
}

func (r *NoteRepository) SearchWithLabels(userID uuid.UUID, query string, labelIDs []uuid.UUID, includeArchived bool) ([]models.Note, error) {
    var notes []models.Note
    
    baseQuery := r.db.Where("user_id = ? AND is_deleted = ?", userID, false)
    
    if !includeArchived {
        baseQuery = baseQuery.Where("is_archived = ?", false)
    }
    
    // Add text search if query is provided
    if query != "" {
        baseQuery = baseQuery.Where("(title ILIKE ? OR content ILIKE ?)", "%"+query+"%", "%"+query+"%")
    }
    
    // Add label filtering if label IDs are provided
    if len(labelIDs) > 0 {
        baseQuery = baseQuery.Joins("INNER JOIN note_labels ON notes.id = note_labels.note_id").
            Where("note_labels.label_id IN ?", labelIDs).
            Group("notes.id")
    }
    
    err := baseQuery.Preload("Labels").
        Order("is_pinned DESC, updated_at DESC").
        Find(&notes).Error
    
    return notes, err
}

func (r *NoteRepository) SearchByColor(userID uuid.UUID, color string, includeArchived bool) ([]models.Note, error) {
    var notes []models.Note
    
    query := r.db.Where("user_id = ? AND is_deleted = ? AND color = ?", userID, false, color)
    
    if !includeArchived {
        query = query.Where("is_archived = ?", false)
    }
    
    err := query.Preload("Labels").
        Order("is_pinned DESC, updated_at DESC").
        Find(&notes).Error
    
    return notes, err
}

func (r *NoteRepository) GetPinnedNotes(userID uuid.UUID) ([]models.Note, error) {
    var notes []models.Note
    err := r.db.Where("user_id = ? AND is_pinned = ? AND is_deleted = ? AND is_archived = ?",
        userID, true, false, false).
        Preload("Labels").
        Order("position ASC, updated_at DESC").
        Find(&notes).Error
    return notes, err
}

func (r *NoteRepository) GetArchivedNotes(userID uuid.UUID) ([]models.Note, error) {
    var notes []models.Note
    err := r.db.Where("user_id = ? AND is_archived = ? AND is_deleted = ?",
        userID, true, false).
        Preload("Labels").
        Order("updated_at DESC").
        Find(&notes).Error
    return notes, err
}