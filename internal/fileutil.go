package internal

import (
	"bytes"
	"errors"
)

func StripPostHeader(content []byte) (stripped []byte, header []byte, err error) {
	delimiter := []byte("@@@@@@@@@@")
	parts := bytes.Split(content, delimiter)
	if len(parts) > 2 {
		err := errors.New("Invalid format. Only one header section allowed.")
		return nil, nil, err
	} else if len(parts) < 2 {
		return parts[0], nil, nil
	} else {
		return parts[1], parts[2], nil
	}
}