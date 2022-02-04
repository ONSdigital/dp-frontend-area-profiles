package utils

// GetFirstError returns the first error or nil if there are no errors
func GetFirstError(errors ...error) error {
	for i := range errors {
		if errors[i] != nil {
			return errors[i]
		}
	}
	return nil
}
