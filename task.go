package canis_rufus

/*
* @author: Chen Chiheng
* @date: 2025/1/4 09:22:16
* @description:
**/

type Future struct {
}

type Task interface {
	Pipeline(Task) Task
}
