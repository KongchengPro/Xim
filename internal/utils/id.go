package utils

import (
	"encoding/hex"
	"github.com/google/uuid"
)

// GenerateID 基于时间戳生成不同时间的唯一ID
func GenerateID() string {
	hash := uuid.New()
	return hex.EncodeToString(hash[:])
}
