package utils

import "errors"

func IsIgnorableError(err error, ignoreList []string) bool {
	for _, ignore := range ignoreList {
		if errors.Is(err, errors.New(ignore)) {
			return true
		}
	}
	return false
}

func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return errors.New(message + ": " + err.Error())
}
