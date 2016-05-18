package base64

var BASE64_CODE      []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=")
var BASE64_SAFE_CODE []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_=")
const BIT_COUNT = 6

func _encode(data []byte, base64Code []byte)string {
	buffer := []byte{}
	var lastCount uint = 0
	var lastNum uint = 0
	var num uint = 0
	for _, code := range data {
		num = uint(uint( uint( code >> uint(8-BIT_COUNT+lastCount)  ) | lastNum) & 0x3F)
		buffer = append(buffer, BASE64_CODE[num])
		lastCount = uint(8-BIT_COUNT+lastCount)
		lastNum = uint( uint(code & (1<<lastCount-1)) << (BIT_COUNT-lastCount))

		if lastCount == 6 {
			buffer = append(buffer, BASE64_CODE[lastNum])
			lastCount = 0
			lastNum = 0
		}
	}

	if(lastCount>0){
		buffer = append(buffer, BASE64_CODE[lastNum])
	}
	if(lastCount==4){
		buffer = append(buffer, '=')
	} else if(lastCount==2){
		buffer = append(buffer, []byte("==")...)
	}
	return string(buffer)
}



func getCode(code byte)uint {
	for index, c := range BASE64_CODE{
		if c == code {
			return uint(index)
		}
	}
	panic("index error")
}

func _decode(str string, base64Code []byte)[]byte {
	buffer := []byte{}
	var lastCount uint8 = 0 //余下的长度
	var lastNum uint8 = 0   //余下的内容
	var num uint8 = 0

	data := []byte(str)
	for _, c := range data {
		code := getCode(c)
		if lastCount >0 {
			leftOff := 6-(8-lastCount)
			num = uint8(uint8(code >> leftOff) | lastNum)

			if(code > 63 && lastNum <=0){
				break
			}
			buffer = append(buffer, byte(num))

			lastCount = leftOff
			lastNum = uint8(code << (8-lastCount))
		} else {
			lastCount = 6
			lastNum = uint8(code << 2)
		}
	}
	return buffer
}

func Encode(data []byte) string {
	return _encode(data, BASE64_CODE)
}

func Decode(str string)[]byte {
	return _decode(str,BASE64_CODE)
}

// '+' -> '-',  '/' -> '_'
func UrlSafeEncode(data []byte) string {
	return _encode(data, BASE64_SAFE_CODE)
}

// '+' -> '-',  '/' -> '_'
func UrlSafeDecode(str string)[]byte {
	return _decode(str,BASE64_SAFE_CODE)
}