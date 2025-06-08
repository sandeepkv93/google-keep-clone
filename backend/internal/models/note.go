package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Note struct {
    ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
    Title       string    `json:"title"`
    Content     string    `json:"content" gorm:"type:text"`
    Color       string    `json:"color" gorm:"default:'#ffffff'"`
    IsPinned    bool      `json:"is_pinned" gorm:"default:false"`
    IsArchived  bool      `json:"is_archived" gorm:"default:false"`
    IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
    Position    int       `json:"position" gorm:"default:0"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

    User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Labels      []Label   `json:"labels,omitempty" gorm:"many2many:note_labels;"`
    Attachments []Attachment `json:"attachments,omitempty" gorm:"foreignKey:NoteID"`
}

type Label struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
    Name      string    `json:"name" gorm:"not null"`
    Color     string    `json:"color" gorm:"default:'#ffffff'"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    Notes     []Note    `json:"notes,omitempty" gorm:"many2many:note_labels;"`
}

type Attachment struct {
    ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    NoteID   uuid.UUID `json:"note_id" gorm:"type:uuid;not null;index"`
    Filename string    `json:"filename" gorm:"not null"`
    URL      string    `json:"url" gorm:"not null"`
    Size     int64     `json:"size"`
    MimeType string    `json:"mime_type"`
    CreatedAt time.Time `json:"created_at"`

    Note     Note      `json:"note,omitempty" gorm:"foreignKey:NoteID"`
}