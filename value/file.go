package value

import (
	"fmt"
	"os"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/4 14:18:44
* @description:
**/

type FileValue struct {
	value string
}

func (v *FileValue) Set(value string) error {
	FileInfo, err := os.Stat(value)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", value)
		}
		return err
	}
	if FileInfo.IsDir() {
		return fmt.Errorf("%s is a directory", value)
	}
	v.value = value
	return nil
}

func (v *FileValue) Type() string {
	return "file"
}

func (v *FileValue) String() string {
	return v.value
}
