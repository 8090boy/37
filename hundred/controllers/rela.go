package controllers

import (
	model "hundred/models"
)

// 关系户推荐的人员数量级下层人员推荐总和
func findRecommandInfo(relaid int64, refMon *model.Monad) (firstCount, secondCount int) {
	tmpRala := &model.Relational{}
	relas := tmpRala.FindByReferrer(relaid)
	firstCount = 0
	secondCount = 0
	if len(relas) == 0 {
		return
	}
	for _, relaA := range relas {
		if relaA == nil {
			continue
		}
		if mainMonadIsCommon(relaA, refMon) {
			firstCount++
		}
		refRefRelas, no := tmpRala.FindRecommended(relaA.Id)
		if no == 0 {
			continue
		}

		for _, rela := range refRefRelas {
			if mainMonadIsCommon(rela, refMon) {
				secondCount++
			}
		}
	}

	tmpRala = nil
	return
}

// 根据关系户查找主单是否正常
func mainMonadIsCommon(rela *model.Relational, refMon *model.Monad) bool {
	if rela == nil {
		return false
	}
	if rela.Id == 0 || rela.CurrentMonad == 0 {
		return false
	}
	mainMonadRef := new(model.Monad).ById(rela.CurrentMonad)

	if refMon != nil {
		if mainMonadRef.Id == refMon.Id {
			return true
		}
	}

	if mainMonadRef != nil {
		if mainMonadRef.State > 0 && mainMonadRef.State != 3 {
			return true
		}
	}
	return false
}
