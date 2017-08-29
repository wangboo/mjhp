package mjhp

import (
	"log"
	"fmt"
)

const (
	RATE_MASK_7DUI         = 0x001 // 7对		0
	RATE_MASK_MEN_QING     = 0x002 // 门清		1
	RATE_MASK_ZHONG_ZHANG  = 0x004 // 中张		2
	RATE_MASK_JIA_XIN5     = 0x008 // 夹心5		3
	RATE_MASK_JIN_GOU      = 0x010 // 金钩胡		4
	RATE_MASK_DUIDUI_HU    = 0x020 // 对对胡		5
	RATE_MASK_QING_YI_SE   = 0x040 // 清一色		6
	RATE_MASK_BENJIN       = 0x080 // 本金暗刻	7
	RATE_MASK_LAIZI8       = 0x100 // 8个癞子	8
	RATE_MASK_LAIZI4       = 0x200 // 4个癞子	9
	RATE_MASK_WU_TING_YONG = 0x400 // 无听用		10
)

func rateToString(v int) string {
	var str string
	if v&RATE_MASK_7DUI == RATE_MASK_7DUI {
		str = fmt.Sprintf(",%s", "7对")
	}
	if v&RATE_MASK_MEN_QING == RATE_MASK_MEN_QING {
		str = fmt.Sprintf("%s,%s", str, "门清")
	}
	if v&RATE_MASK_ZHONG_ZHANG == RATE_MASK_ZHONG_ZHANG {
		str = fmt.Sprintf("%s,%s", str, "中张")
	}
	if v&RATE_MASK_JIA_XIN5 == RATE_MASK_JIA_XIN5 {
		str = fmt.Sprintf("%s,%s", str, "夹心5")
	}
	if v&RATE_MASK_JIN_GOU == RATE_MASK_JIN_GOU {
		str = fmt.Sprintf("%s,%s", str, "金钩胡")
	}
	if v&RATE_MASK_DUIDUI_HU == RATE_MASK_DUIDUI_HU {
		str = fmt.Sprintf("%s,%s", str, "对对胡")
	}
	if v&RATE_MASK_QING_YI_SE == RATE_MASK_QING_YI_SE {
		str = fmt.Sprintf("%s,%s", str, "清一色")
	}
	if v&RATE_MASK_BENJIN == RATE_MASK_BENJIN {
		str = fmt.Sprintf("%s,%s", str, "本金暗刻")
	}
	if v&RATE_MASK_LAIZI8 == RATE_MASK_LAIZI8 {
		str = fmt.Sprintf("%s,%s", str, "8个癞子")
	}
	if v&RATE_MASK_LAIZI4 == RATE_MASK_LAIZI4 {
		str = fmt.Sprintf("%s,%s", str, "4个癞子")
	}
	if v&RATE_MASK_WU_TING_YONG == RATE_MASK_WU_TING_YONG {
		str = fmt.Sprintf("%s,%s", str, "无听用")
	}
	return str
}

type RateResult struct {
	Mask      int
	GangCount int
}

// 判断翻数
// 0x00000000 低8bit，存放基本番，高24bit存放特殊翻数标识
func JudgeRate(req *JudgeReq) (rate *RateResult) {
	if req.MaxRate == 0 {
		req.MaxRate = 6
	}
	lenOfHands := len(req.handsWithLz)
	rate = &RateResult{}
	if lenOfHands%3 != 2 {
		log.Printf("JudgeRate hands number error: %v \n", req.hands)
		return
	}
	judgeRateImpl(rate, req)
	judgeBaseRateWithRate(rate, req)
	//judgeSpecMask(rate, req.hands, req.Events)
	return
}

// 判断特殊牌型
func judgeSpecMask(rate *RateResult, hands []byte, events []MjEvent) {
	lenOfHands := len(hands)
	// 1. 判断门清
	if lenOfHands == 14 {
		rate.Mask |= RATE_MASK_MEN_QING
	}
	// 2. 判断中张, 没有1，9
	isZhognZhang := true
	for _, m := range hands {
		switch m {
		case 0x01:
		case 0x09:
		case 0x11:
		case 0x19:
		case 0x21:
		case 0x29:
			isZhognZhang = false
			break
		}
	}
	if isZhognZhang {
		rate.Mask |= RATE_MASK_ZHONG_ZHANG
	}
	// 3. 夹心5 TODO
	// 4. 幺九将对 TODO
}

// 判断基本番数
func judgeBaseRate(req *JudgeReq) (rate *RateResult) {
	rate = &RateResult{}
	judgeBaseRateWithRate(rate, req)
	return
}

func judgeBaseRateWithRate(rate *RateResult, req *JudgeReq) {
	hands := req.hands
	events := req.Events
	lenOfHands := len(hands)
	// 1. 判断4张，和手牌1张，碰1张情况
	if lenOfHands > 4 {
		for i := 0; i < lenOfHands-4; {
			if hands[i] == hands[i+1] && hands[i] == hands[i+2] && hands[i] == hands[i+3] && hands[i] == hands[i+4] {
				rate.GangCount++
				i += 4
			} else {
				if events != nil {
					for _, e := range events {
						if e.IsPeng() && Mj(e.Key).ToByte() == hands[i] {
							rate.GangCount++
						}
					}
				}
				i++
			}
		}
	}
	// 2. 判断杠
	if events != nil {
		for _, e := range events {
			if e.IsGang() {
				rate.GangCount++
			}
		}
	}
	// 3. 清一色
	if req.colorCount == 0x01 {
		rate.Mask |= RATE_MASK_QING_YI_SE
	}
}

func judgeRateImpl(rate *RateResult, req *JudgeReq) {
	lenOfHands := len(req.handsWithLz)
	if lenOfHands == 14 && judge7Dui(req.handsWithLz) {
		rate.Mask |= RATE_MASK_7DUI
	}
	// 金钩胡
	if lenOfHands == 2 {
		rate.Mask |= RATE_MASK_JIN_GOU
	}
	// 对对胡
	if judgeDuiDuiHu(req.handsWithLz) {
		rate.Mask |= RATE_MASK_DUIDUI_HU
	}
}

// 判断对对胡，23，233，2333 牌型
func judgeDuiDuiHu(hands []byte) bool {
	var dui bool = false
	for i := 0; i < len(hands); {
		if hands[i] == hands[i+1] && hands[i] == hands[i+2] {
			if hands[i] == hands[i+3] {
				i += 3
			} else if dui {
				return false
			} else {
				dui = true
				i += 2
			}
		} else {
			return false
		}
	}
	return true
}

// 七对
func judge7Dui(hands []byte) bool {
	for i := 0; i < 14; i += 2 {
		if hands[i] != hands[i+1] {
			return false
		}
	}
	return true
}
