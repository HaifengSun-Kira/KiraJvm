package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goString string) *Object {
	if internedStr, ok := internedStrings[goString]; ok {
		return internedStr
	}

	chars := stringToUtf16(goString)
	jChars := &Object{loader.LoadClass("[C"), chars, nil}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goString] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) // utf32
	return utf16.Encode(runes)
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)  // utf8
	return string(runes)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jStr
	return jStr
}