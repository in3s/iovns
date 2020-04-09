package iovnsd

import (
	"bytes"
	"fmt"
	"time"
)

// SecondsToTime converts unix seconds to time
func SecondsToTime(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}

// TimeToSeconds converts a time.Time to unix seconds timestamp
func TimeToSeconds(t time.Time) int64 {
	return t.Unix()
}

// GetAccountKey returns an account key in the form of string
// given the domain name and the account name
func GetAccountKey(domain, name string) string {
	return fmt.Sprintf("%s*%s", domain, name)
}

// SplitAccountKey takes an account key and splits it
// into domain name and account name, panics on nil keys.
func SplitAccountKey(key []byte) (domainName, accountName string) {
	resp := bytes.Split(key, []byte("*"))
	return string(resp[0]), string(resp[1])
}