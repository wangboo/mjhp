package mjhp

type RateAlgorithm interface {
	Calculate(req *JudgeReq, rst *RateResult) int // 计算额外倍率
}

type YiBinRate struct{}
type XueZhanRate struct{}

var (
	yiBinRate   = &YiBinRate{}
	xueZhanRate = &YiBinRate{}
)

// 	宜宾麻将
func (this *YiBinRate) Calculate(req *JudgeReq, rst *RateResult) int {
	// 清一色
	isQingYiSe := rst.Mask&RATE_MASK_QING_YI_SE == RATE_MASK_QING_YI_SE
	isQiDui := rst.Mask&RATE_MASK_7DUI == RATE_MASK_7DUI
	isDuiDuiHu := rst.Mask&RATE_MASK_DUIDUI_HU == RATE_MASK_DUIDUI_HU
	base := 0
	if isQingYiSe {
		if isQiDui {
			// 清七对 和 清龙七对叠加一样
			base = 5 + rst.GangCount
		} else if isDuiDuiHu {
			// 对对胡
			base = 5 + rst.GangCount
		} else {
			base = 3
		}
	} else {
		// 不是清一色
		if isQiDui {
			base = 3
		} else if isDuiDuiHu {
			base = 3
		}
	}
	// 8个癞子2番，4个1番
	if req.LzTotal >= 8 {
		rst.Mask |= RATE_MASK_LAIZI8
		base += 2
	} else if req.LzTotal >= 4 {
		rst.Mask |= RATE_MASK_LAIZI4
		base += 1
	}
	// 无听用2番，事件区无听用1番
	//if req.Events == nil || req.LzTotal-req.LzCount == 0 {
	//	base += 1
	//}
	if req.LzTotal == 0 {
		rst.Mask |= RATE_MASK_WU_TING_YONG
		base += 2
	}
	// 4. 宜宾麻将，本金加番判断
	benJin := req.GetBenJinByte()
	count := 0
	for _, m := range req.hands {
		if m == benJin {
			rst.Mask |= RATE_MASK_BENJIN
			count ++
		} else if m > benJin {
			break
		}
	}
	if count >= 3 {
		rst.Mask |= RATE_MASK_BENJIN
		base ++
	} else {
		if req.Events != nil {
			for _, e := range req.Events {
				if e.Key == req.BenJin {
					rst.Mask |= RATE_MASK_BENJIN
					base ++
				}
			}
		}
	}
	// 宜宾麻将暂无其他牌型
	return base
}

// 	血战麻将
func (this *XueZhanRate) Calculate(rst *RateResult) int {
	return 0
}
