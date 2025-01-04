package value

import "time"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 14:14:49
* @description:
**/

type TimeDuration struct {
	value time.Duration
}

func (v *TimeDuration) Set(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	v.value = duration
	return nil
}

func (v *TimeDuration) Type() string {
	return "time duration"
}

func (v *TimeDuration) String() string {
	return v.value.String()
}
