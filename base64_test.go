package base64

import "testing"

func TestEncode(t *testing.T) {
	str := "http://www1.tc711.com/tool/BASE64.htm"
	code := Encode([]byte(str))
	if code != "aHR0cDovL3d3dzEudGM3MTEuY29tL3Rvb2wvQkFTRTY0Lmh0bQ==" {
		t.Error("encode error", code)
	}

}

func TestDecode(t *testing.T) {
	code := "aHR0cDovL3d3dzEudGM3MTEuY29tL3Rvb2wvQkFTRTY0Lmh0bQ=="
	str := Decode(code)
	if string(str) != "http://www1.tc711.com/tool/BASE64.htm" {
		t.Error("decode error", string(str))
	}
}

func TestChinese(t *testing.T) {
	str := "我是中国人 i am a chinese"
	code := Encode([]byte(str))
	t.Log(code)
	str1 := string(Decode(code))
	if str1 != str {
		t.Error("error")
	}
}