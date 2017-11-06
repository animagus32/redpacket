package utils

import "container/list"

type RedpacketItem struct {

	Id uint
	Timeout	int64
}


type RedpacketRetreiver struct {
	List list.List
}

func (r *RedpacketRetreiver) Insert(id uint,timeout int64)  {
	timeout += 10
	r.List.PushBack(RedpacketItem{id,timeout})
}

