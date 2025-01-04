package canis_rufus

/*
* @author: Chen Chiheng
* @date: 2025/1/3 22:25:09
* @description:
**/

type Validator interface {
	Validate(flag *Flag) (Validator, Task, error)
}
