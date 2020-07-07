package data

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
)

func getMd5Hash(obj interface{}) [16]byte {
	// convert object to json string
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	// hash json string using md5 algo
	return md5.Sum(bytes)
}

func getMd5HashStr(obj interface{}) string {
	// convert object to json string
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	// hash json string using md5 algo
	hasher := md5.New()
	hasher.Write(bytes)
	// hash := md5.Sum(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
