package utils 

import (
	"crypto/rand"
	"io"
	"fmt"
	"strconv"
	"time"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	//uuid[8] = uuid[8] & ^0xC0 | 0x80
	//uuid[6] = uuid[6] & ^0xF0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), err
}

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1000000, 10))
	return ts 
}
