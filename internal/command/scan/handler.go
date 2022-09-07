package scan

import (
	"github.com/go-git/go-git/v5"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args *Args
}

// New instantiates new scan command handler.
func New(args *Args) *Handler {
	return &Handler{Args: args}
}

// Execute runs scan command.
func (h *Handler) Execute() error {
	if err := h.Args.Validate(); err != nil {
		return err
	}

	return nil
}
