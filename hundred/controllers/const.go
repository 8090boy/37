package controllers

var INCOME []int64 = []int64{0, 200, 200, 500, 1500, 5000, 0, 0, 0, 0}

const (
	RELA_STATUS_UNBORN  int = iota
	RELA_STATUS_NORMAL      //正常 normal 1
	RELA_STATUS_FREEZE      // 冻结状态 2
	RELA_STATUS_DISCARD     // 非正常出局 3
	RELA_STATUS_FOUR        // 未完成任务冻结 4
	RELA_STATUS_FIVE
	RELA_STATUS_SIX
	RELA_STATUS_SEVEN
	RELA_STATUS_EIGHT
	RELA_STATUS_Retired // 正常出局 9
	RELA_STATUS_TEN
)
