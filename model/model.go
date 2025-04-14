package model

import (
	"time"
	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID
	Filename  string
	Type	  string 
	Size      int64
	CreatedAt time.Time
}