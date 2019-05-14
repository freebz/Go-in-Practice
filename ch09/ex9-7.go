// Listing 9.7  Function to check whether the application is available

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		es := "Could not find '%s' in PATH: %s"
		return fmt.Errorf(es, name, err)
	}

	return nil
}
