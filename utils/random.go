package utils

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomNumber(n int) string {
	var alphabet string = "1234567890"
	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(5)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(10))
}

func RandomFileName(file *multipart.FileHeader) string {
	fileName := filepath.Base(file.Filename)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return fmt.Sprintf("%s-%d-%s", fileNameWithoutExt, time.Now().UnixNano(), RandomString(8)+filepath.Ext(file.Filename))
}

func RandomTransactionTypes() string {
	var trxType = []string{"TRANSFER", "TOPUP", "WITHDRAWAL", "PAYMENT"}
	return trxType[rand.Intn(len(trxType))]
}
