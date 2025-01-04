package value

import (
	"fmt"
	"os"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/4 13:53:29
* @description:
**/

type DIrValue struct {
	value string
}

func (v *DIrValue) Set(value string) error {
	fileInfo, err := os.Stat(value)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", value)
		}
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", value)
	}
	v.value = value
	return nil
}

func (v *DIrValue) Type() string {
	return "dir"
}

func (v *DIrValue) String() string {
	return v.value
}
