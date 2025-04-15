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

type Chunk struct {
	ID        		uuid.UUID
	Current_size 	int64
	Max_size      	int64
	CreatedAt 		time.Time
}