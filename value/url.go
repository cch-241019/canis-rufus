package value

import "net/url"

/*
* @author: Chen Chiheng
* @date: 2025/1/4 14:22:43
* @description:
**/

type UrlValue struct {
	value string
}

func (v *UrlValue) Set(value string) error {
	u, err := url.ParseRequestURI(value)
	if err != nil {
		return err
	}
	v.value = u.String()
	return nil
}

func (v *UrlValue) Type() string {
	return "url"
}

func (v *UrlValue) String() string {
	return v.value
}
