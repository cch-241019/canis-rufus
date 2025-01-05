package value

import (
	"strconv"
	"strings"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/4 19:29:05
* @description:
**/

type BoolSliceValue struct {
	value   []bool
	changed bool
}

func (v *BoolSliceValue) Set(value string) error {
	boolSlice := strings.Split(value, ",")

	out := make([]bool, 0, len(boolSlice))
	for _, bs := range boolSlice {
		b, err := strconv.ParseBool(bs)
		if err != nil {
			return err
		}
		out = append(out, b)
	}

	if !v.changed {
		v.value = out
	} else {
		v.value = append(v.value, out...)
	}
	v.changed = true
	return nil
}

func (v *BoolSliceValue) Type() string {
	return "bool slice"
}

func (v *BoolSliceValue) String() string {
	strSlice := make([]string, len(v.value))
	for i, bs := range v.value {
		strSlice[i] = strconv.FormatBool(bs)
	}
	return "[" + strings.Join(strSlice, ",") + "]"
}
