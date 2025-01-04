package value

import "strconv"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 19:26:11
* @description:
**/

type Float32Value struct {
	value float32
}

func (v *Float32Value) Set(value string) error {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	v.value = float32(f)
	return nil
}

func (v *Float32Value) Type() string {
	return "float32"
}

func (v *Float32Value) String() string {
	return strconv.FormatFloat(float64(v.value), 'g', -1, 32)
}
