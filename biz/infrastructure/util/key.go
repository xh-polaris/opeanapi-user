package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"io"
	"strings"
	"time"
)

// 单例模式
var keyManager *KeyManager

func GetKeyManager() *KeyManager {
	if keyManager == nil {
		c, _ := config.NewConfig()
		keyManager = &KeyManager{
			Config: c,
		}
	}
	return keyManager
}

type KeyManager struct {
	Config *config.Config
}

/*
 * 使用AES对称加密，将密钥id、用户id、刷新时间作为参数拼接
 * 其中对密钥参数issue有特殊要求，要求为128字节，即16位ASCII码
 */

func (km *KeyManager) IssueKey(id string, userId string, freshTime time.Time) (string, error) {
	// 拼接参数
	data := fmt.Sprintf("%s|%s|%s", id, userId, freshTime.Format(time.RFC3339))

	//获取 AES 密钥
	sk := []byte(km.Config.Security.Issue)

	//创建AES加密器
	block, err := aes.NewCipher(sk)
	if err != nil {
		return "", err
	}

	// 创建GCM模式加密器
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密数据
	encrypted := gcm.Seal(nonce, nonce, []byte(data), nil)

	// 使用Base64编码
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (km *KeyManager) ParseKey(key string) (id string, userId string, freshTime time.Time, err error) {
	// 解码Base64编码
	encrypted, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", "", time.Time{}, err
	}

	// 提取AES密钥
	sk := []byte(km.Config.Security.Issue)

	// 创建AES解密器
	block, err := aes.NewCipher(sk)
	if err != nil {
		return "", "", time.Time{}, err
	}

	//创建GMC解密器
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", time.Time{}, err
	}

	// 提取nonce
	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return "", "", time.Time{}, errors.New("invalid key")
	}
	nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]

	// 解密数据
	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", "", time.Time{}, err
	}

	// 解析原始数据
	data := strings.Split(string(decrypted), "|")
	if len(data) != 3 {
		return "", "", time.Time{}, errors.New("invalid key")
	}

	// 解析时间
	freshTime, err = time.Parse(time.RFC3339, data[2])
	if err != nil {
		return "", "", time.Time{}, err
	}

	return data[0], data[1], freshTime, nil
}
