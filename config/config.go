package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	ServerPort        string `json:"serverPort"`
	DBUrl             string `json:"dbUrl"`
	DBName            string `json:"dbName"`
	CacheHost         string `json:"cacheHost"`
	CachePwd          string `json:"cachePwd"`
	EncryptionKey     []byte `json:"encryptionKey"`
	AccessTokenLength int    `json:"accessTokenLength"`
	AccessTokenExpSec int    `json:"accessTokenExpSec"`
	CodeLength        int    `json:"codeLength"`
	CodeExpSec        int    `json:"codeExpSec"`
	AuthType          string `json:"authType"`
	VerifyTokenType   string `json:"verifyTokenType"`
	MFAType           string `json:"mfaType"`
	FirebaseCfg       string `json:"firebaseCfg"`
}

func NewAppConfig() *AppConfig {
	tlenStr := os.Getenv("ACCESS_TOKEN_LENGTH")
	tExpSecStr := os.Getenv("ACCESS_TOKEN_EXP_SEC")
	tlen, _ := strconv.Atoi(tlenStr)
	tExpSec, _ := strconv.Atoi(tExpSecStr)

	clenStr := os.Getenv("CODE_LENGTH")
	cExpSecStr := os.Getenv("CODE_EXP_SEC")
	clen, _ := strconv.Atoi(clenStr)
	cExpSec, _ := strconv.Atoi(cExpSecStr)

	return &AppConfig{
		ServerPort:        os.Getenv("SERVER_PORT"),
		DBUrl:             os.Getenv("DB_URL"),
		DBName:            os.Getenv("DB_NAME"),
		CacheHost:         os.Getenv("CACHE_HOST"),
		CachePwd:          os.Getenv("CACHE_PWD"),
		EncryptionKey:     []byte(os.Getenv("ENCRYPTION_KEY")),
		AccessTokenLength: tlen,
		AccessTokenExpSec: tExpSec,
		CodeLength:        clen,
		CodeExpSec:        cExpSec,
		AuthType:          os.Getenv("AUTH_TYPE"),
		VerifyTokenType:   os.Getenv("VERIFY_TOKEN_TYPE"),
		MFAType:           os.Getenv("MFA_TYPE"),
		FirebaseCfg:       os.Getenv("FIREBASE_CFG"),
	}
}
