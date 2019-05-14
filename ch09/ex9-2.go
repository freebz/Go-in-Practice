// Listing 9.2  Simple implementation of cloud file storage

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {

	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}
