package utils

import (
	"unicode"
	"unicode/utf8"
)

const (
	// ascii 表中可见字符从!开始，偏移位置为33
	DBC_CHAR_START = 33 // 半角
	// ascii 表中可见字符到～结束，偏移位置为126
	DBC_CHAR_END = 126
	// 全角对应于ascii表的可见字符从！开始，偏移值为65281
	SBC_CHAR_START = 65281
	// 全角对应于ascii表的可见字符到～结束，偏移值为65374
	SBC_CHAR_END = 65374

	// ASCII表中除空格外的可见字符与对应的全角字符的相对偏移
	CONVERT_STEP = 65248 // 全角半角转换间隔
	// 全角空格的值，它没有遵从与ASCII的相对偏移，必须单独处理
	SBC_SPACE = 12288 // 全角空格

	// 半角空格的值，在ASCII中为32(Decimal)
	DBC_SPACE = ' ' // 半角空格
)

// 半角字符 -> 全角字符 转换
// 只处理空格，!到~之间的字符，忽略其他
func S2DConvertString(src string) string {
	if src == "" {
		return src
	}
	var buf String = NewString()
	var ca []Char = String(src).ToCharArray()
	for i := 0; i < len(ca); i++ {
		if ca[i].ToInt() == DBC_SPACE {
			buf.Append(Char(SBC_SPACE).ToString())
		}else if (ca[i].ToInt() >= DBC_CHAR_START) && (ca[i].ToInt() <= DBC_CHAR_END) {
			buf.Append(Char(ca[i].ToInt() + CONVERT_STEP).ToString())
		}else {
			buf.Append(ca[i].ToString())
		}
	}

	return buf.ToString()
}

// 半角->全角转换
func S2DConvertChar(src Char) string {
	var r int = src.ToInt()
	if src == DBC_SPACE {// 如果是半角空格，直接用全角空格替代
		r = SBC_SPACE
	}else if (src.ToInt() >= DBC_CHAR_START) && (src <= DBC_CHAR_END) {// 字符是!到~之间的可见字符
		r = src.ToInt() + CONVERT_STEP
	}

	return Char(r).ToString()
}

// 全角字符->半角字符转换
// 只处理全角的空格，全角！到全角～之间的字符，忽略其他
func D2SConvertString(src string) string {
	if src == "" {
		return src
	}
	var buf String = NewString()
	var ca []Char = String(src).ToCharArray()
	for i := 0; i < len(ca); i++ {
		if ca[i].ToInt() >= SBC_CHAR_START && ca[i].ToInt() <= SBC_CHAR_END {
			// 如果位于全角！到全角～区间内
			buf.Append(Char(ca[i].ToInt() - CONVERT_STEP).ToString())
		}else if ca[i].ToInt() == SBC_SPACE {
			// 如果是全角空格
			buf.Append(Char(DBC_SPACE).ToString())
		}else {
			// 不处理全角空格，全角！到全角～区间外的字符
			buf.Append(ca[i].ToString())
		}
	}
	return buf.ToString()
}

// 全角转换半角
func D2SConvertChar(src Char) string {
	var r int = src.ToInt()
	if src.ToInt() >= SBC_CHAR_START && src <= SBC_CHAR_END {
		r = src.ToInt() - CONVERT_STEP
	}else if src.ToInt() == SBC_SPACE {
		r = DBC_SPACE
	}
	return Char(r).ToString()
}

// 将文本划分成字元
func SplitTextToWords(text Text) []Text {
	output := make([]Text, 0, len(text)/3)
	current := 0
	inAlphanumeric := true
	alphanumericStart := 0
	for current < len(text) {
		r, size := utf8.DecodeRune(text[current:])
		if size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			// 当前是拉丁字母或数字（非中日韩文字）
			if !inAlphanumeric {
				alphanumericStart = current
				inAlphanumeric = true
			}
		} else {
			if inAlphanumeric {
				inAlphanumeric = false
				if current != 0 {
					output = append(output, ToLower(text[alphanumericStart:current]))
				}
			}
			output = append(output, text[current:current+size])
		}
		current += size
	}

	// 处理最后一个字元是英文的情况
	if inAlphanumeric {
		if current != 0 {
			output = append(output, ToLower(text[alphanumericStart:current]))
		}
	}

	return output
}

// 将英文词转化为小写
func ToLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}