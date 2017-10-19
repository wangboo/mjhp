package mjhp

import (
	"log"
)

type RateAlgorithm interface {
	Calculate(req *JudgeReq, rst *RateResult) int // 计算额外倍率
}

type YiBinRate struct{}    // 宜宾麻将
type XueZhanRate struct{}  // 血战、三人两房麻将
type NanChongRate struct{} // 南充麻将,不需要缺一门，如果缺一门，加一番

var (
	yiBinRate   = &YiBinRate{}
	xueZhanRate = &XueZhanRate{}
)

// 	宜宾麻将
func (this *YiBinRate) Calculate(req *JudgeReq, rst *RateResult) int {
	//printMj(req.handsWithLz)
	// 清一色
	isQingYiSe := rst.Mask&RATE_MASK_QING_YI_SE == RATE_MASK_QING_YI_SE
	isQiDui := rst.Mask&RATE_MASK_7DUI == RATE_MASK_7DUI
	isDuiDuiHu := rst.Mask&RATE_MASK_DUIDUI_HU == RATE_MASK_DUIDUI_HU
	// 先移除掉特征码
	if isQingYiSe {
		rst.Mask &= RATE_MASK_QING_YI_SE ^ RATE_MASK
	}
	if isQiDui {
		rst.Mask &= RATE_MASK_7DUI ^ RATE_MASK
	}
	if isDuiDuiHu {
		rst.Mask &= RATE_MASK_DUIDUI_HU ^ RATE_MASK
	}

	// 7对逻辑修复, 7对癞子变的杠不加番，但是要算牌型（龙7对，双龙七对）!req.IsChaJiao() && 略去
	//log.Println("before isQiDui = ", isQiDui, ", rst.GangCount = ", rst.GangCount)
	if isQiDui {
		rst.GangCount = getGangWithLz(req.handsWithLz)
	}
	//log.Println("after isQiDui = ", isQiDui, ", rst.GangCount = ", rst.GangCount)

	base := 0
	if isQingYiSe {
		if isQiDui {
			if rst.GangCount >= 2 {
				// 清双龙七对
				base = 6
				rst.Mask |= RATE_MASK_QING_7DUI_LONG2
			} else if rst.GangCount == 1 {
				// 清一色龙七对
				base = 6
				rst.Mask |= RATE_MASK_QING_7DUI_LONG
			} else {
				// 清七对
				base = 6
				rst.Mask |= RATE_MASK_QING_7DUI
			}
			rst.GangCount = 0
		} else if isDuiDuiHu {
			// 清对
			base = 6
			rst.Mask |= RATE_MASK_QING_DUI
		} else {
			// 清一色
			rst.Mask |= RATE_MASK_QING_YI_SE
			base = 3
		}
	} else {
		// 不是清一色
		if isQiDui {
			if rst.GangCount >= 3 {
				rst.Mask |= RATE_MASK_LONG_7DUI3
				base = 6
			} else if rst.GangCount >= 2 {
				rst.Mask |= RATE_MASK_LONG_7DUI2
				base = 5
			} else if rst.GangCount == 1 {
				rst.Mask |= RATE_MASK_LONG_7DUI
				base = 4
			} else {
				rst.Mask |= RATE_MASK_7DUI
				base = 3
			}
			rst.GangCount = 0
		} else if isDuiDuiHu {
			base = 3
			rst.Mask |= RATE_MASK_DUIDUI_HU
		}
	}
	if rst.Mask&RATE_MASK_JIN_GOU == RATE_MASK_JIN_GOU {
		// 金钩胡
		base += 4
		rst.Mask |= RATE_MASK_DUIDUI_HU
	}
	// 以上为基本牌型，以下为附加番
	// 不是花猪大叫，0番设为1番
	if rst.Mask&RATE_MASK_HUA_ZU != RATE_MASK_HUA_ZU && rst.Mask&RATE_MASK_DA_JIAO != RATE_MASK_DA_JIAO {
		if base == 0 {
			base = 1
		}
	}
	base += rst.GangCount
	// 8个癞子2番，4个1番， 如果是最后查叫，默认添加一个癞子，这时候计算癞子数量要去掉
	lzCount := req.LzTotal
	if req.IsChaJiao() {
		lzCount -= 1
	}
	if lzCount >= 8 {
		rst.Mask |= RATE_MASK_LAIZI8
		base += 2
	}
	// 去掉四个癞子
	//else if lzCount >= 4 {
	//	rst.Mask |= RATE_MASK_LAIZI4
	//	base += 1
	//}
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
				if !e.IsFei() && e.Key == req.BenJin {
					rst.Mask |= RATE_MASK_BENJIN
					base ++
				}
			}
		}
	}
	// 宜宾麻将暂无其他牌型
	return base
}

func getGangWithLz(hands []byte) (gangCount int) {
	lenOfHands := len(hands)
	for i := 0; i < lenOfHands-3; {
		//log.Printf("i = %v, hands[i] = %v, hands[i+1] = %v, hands[i+2] = %v, hands[i+3] = %v\n",
		//i, hands[i], hands[i+1], hands[i+2], hands[i+3])
		if hands[i] == hands[i+1] && hands[i] == hands[i+2] && hands[i] == hands[i+3] {
			//log.Println("i = ", gangCount)
			gangCount++
			i += 4
		} else {
			i++
		}
	}
	return
}

// 	血战麻将
func (this *XueZhanRate) Calculate(req *JudgeReq, rst *RateResult) int {
	isQingYiSe := rst.Mask&RATE_MASK_QING_YI_SE == RATE_MASK_QING_YI_SE
	isQiDui := rst.Mask&RATE_MASK_7DUI == RATE_MASK_7DUI
	isDuiDuiHu := rst.Mask&RATE_MASK_DUIDUI_HU == RATE_MASK_DUIDUI_HU
	if isQingYiSe {
		rst.Mask &= RATE_MASK_QING_YI_SE ^ RATE_MASK
	}
	if isQiDui {
		rst.Mask &= RATE_MASK_7DUI ^ RATE_MASK
	}
	if isDuiDuiHu {
		rst.Mask &= RATE_MASK_DUIDUI_HU ^ RATE_MASK
	}

	base := 0
	if isQingYiSe {
		if isQiDui {
			if rst.GangCount == 2 {
				// 清双龙七对
				rst.Mask |= RATE_MASK_QING_7DUI_LONG2
			} else if rst.GangCount == 1 {
				// 清一色龙七对
				rst.Mask |= RATE_MASK_QING_7DUI_LONG
			} else {
				// 清七对
				rst.Mask |= RATE_MASK_QING_7DUI
			}
			base = 4
		} else if isDuiDuiHu {
			// 清对
			if req.IsDdh2Rate() {
				base = 4
			} else {
				base = 3
			}
			rst.Mask |= RATE_MASK_QING_DUI
		} else {
			// 清一色
			rst.Mask |= RATE_MASK_QING_YI_SE
			base = 2
		}
	} else {
		// 不是清一色
		if isQiDui {
			if rst.GangCount >= 2 {
				rst.Mask |= RATE_MASK_LONG_7DUI2
				base = 4
			} else if rst.GangCount == 1 {
				rst.Mask |= RATE_MASK_LONG_7DUI
				base = 3
			} else {
				rst.Mask |= RATE_MASK_7DUI
				base = 2
			}
			rst.GangCount = 0
		} else if isDuiDuiHu {
			if req.IsDdh2Rate() {
				base = 2
			} else {
				base = 1
			}
			rst.Mask |= RATE_MASK_DUIDUI_HU
		}
	}
	// 将对, 全是2，5，8对 ()
	if req.IsYJJD() {
		if isJiangDui(req) {
			if isQiDui {
				rst.Mask |= RATE_MASK_JIANG_DUI
				base += 1
			} else {
				rst.Mask |= RATE_MASK_JIANG_DUI
				base += 3
				if isDuiDuiHu {
					if req.IsDdh2Rate() {
						base -= 2
					} else {
						base -= 1
					}
					// 有将对，擦除对对胡标识
					rst.Mask &= RATE_MASK_DUIDUI_HU ^ RATE_MASK
				} else if rst.Mask&RATE_MASK_JIN_GOU == RATE_MASK_JIN_GOU {
					// 金钩胡, 擦除对对胡标识
					rst.Mask &= RATE_MASK_DUIDUI_HU ^ RATE_MASK
				}
			}
		}
		if isYaoJiu(req) {
			rst.Mask |= RATE_MASK_YAO_JIU
			base += 3
		}
	}
	if rst.Mask&RATE_MASK_JIN_GOU == RATE_MASK_JIN_GOU {
		if rst.Mask&RATE_MASK_JIANG_DUI == RATE_MASK_JIANG_DUI {
			// 将对，加一番
			base += 1
		} else {
			if req.IsDdh2Rate() {
				base += 2
			} else {
				base += 1
			}
			base += 1
			rst.Mask |= RATE_MASK_DUIDUI_HU
		}
	}
	// 门清中张（7对不算门清）
	if !isQiDui && req.IsMenQing() && isMenQing(req) {
		rst.Mask |= RATE_MASK_MEN_QING
		base += 1
	}
	// 中张  （将对不算中张）
	if rst.Mask&RATE_MASK_JIANG_DUI != RATE_MASK_JIANG_DUI && req.IsZhongZhang() && isZhongZhang(req) {
		rst.Mask |= RATE_MASK_ZHONG_ZHANG
		base += 1
	}
	// 夹心五
	//log.Println("req.IsJX5() = ", req.IsJX5())
	if req.IsJX5() && (req.JudgeMj == 4 || req.JudgeMj == 13 || req.JudgeMj == 22) && isJX5(req) {
		//log.Println("夹心5 + 1番")
		rst.Mask |= RATE_MASK_JIA_XIN5
		base += 1
	}
	return base + rst.GangCount
}

func isMenQing(req *JudgeReq) bool {
	if req.Events == nil {
		return true
	}
	for _, e := range req.Events {
		if !e.IsAnGang() {
			return false
		}
	}
	return true
}

// 中张
func isZhongZhang(req *JudgeReq) bool {
	if req.Events != nil {
		for _, e := range req.Events {
			v := e.Key % 9
			if v == 0 || v == 8 {
				return false
			}
		}
	}
	for _, m := range req.handsWithLz {
		v := (int(m) % 16) % 9
		if v == 1 || v == 0 {
			return false
		}
	}
	return true
}

// 幺九
func isYaoJiu(req *JudgeReq) bool {
	if req.Events != nil {
		for _, e := range req.Events {
			v := e.Key % 9
			if v != 0 && v != 8 {
				return false
			}
		}
	}
	// 使用幺九表
	v := bytesToStr(transFrom2Color(req.handsWithLz))
	printMj(req.handsWithLz)
	_, ok := lys[v]
	log.Println("ok = ", ok)
	return ok
}

// 将对
func isJiangDui(req *JudgeReq) bool {
	if req.Events != nil {
		for _, e := range req.Events {
			// 1，4，7，10 ...
			if e.Key%3 != 1 {
				return false
			}
		}
	}
	for _, m := range req.handsWithLz {
		if (m%16)%3 != 2 {
			return false
		}
	}
	return true
}

func isJX5(req *JudgeReq) bool {
	lenOfHands := len(req.handsWithLz)
	if lenOfHands < 5 {
		return false
	}
	if req.JudgeMj != 4 && req.JudgeMj != 13 && req.JudgeMj != 22 {
		return false
	}
	pre, after := tiles[req.JudgeMj-1], tiles[req.JudgeMj+1]
	delPre, delCenter, delAfter := -1, -1, -1
	for idx, v := range req.handsWithLz {
		if delPre == -1 && v == pre {
			delPre = idx
		} else if delAfter == -1 && v == after {
			delAfter = idx
		} else if delCenter == -1 && v == tiles[req.JudgeMj] {
			delCenter = idx
		}
	}
	//log.Printf("delPre = %d, delCenter = %d, delAfter = %d \n", delPre, delCenter, delAfter)
	if delPre == -1 || delAfter == -1 {
		return false
	}
	jx5Hands := make([]byte, 0, lenOfHands)
	for idx, v := range req.handsWithLz {
		if idx == delPre || idx == delAfter || idx == delCenter {
			continue
		}
		jx5Hands = append(jx5Hands, v)
	}
	printMj(jx5Hands)
	jx5Req := &JudgeReq{
		hands:       jx5Hands,
		handsWithLz: jx5Hands,
		Mask:        0x01,
		colorCount:  req.colorCount,
	}
	jx5Rst := judgeHu(jx5Req)
	return jx5Rst.Result
}

// 南充麻将
func (this *NanChongRate) Calculate(req *JudgeReq, rst *RateResult) (rate int) {
	isQingYiSe := rst.Mask&RATE_MASK_QING_YI_SE == RATE_MASK_QING_YI_SE
	isQiDui := rst.Mask&RATE_MASK_7DUI == RATE_MASK_7DUI
	isDuiDuiHu := rst.Mask&RATE_MASK_DUIDUI_HU == RATE_MASK_DUIDUI_HU
	// 先移除掉特征码
	if isQingYiSe {
		rst.Mask &= RATE_MASK_QING_YI_SE ^ RATE_MASK
	}
	if isQiDui {
		rst.Mask &= RATE_MASK_7DUI ^ RATE_MASK
	}
	if isDuiDuiHu {
		rst.Mask &= RATE_MASK_DUIDUI_HU ^ RATE_MASK
	}
	if isQingYiSe {
		rate = 6
		rst.Mask |= RATE_MASK_QING_YI_SE
	} else if isQiDui {
		rate = 6
		rst.Mask |= RATE_MASK_7DUI
	} else if isDuiDuiHu && isNanChongDuiDuiHu(req) {
		rate = 4
		rst.Mask |= RATE_MASK_DDZ
	}
	// 七对和对对胡不算一般高
	if !isQiDui && !isDuiDuiHu && isYiBanGao(req.hands) {
		rate += 1
	}
	// 缺一门，加一番
	if !isQingYiSe && req.colorCount == 1 {
		rate += 1
		rst.Mask |= RATE_MASK_QUEMEN
	}
	return rate
}

// 南充对对胡，单吊才算对对胡
func isNanChongDuiDuiHu(req *JudgeReq) bool {
	huByte := mjIntToByte(req.JudgeMj)
	count := 0
	for _, v := range req.hands {
		if huByte == v {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return count == 2
}

// 一般高, 只计算一次, 2,2,3,3,4,4
func isYiBanGao(hands []byte) bool {
	return true 
}
