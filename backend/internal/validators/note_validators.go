package validators

import (
    "errors"
    "strings"
    "regexp"
)

type CreateNoteRequest struct {
    Title    string `json:"title"`
    Content  string `json:"content"`
    Color    string `json:"color"`
    IsPinned *bool  `json:"is_pinned"`
}

type UpdateNoteRequest struct {
    Title      *string `json:"title"`
    Content    *string `json:"content"`
    Color      *string `json:"color"`
    IsPinned   *bool   `json:"is_pinned"`
    IsArchived *bool   `json:"is_archived"`
    Position   *int    `json:"position"`
}

type ColorUpdateRequest struct {
    Color string `json:"color" validate:"required"`
}

type SearchRequest struct {
    Query string `json:"query"`
    Limit int    `json:"limit"`
    Page  int    `json:"page"`
}

type AdvancedSearchRequest struct {
    Query           string   `json:"query,omitempty"`
    LabelIDs        []string `json:"label_ids,omitempty"`
    Color           string   `json:"color,omitempty"`
    IncludeArchived bool     `json:"include_archived,omitempty"`
}

func ValidateCreateNoteRequest(req *CreateNoteRequest) error {
    // At least title or content must be provided
    if strings.TrimSpace(req.Title) == "" && strings.TrimSpace(req.Content) == "" {
        return errors.New("either title or content must be provided")
    }

    // Validate title length
    if len(req.Title) > 255 {
        return errors.New("title must be less than 255 characters")
    }

    // Validate content length (reasonable limit for notes)
    if len(req.Content) > 10000 {
        return errors.New("content must be less than 10,000 characters")
    }

    // Validate color format if provided
    if req.Color != "" {
        if err := validateColor(req.Color); err != nil {
            return err
        }
    }

    return nil
}

func ValidateUpdateNoteRequest(req *UpdateNoteRequest) error {
    // Validate title length if provided
    if req.Title != nil && len(*req.Title) > 255 {
        return errors.New("title must be less than 255 characters")
    }

    // Validate content length if provided
    if req.Content != nil && len(*req.Content) > 10000 {
        return errors.New("content must be less than 10,000 characters")
    }

    // Validate color format if provided
    if req.Color != nil && *req.Color != "" {
        if err := validateColor(*req.Color); err != nil {
            return err
        }
    }

    // Validate position if provided
    if req.Position != nil && *req.Position < 0 {
        return errors.New("position must be non-negative")
    }

    return nil
}

func ValidateColorUpdateRequest(req *ColorUpdateRequest) error {
    return validateColor(req.Color)
}

func ValidateSearchRequest(req *SearchRequest) error {
    // Validate query length
    if len(req.Query) > 100 {
        return errors.New("search query must be less than 100 characters")
    }

    // Validate limit
    if req.Limit < 0 || req.Limit > 100 {
        return errors.New("limit must be between 0 and 100")
    }

    // Validate page
    if req.Page < 0 {
        return errors.New("page must be non-negative")
    }

    return nil
}

func ValidateAdvancedSearchRequest(req *AdvancedSearchRequest) error {
    // Validate query length
    if len(req.Query) > 100 {
        return errors.New("search query must be less than 100 characters")
    }

    // Validate color format if provided
    if req.Color != "" {
        if err := validateColor(req.Color); err != nil {
            return err
        }
    }

    // Validate label IDs if provided
    if len(req.LabelIDs) > 10 {
        return errors.New("cannot filter by more than 10 labels at once")
    }

    return nil
}

func validateColor(color string) error {
    if color == "" {
        return nil // Allow empty color to reset to default
    }

    // Check if it's a valid hex color
    hexColorRegex := regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
    if hexColorRegex.MatchString(color) {
        return nil
    }

    // Check if it's a predefined color name
    predefinedColors := map[string]bool{
        "white":   true,
        "red":     true,
        "orange":  true,
        "yellow":  true,
        "green":   true,
        "teal":    true,
        "blue":    true,
        "purple":  true,
        "pink":    true,
        "brown":   true,
        "gray":    true,
        "grey":    true,
    }

    if predefinedColors[strings.ToLower(color)] {
        return nil
    }

    return errors.New("invalid color format. Use hex color (#rrggbb) or predefined color name")
}