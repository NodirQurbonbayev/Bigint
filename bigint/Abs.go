package bigint

import (
	"errors"
	"strings"
)


func ValidateNum(num string) error {
	allowed :="123456789"
	if !strings.Contains("+"+allowed,string(num[0])) {
		return errors.New("ErrorBadInput")
	}
	for i:=1; i<len(num); i++ {
		if !strings.Contains("+"+allowed,string(num[i])){
		// return ErrorBadInput
		}
	}
	return nil

}
