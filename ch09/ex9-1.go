// Listing 9.1  Interface for cloud functionality

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}
