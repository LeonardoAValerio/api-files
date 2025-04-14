package Model

type File struct {
	ID        uint
	Filename  string
	MimeType  string 
	Size      int64
	CreatedAt time.Time
}