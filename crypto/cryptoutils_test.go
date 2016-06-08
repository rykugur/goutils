package cryptoutils

import "testing"

func TestSha512(t *testing.T) {
	testStr := "thought is the arrow of time"

	hash, err := Sha512(testStr)

	if err != nil {
		t.Error("err was not nil", err)
	}

	if hash != "capkM5iwzXzbKnjGDhbNI_eoEkjSOppsGLGamU-43Vt5rioCe9Xi0qNgPrC1urqSUhEHmNRDnEjMsuLRGofiSg==" {
		t.Error("hash value didn't equal expected hash value")
	}
}
