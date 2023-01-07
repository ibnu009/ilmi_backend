package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nama") {
		return errors.New("Nama Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "Password") {
		return errors.New("Password Salah")
	}
	return errors.New("Email Salah")
}
