package util

func ErrToString(errs []error) (stringsErrs []string) {
	for _, err := range errs {
		stringsErrs = append(stringsErrs, err.Error())
	}

	return
}
