package value

import "strconv"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 19:28:58
* @description:
**/

type BoolValue struct {
	value bool
}

func (v *BoolValue) Set(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	v.value = b
	return nil
}

func (v *BoolValue) Type() string {
	return "bool"
}

func (v *BoolValue) String() string {
	return strconv.FormatBool(v.value)
}
