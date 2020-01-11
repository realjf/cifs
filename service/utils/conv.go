package utils

import "reflect"

func chinaNumberMap() map[string]string {
	return map[string]string{
		"零":"0",
		"壹":"1",
		"贰":"2",
		"叁":"3",
		"肆":"4",
		"伍":"5",
		"陆":"6",
		"柒":"7",
		"捌":"8",
		"玖":"9",
	}
}

func chinaNumberMap2() map[string]string {
	return map[string]string{
		"0":"零",
		"1":"壹",
		"2":"贰",
		"3":"叁",
		"4":"肆",
		"5":"伍",
		"6":"陆",
		"7":"柒",
		"8":"捌",
		"9":"玖",
	}
}


// 中文数字转阿拉伯数字
func C2ANumber(str string) string {
	if str == "" {
		return str
	}
	s := String(str).ToCharArray()
	r := NewString()
	for i := 0; i < len(s); i++ {
		if o, ok := chinaNumberMap()[s[i].ToString()]; ok {
			r.Append(o)
		}else{
			r.Append(s[i].ToString())
		}
	}
	return r.ToString()
}

// 阿拉伯数字转中文数字
func A2ChinaNumber(str string) string {
	if str == "" {
		return str
	}
	s := String(str).ToCharArray()
	r := NewString()
	table := chinaNumberMap2()
	for i := 0; i < len(s); i++ {
		if o, ok := table[s[i].ToString()]; ok {
			r.Append(o)
		}else{
			r.Append(s[i].ToString())
		}
	}
	return r.ToString()
}

//把类似slice的map转为slice
func MapToMapEntrySet(input interface{}) []EntrySet {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Map {
		return nil
	}
	keys := v.MapKeys()
	output := []EntrySet{}
	for i, l := 0, v.Len(); i < l; i++ {
		tmp := EntrySet{}
		tmp.SetKey(keys[i].Interface())
		tmp.SetValue(v.MapIndex(keys[i]).Interface())
		output = append(output, tmp)
	}
	return output
}
