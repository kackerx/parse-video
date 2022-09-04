package tools

import (
    "fmt"
    "testing"
)

func TestAes(t *testing.T) {
    encodeStr := []byte("A0Jy37StW3zNwasuz8eYxhGZQIM3LkUwwOTF0gRTNJ3qPVOELdveEFrMeoVym3LuDwwHoNgbAKqe6RRJeg6tzKuiEb8W3/z0B/4ln1Jf8Os=")
    //key := []byte("loveme@nxflv@com")
    decrypt := AESDecrypt(string(encodeStr))
    fmt.Println(decrypt)

    res := AESEncrypt("https://www.iqiyi.com/v_xzqbkelysg.html")

    fmt.Println(res)
}
