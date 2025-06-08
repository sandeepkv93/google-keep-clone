package validators

import (
	"errors"
	"strings"
)

type CreateLabelRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=50"`
	Color string `json:"color,omitempty" validate:"omitempty,hexcolor"`
}

type UpdateLabelRequest struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,min=1,max=50"`
	Color *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
}

type AttachLabelRequest struct {
	LabelID string `json:"label_id" validate:"required,uuid"`
}

func ValidateCreateLabelRequest(req *CreateLabelRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}

	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) == 0 {
		return errors.New("name cannot be empty")
	}

	if len(req.Name) > 50 {
		return errors.New("name cannot exceed 50 characters")
	}

	if req.Color != "" && !isValidHexColor(req.Color) {
		return errors.New("color must be a valid hex color")
	}

	return nil
}

func ValidateUpdateLabelRequest(req *UpdateLabelRequest) error {
	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if len(name) == 0 {
			return errors.New("name cannot be empty")
		}
		if len(name) > 50 {
			return errors.New("name cannot exceed 50 characters")
		}
		*req.Name = name
	}

	if req.Color != nil && !isValidHexColor(*req.Color) {
		return errors.New("color must be a valid hex color")
	}

	return nil
}

func isValidHexColor(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}
	for i := 1; i < len(color); i++ {
		c := color[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}