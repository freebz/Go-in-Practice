// Listing 9.4  Errors to go with cloud functionality interface

var (
	ErrFileNotFound   = errors.New("File not found")
	ErrCannotLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)
