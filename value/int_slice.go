package value

import (
	"strconv"
	"strings"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/4 18:47:51
* @description:
**/

type IntSliceValue struct {
	value   []int
	changed bool
}

func (v *IntSliceValue) Set(value string) error {
	elems := strings.Split(value, ",")
	out := make([]int, len(elems))
	for i, e := range elems {
		var err error
		out[i], err = strconv.Atoi(e)
		if err != nil {
			return err
		}
	}
	if !v.changed {
		v.value = out
	} else {
		v.value = append(v.value, out...)
	}
	v.changed = true
	return nil
}

func (v *IntSliceValue) Type() string {
	return "int slice"
}

func (v *IntSliceValue) String() string {
	out := make([]string, len(v.value))
	for i, e := range v.value {
		out[i] = strconv.Itoa(e)
	}
	return "[" + strings.Join(out, ",") + "]"
}
