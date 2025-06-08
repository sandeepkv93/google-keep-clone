package repositories

import (
	"github.com/google/uuid"
	"google-keep-clone/internal/models"
	"gorm.io/gorm"
)

type LabelRepository struct {
	db *gorm.DB
}

func NewLabelRepository(db *gorm.DB) *LabelRepository {
	return &LabelRepository{db: db}
}

func (r *LabelRepository) Create(label *models.Label) error {
	return r.db.Create(label).Error
}

func (r *LabelRepository) GetByUserID(userID uuid.UUID) ([]models.Label, error) {
	var labels []models.Label
	err := r.db.Where("user_id = ?", userID).Order("name ASC").Find(&labels).Error
	return labels, err
}

func (r *LabelRepository) GetByID(id, userID uuid.UUID) (*models.Label, error) {
	var label models.Label
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&label).Error
	return &label, err
}

func (r *LabelRepository) Update(label *models.Label) error {
	return r.db.Save(label).Error
}

func (r *LabelRepository) Delete(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Label{}).Error
}

func (r *LabelRepository) GetByName(userID uuid.UUID, name string) (*models.Label, error) {
	var label models.Label
	err := r.db.Where("user_id = ? AND name = ?", userID, name).First(&label).Error
	return &label, err
}

func (r *LabelRepository) AttachToNote(noteID, labelID uuid.UUID) error {
	return r.db.Exec("INSERT INTO note_labels (note_id, label_id) VALUES (?, ?) ON CONFLICT DO NOTHING", noteID, labelID).Error
}

func (r *LabelRepository) DetachFromNote(noteID, labelID uuid.UUID) error {
	return r.db.Exec("DELETE FROM note_labels WHERE note_id = ? AND label_id = ?", noteID, labelID).Error
}

func (r *LabelRepository) GetNotesByLabel(labelID, userID uuid.UUID) ([]models.Note, error) {
	var notes []models.Note
	err := r.db.Table("notes").
		Joins("INNER JOIN note_labels ON notes.id = note_labels.note_id").
		Where("note_labels.label_id = ? AND notes.user_id = ? AND notes.is_deleted = ?", labelID, userID, false).
		Preload("Labels").
		Order("notes.updated_at DESC").
		Find(&notes).Error
	return notes, err
}
