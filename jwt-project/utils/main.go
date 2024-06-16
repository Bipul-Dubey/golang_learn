package utils

import "time"

func GetTokenExpireTime() time.Time {
	return time.Now().Add(24 * time.Hour)
}
