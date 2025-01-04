package value

import "strconv"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 18:42:35
* @description:
**/

type Int8Value struct {
	value int8
}

func (v *Int8Value) Set(value string) error {
	i, err := strconv.ParseInt(value, 0, 8)
	if err != nil {
		return err
	}
	v.value = int8(i)
	return nil
}

func (v *Int8Value) Type() string {
	return "int8"
}

func (v *Int8Value) String() string {
	return strconv.FormatInt(int64(v.value), 10)
}
