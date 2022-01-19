package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// GenerateID 基于时间戳生成不同时间的唯一ID
func GenerateID() string {
	s := strconv.FormatInt(time.Now().UnixNano(), 10)
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
