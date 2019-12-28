package aes

import "testing"

func TestAES(t *testing.T) {
	orig := "Hello World"
	key := "piLQJ0unsBPxZx7+QEn1/Q=="
	// fmt.Println("原文：", orig)

	encryptCode := AesEncrypt(orig, key)
	// fmt.Println("密文：", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key)
	// fmt.Println("解密结果：", decryptCode)
	if decryptCode != orig {
		t.Error("AES encrypt or decrypt have problem.")
	}
}
