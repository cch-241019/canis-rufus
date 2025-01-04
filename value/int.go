package value

import "strconv"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 18:39:26
* @description:
**/

type IntValue struct {
	value int
}

func (v *IntValue) Set(value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	v.value = i
	return nil
}

func (v *IntValue) Type() string {
	return "int"
}

func (v *IntValue) String() string {
	return strconv.Itoa(v.value)
}
