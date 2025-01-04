package value

/*
* @author: Chen Chiheng
* @date: 2025/1/4 13:52:48
* @description:
**/

type Value interface {
	Set(value string) error
	Type() string
	String() string
}


