package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/plab0n/koinoni_backend/internal/storage"
	"github.com/plab0n/koinoni_backend/pkg/httputils"
)

// Handlers implements all the handler functions and has the dependencies that they use (Sender, Storage).
type Handlers struct {
	Sender  *httputils.Sender
	Storage storage.StorageInterface
}

// Validate is a singleton that provides validation services for in handlers.
var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
