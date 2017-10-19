package mjhp

import (
	"log"
	"fmt"
)

const (
	RATE_MASK_7DUI            = 1 << 0  // 7对
	RATE_MASK_MEN_QING        = 1 << 1  // 门清
	RATE_MASK_ZHONG_ZHANG     = 1 << 2  // 中张
	RATE_MASK_JIA_XIN5        = 1 << 3  // 夹心5
	RATE_MASK_JIN_GOU         = 1 << 4  // 金钩胡
	RATE_MASK_DUIDUI_HU       = 1 << 5  // 对对胡
	RATE_MASK_QING_YI_SE      = 1 << 6  // 清一色
	RATE_MASK_BENJIN          = 1 << 7  // 本金暗刻
	RATE_MASK_LAIZI8          = 1 << 8  // 8个癞子
	RATE_MASK_LAIZI4          = 1 << 9  // 4个癞子
	RATE_MASK_WU_TING_YONG    = 1 << 10 // 无听用
	RATE_MASK_TIAN_HU         = 1 << 11 // 天胡
	RATE_MASK_DI_HU           = 1 << 12 // 地胡
	RATE_MASK_JIANG_DUI       = 1 << 13 // 将对
	RATE_MASK_HAI_DI_ZM       = 1 << 14 // 海底自摸
	RATE_MASK_HAI_DI_PAO      = 1 << 15 // 海底炮
	RATE_MASK_GSH             = 1 << 16 // 杠上花
	RATE_MASK_GSP             = 1 << 17 // 杠上炮
	RATE_MASK_GQH             = 1 << 18 // 抢杠胡
	RATE_MASK_LONG_7DUI       = 1 << 19 // 龙7对
	RATE_MASK_QING_7DUI       = 1 << 20 // 清七对
	RATE_MASK_LONG_7DUI2      = 1 << 21 // 双龙七对
	RATE_MASK_YAO_JIU         = 1 << 22 // 幺九
	RATE_MASK_DA_JIAO         = 1 << 23 // 查大觉
	RATE_MASK_HUA_ZU          = 1 << 24 // 查花猪
	RATE_MASK_QING_7DUI_LONG  = 1 << 25 // 清一色龙7对
	RATE_MASK_QING_7DUI_LONG2 = 1 << 26 // 清一色双龙7对
	RATE_MASK_QING_DUI        = 1 << 27 // 清对
	RATE_MASK_LONG_7DUI3      = 1 << 28 // 三龙七对
	RATE_MASK_QUEMEN          = 1 << 29 // 缺门加番
	RATE_MASK_PIN_HU          = 1 << 30 // 平胡
	RATE_MASK_DDZ             = 1 << 31 // 大对子
	RATE_MASK_BAI_PAI         = 1 << 32 // 摆牌
	RATE_MASK_BAI_DU_PAI      = 1 << 33 // 摆独牌
	RATE_MASK_YI_BAN_GAO      = 1 << 34 // 一般高

	RATE_MASK = int64(0xffffffffffffff)
)

func rateToString(v int64) string {
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
	if v&RATE_MASK_JIANG_DUI == RATE_MASK_JIANG_DUI {
		str = fmt.Sprintf("%s,%s", str, "将对")
	}
	if v&RATE_MASK_YAO_JIU == RATE_MASK_YAO_JIU {
		str = fmt.Sprintf("%s,%s", str, "幺九")
	}
	if v&RATE_MASK_HUA_ZU == RATE_MASK_HUA_ZU {
		str = fmt.Sprintf("%s,%s", str, "查花猪")
	}
	if v&RATE_MASK_DA_JIAO == RATE_MASK_DA_JIAO {
		str = fmt.Sprintf("%s,%s", str, "查大叫")
	}
	if v&RATE_MASK_QING_7DUI_LONG == RATE_MASK_QING_7DUI_LONG {
		str = fmt.Sprintf("%s,%s", str, "清龙七对")
	}
	if v&RATE_MASK_QING_7DUI_LONG2 == RATE_MASK_QING_7DUI_LONG2 {
		str = fmt.Sprintf("%s,%s", str, "清龙双七对")
	}
	if v&RATE_MASK_QING_DUI == RATE_MASK_QING_DUI {
		str = fmt.Sprintf("%s,%s", str, "清对")
	}
	if v&RATE_MASK_LONG_7DUI == RATE_MASK_LONG_7DUI {
		str = fmt.Sprintf("%s,%s", str, "龙七对")
	}
	if v&RATE_MASK_QING_7DUI == RATE_MASK_QING_7DUI {
		str = fmt.Sprintf("%s,%s", str, "清七对")
	}
	if v&RATE_MASK_LONG_7DUI2 == RATE_MASK_LONG_7DUI2 {
		str = fmt.Sprintf("%s,%s", str, "双龙七对")
	}
	if v&RATE_MASK_PIN_HU == RATE_MASK_PIN_HU {
		str = fmt.Sprintf("%s,%s", str, "平胡")
	}
	if v&RATE_MASK_DDZ == RATE_MASK_DDZ {
		str = fmt.Sprintf("%s,%s", str, "大对子")
	}
	if v&RATE_MASK_BAI_PAI == RATE_MASK_BAI_PAI {
		str = fmt.Sprintf("%s,%s", str, "摆牌")
	}
	if v&RATE_MASK_BAI_DU_PAI == RATE_MASK_BAI_DU_PAI {
		str = fmt.Sprintf("%s,%s", str, "摆独牌")
	}
	if v&RATE_MASK_YI_BAN_GAO == RATE_MASK_YI_BAN_GAO {
		str = fmt.Sprintf("%s,%s", str, "一般高")
	}
	return str
}

const (
	CMASK_DDH        = 0x1
	CMASK_YJJD       = 0x2
	CMASK_JX5        = 0x4
	CMASK_MENQING    = 0x8
	CMASK_ZHONGZHANG = 0x10
)

type RateResult struct {
	Mask      int64
	GangCount int
}

func (this *JudgeReq) IsDdh2Rate() bool {
	return this.CMask&CMASK_DDH == CMASK_DDH
}

func (this *JudgeReq) IsYJJD() bool {
	return this.CMask&CMASK_YJJD == CMASK_YJJD
}

func (this *JudgeReq) IsJX5() bool {
	return this.CMask&CMASK_JX5 == CMASK_JX5
}

func (this *JudgeReq) IsMenQing() bool {
	return this.CMask&CMASK_MENQING == CMASK_MENQING
}

func (this *JudgeReq) IsZhongZhang() bool {
	return this.CMask&CMASK_ZHONGZHANG == CMASK_ZHONGZHANG
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
}

// 判断基本番数
func judgeBaseRate(req *JudgeReq) (rate *RateResult) {
	rate = &RateResult{}
	judgeBaseRateWithRate(rate, req)
	return
}

// 基本番数，要计入癞子加番
// 牌型：1,1,1,2,2,2,2,3,3,3
// 1. 血战麻将，只能胡1，3万，番型 清一色 + 1根
// 2. 宜宾麻将，胡1，2，3万，胡2万时，牌型：清一色对对胡 + 1根
func judgeBaseRateWithRate(rate *RateResult, req *JudgeReq) {
	hands := req.handsWithLz
	lenOfHands := len(hands)
	// 1. 判断4张，和手牌1张，碰1张情况
  req.GenCount = 0
	if lenOfHands > 4 {
		for i := 0; i < lenOfHands-3; {
			if hands[i] == hands[i+1] && hands[i] == hands[i+2] && hands[i] == hands[i+3] {
				if !req.IsChaJiao() || !req.IsYiBinMj() || getMjCount(req.hands, hands[i]) > 2 {
					rate.GangCount++
					req.GenCount++
				}
				i += 4
			} else {
				i++
			}
		}
	}

	// 2. 判断杠
	events := req.Events
	if events != nil {
		for _, e := range events {
			if e.IsGang() {
				rate.GangCount++
			}
			if (e.IsPeng() || e.IsFei()) && hasMj(req.handsWithLz, mjIntToByte(e.Key)) {
				rate.GangCount++
				req.GenCount++
			}
		}
	}
	// 3. 清一色
	if req.colorCount == 0x01 {
		rate.Mask |= RATE_MASK_QING_YI_SE
	}
}

func hasMj(hands []byte, v byte) bool {
	for _, m := range hands {
		if m == v {
			return true
		}
	}
	return false
}

// 老版本，癞子不计入番数
//func judgeBaseRateWithRate(rate *RateResult, req *JudgeReq) {
//	// 基本番数，不计入癞子
//	hands := req.hands
//	if req.IsChaJiao() {
//		hands = req.handsWithLz
//	}
//	events := req.Events
//	lenOfHands := len(hands)
//	// 1. 判断4张，和手牌1张，碰1张情况
//	if lenOfHands > 4 {
//		for i := 0; i < lenOfHands-3; {
//			if hands[i] == hands[i+1] && hands[i] == hands[i+2] && hands[i] == hands[i+3] {
//				if !req.IsChaJiao() || getMjCount(req.hands, hands[i]) > 2 {
//					rate.GangCount++
//				}
//				i += 4
//			} else {
//				i++
//			}
//		}
//	}
//	add1 := 0
//	if req.IsChaJiao() {
//		// 查叫，先查基本加番
//		for i := 0; i < len(req.hands); i++ {
//			if events != nil {
//				for _, e := range events {
//					if e.IsPeng() && Mj(e.Key).ToByte() == req.hands[i] {
//						add1++
//					}
//				}
//			}
//		}
//	}
//	add2 := 0
//	// 再查癞子增加番
//	if events != nil {
//		for i := 0; i < lenOfHands; i++ {
//			for _, e := range events {
//				if e.IsPeng() && Mj(e.Key).ToByte() == hands[i] {
//					add2++
//				}
//			}
//		}
//	}
//	//printMj(req.handsWithLz)
//	//log.Println("add1 = ", add1, ", add2 = ", add2)
//	if req.IsChaJiao() && add2-add1 > 1 {
//		rate.GangCount += add1 + 1
//	} else {
//		rate.GangCount += add1 + add2
//	}
//	// 2. 判断杠
//	if events != nil {
//		for _, e := range events {
//			if e.IsGang() {
//				rate.GangCount++
//			}
//		}
//	}
//	// 3. 清一色
//	if req.colorCount == 0x01 {
//		rate.Mask |= RATE_MASK_QING_YI_SE
//	}
//}

func getMjCount(hands []byte, v byte) int {
	count := 0
	for _, h := range hands {
		if h == v {
			count++
		} else if h > v {
			return count
		}
	}
	return count
}

func judgeRateImpl(rate *RateResult, req *JudgeReq) {
	lenOfHands := len(req.handsWithLz)
	if lenOfHands == 14 && judge7Dui(req.handsWithLz) {
		//log.Println("7dui")
		rate.Mask |= RATE_MASK_7DUI
	}
	// 金钩胡
	if lenOfHands == 2 {
		rate.Mask |= RATE_MASK_JIN_GOU
	}
	// 对对胡
	if judgeDuiDuiHu(req.handsWithLz) {
		//log.Println("duiduiHu")
		rate.Mask |= RATE_MASK_DUIDUI_HU
	}
}

// 判断对对胡，23，233，2333 牌型
func judgeDuiDuiHu(hands []byte) bool {
	lenOfHands := len(hands)
	var dui bool = false
	for i := 0; i < lenOfHands-1; {
		if hands[i] == hands[i+1] {
			if i+2 < lenOfHands && hands[i] == hands[i+2] {
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
