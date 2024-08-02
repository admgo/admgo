package random

import (
	"github.com/google/uuid"
)

// GenerateSessionID 生成一个新的 SessionID
func GenerateSessionID() string {
	// 使用 UUID 生成 SessionID
	sessionID := uuid.New().String()
	return sessionID
}
