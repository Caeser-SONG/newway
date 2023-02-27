package utils

import (
	"log"
	"strconv"

	"github.com/google/uuid"
)

func CreateUid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	uid_string := strconv.Itoa(int(uid.ID()))
	return uid_string, err
}
