package tools

import (
    "encoding/base64"
    "github.com/forgoer/openssl"
)

func AESDecrypt(transStr string) string {
    key := []byte("loveme@nxflv@com")
    btransStr, _ := base64.StdEncoding.DecodeString(transStr)
    dst, _ := openssl.AesECBDecrypt(btransStr, key, openssl.PKCS7_PADDING)
    return string(dst)
}

func AESEncrypt(transStr string) string {
    key := []byte("yourme@nxflv@com")
    dst, _ := openssl.AesECBEncrypt([]byte(transStr), key, openssl.PKCS7_PADDING)
    return base64.StdEncoding.EncodeToString(dst) // yXVUkR45PFz0UfpbDB8/ew==
}
