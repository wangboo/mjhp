package mjhp

import (
	"os"
	"bufio"
	"strconv"
	"log"
	"sort"
)

func CheckDataFiles() {
	checkOrCreate("data/set14.data", FactorCreateSet14)
	checkOrCreate("data/set11.data", FactorCreateSet11)
	checkOrCreate("data/set8.data", FactorCreateSet8)
	checkOrCreate("data/set5.data", FactorCreateSet5)
	checkOrCreate("data/lz6_8.data", FactorCreateLz6_8)
	checkOrCreate("data/lz5_9.data", FactorCreateLz5_9)
	checkOrCreate("data/lz5_6.data", FactorCreateLz5_6)
	checkOrCreate("data/yj.data", FactorCreateYaoJiu)
	checkOrCreate("data/yibangao.data", FactorCreateYiBanGao)
}

func checkOrCreate(file string, createFunc func()) {
	_, err := os.Stat(file)
	if err != nil {
		createFunc()
	} else {
		log.Printf("data file: %s exists.\n", file)
	}
}

func FactorCreateSet14() {
	log.Println("making set14.data")
	f, err := os.Create("data/set14.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cache := make(map[int]bool, 8800)
	buf := bufio.NewWriterSize(f, 2048)
	pair := findPairs()
	group := findGroups()
	hands := make([]byte, 14, 14)
	for _, p := range pair {
		for _, g0 := range group {
			for _, g1 := range group {
				for _, g2 := range group {
					for _, g3 := range group {
						hands[0] = p[0]
						hands[1] = p[1]
						hands[2] = g0[0]
						hands[3] = g0[1]
						hands[4] = g0[2]
						hands[5] = g1[0]
						hands[6] = g1[1]
						hands[7] = g1[2]
						hands[8] = g2[0]
						hands[9] = g2[1]
						hands[10] = g2[2]
						hands[11] = g3[0]
						hands[12] = g3[1]
						hands[13] = g3[2]
						if checkValid(hands) {
							v := bytesToInt(hands)
							cache[v] = true
						}
					}
				}
			}
		}
	}
	for k := range cache {
		buf.WriteString(strconv.Itoa(k))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

func FactorCreateSet11() {
	log.Println("making set11.data")
	f, err := os.Create("data/set11.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cache := make(map[int]bool, 2000)
	buf := bufio.NewWriterSize(f, 2048)
	pair := findPairs()
	group := findGroups()
	hands := make([]byte, 11, 11)
	for _, p := range pair {
		for _, g0 := range group {
			for _, g1 := range group {
				for _, g2 := range group {
					hands[0] = p[0]
					hands[1] = p[1]
					hands[2] = g0[0]
					hands[3] = g0[1]
					hands[4] = g0[2]
					hands[5] = g1[0]
					hands[6] = g1[1]
					hands[7] = g1[2]
					hands[8] = g2[0]
					hands[9] = g2[1]
					hands[10] = g2[2]
					if checkValid(hands) {
						v := bytesToInt(hands)
						cache[v] = true
					}
				}
			}
		}
	}
	for k := range cache {
		buf.WriteString(strconv.Itoa(k))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

func FactorCreateSet8() {
	log.Println("making set8.data")
	f, err := os.Create("data/set8.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cache := make(map[int]bool, 300)
	buf := bufio.NewWriterSize(f, 2048)
	pair := findPairs()
	group := findGroups()
	hands := make([]byte, 8, 8)
	for _, p := range pair {
		for _, g0 := range group {
			for _, g1 := range group {
				hands[0] = p[0]
				hands[1] = p[1]
				hands[2] = g0[0]
				hands[3] = g0[1]
				hands[4] = g0[2]
				hands[5] = g1[0]
				hands[6] = g1[1]
				hands[7] = g1[2]
				if checkValid(hands) {
					v := bytesToInt(hands)
					cache[v] = true
				}
			}
		}
	}
	for k := range cache {
		buf.WriteString(strconv.Itoa(k))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

func FactorCreateSet5() {
	log.Println("making set5.data")
	f, err := os.Create("data/set5.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cache := make(map[int]bool, 100)
	buf := bufio.NewWriterSize(f, 2048)
	pair := findPairs()
	group := findGroups()
	hands := make([]byte, 5, 5)
	for _, p := range pair {
		for _, g0 := range group {
			hands[0] = p[0]
			hands[1] = p[1]
			hands[2] = g0[0]
			hands[3] = g0[1]
			hands[4] = g0[2]
			if checkValid(hands) {
				v := bytesToInt(hands)
				cache[v] = true
			}
		}
	}
	for k := range cache {
		buf.WriteString(strconv.Itoa(k))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

var tiles2 = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
}

var tiles6_8_1 = [][]byte{
	{0x01, 0x02, 0x03}, {0x04, 0x05, 0x06}, {0x07, 0x08, 0x09},
	{0x02, 0x03, 0x04}, {0x05, 0x06, 0x07},
	{0x03, 0x04, 0x05}, {0x06, 0x07, 0x08},
}

var tiles6_8_2 = [][]byte{
	{0x11, 0x12, 0x13}, {0x14, 0x15, 0x16}, {0x17, 0x18, 0x19},
	{0x12, 0x13, 0x14}, {0x15, 0x16, 0x17}, {0x16, 0x17, 0x18},
	{0x13, 0x14, 0x15},
}

func findPairs2() [][]byte {
	pairs := make([][]byte, 0, 18)
	for _, v := range tiles2 {
		pairs = append(pairs, []byte{v, v})
	}
	return pairs
}

func findGroups2Only3() [][]byte {
	groups := make([][]byte, 0, 18)
	for _, v := range tiles2 {
		groups = append(groups, []byte{v, v, v})
	}
	return groups
}

func findGroups2WithOut3() [][]byte {
	groups := make([][]byte, 0, 14)
	for i := 0; i < 16; i++ {
		if tiles2[i] == tiles2[i+1]-1 && tiles2[i] == tiles2[i+2]-2 {
			groups = append(groups, []byte{tiles2[i], tiles2[i+1], tiles2[i+2]})
		}
	}
	return groups
}

// 1008 种胡牌组合，每种暴力拆分
func FactorCreateLz6_8() {
	pairs := findPairs2()
	log.Printf("pairs = %v \n", pairs)
	hands := make([]byte, 14, 14)
	droped := make([]byte, 8, 8)
	cache := make(map[int64]bool, 100)
	// 3 - 1 搭配,
	lenOfTiles := len(tiles6_8_1)
	count := 0
	for _, p := range pairs {
		for i := 0; i < lenOfTiles; i++ {
			g0 := tiles6_8_1[i]
			for j := i + 1; j < lenOfTiles; j++ {
				g1 := tiles6_8_1[j]
				if &g0 == &g1 {
					continue
				}
				for k := j + 1; k < lenOfTiles; k++ {
					g2 := tiles6_8_1[k]
					if &g1 == &g2 || &g0 == &g2 {
						continue
					}
					for l := 0; l < lenOfTiles; l++ {
						g3 := tiles6_8_2[l]
						count++
						breakTheHands6(cache, hands, droped, p, g0, g1, g2, g3)
						if count%100 == 0 {
							log.Printf("count = %d, cache.size = %d \n", count, len(cache))
						}
					}
				}
			}
		}
	}
	// 2 - 2 配
	for _, p := range pairs {
		for i := 0; i < lenOfTiles; i++ {
			g0 := tiles6_8_1[i]
			for j := i + 1; j < lenOfTiles; j++ {
				g1 := tiles6_8_1[j]
				if &g0 == &g1 {
					continue
				}
				for k := j + 1; k < lenOfTiles; k++ {
					g2 := tiles6_8_2[k]
					for l := 0; l < lenOfTiles; l++ {
						g3 := tiles6_8_2[l]
						if &g2 == &g3 {
							continue
						}
						count++
						breakTheHands6(cache, hands, droped, p, g0, g1, g2, g3)
						if count%100 == 0 {
							log.Printf("count = %d, cache.size = %d \n", count, len(cache))
						}
					}
				}
			}
		}
	}
	// 1 - 3 配
	for _, p := range pairs {
		for i := 0; i < lenOfTiles; i++ {
			g0 := tiles6_8_1[i]
			for j := i + 1; j < lenOfTiles; j++ {
				g1 := tiles6_8_2[j]
				for k := j + 1; k < lenOfTiles; k++ {
					g2 := tiles6_8_2[k]
					if &g1 == &g2 {
						continue
					}
					for l := 0; l < lenOfTiles; l++ {
						g3 := tiles6_8_2[l]
						if &g2 == &g3 || &g1 == &g3 {
							continue
						}
						count++
						breakTheHands6(cache, hands, droped, p, g0, g1, g2, g3)
						if count%100 == 0 {
							log.Printf("count = %d, cache.size = %d \n", count, len(cache))
						}
					}
				}
			}
		}
	}
	log.Println("count = ", count)
	log.Println("making lz6_8.data")
	f, err := os.Create("data/lz6_8.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriterSize(f, 2048)
	for k := range cache {
		buf.WriteString(strconv.FormatInt(k, 10))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

//func fillHands(hands []byte, p []byte, g0 []byte, g1 []byte, g2 []byte, g3 []byte) {
func fillHands(hands []byte, p []byte, g0 ...[]byte) {
	hands[0] = p[0]
	hands[1] = p[1]
	for i := 0; i < len(g0); i++ {
		hands[2+i*3] = g0[i][0]
		hands[3+i*3] = g0[i][1]
		hands[4+i*3] = g0[i][2]
	}
}

func copyNotEmpty(src []byte, dst []byte) {
	idx := 0
	for _, v := range src {
		if v != 0x00 {
			dst[idx] = v
			idx++
		}
	}
}

func pairRepeated(g0 []byte, g1 []byte) bool {
	return g0[0] == g1[0] || g0[0] == g1[1] || g0[0] == g1[2] ||
		g0[1] == g1[0] || g0[1] == g1[1] || g0[1] == g1[2]
}

func groupRepeated(g0 []byte, g1 []byte) bool {
	return g0[0] == g1[0] || g0[0] == g1[1] || g0[0] == g1[2] ||
		g0[1] == g1[0] || g0[1] == g1[1] || g0[1] == g1[2] ||
		g0[2] == g1[0] || g0[2] == g1[1] || g0[2] == g1[2]
}

// 5个癞子，9张手牌 c2*2+c3*3 < 5
// 所以不存在2对，和2个三张以上
// 所以gt3 最多一次，gt的牌不可能重合
func FactorCreateLz5_9() {
	pairs := findPairs2()
	gt := findGroups2WithOut3()
	gt3 := findGroups2Only3()
	lenOfGt := len(gt)
	hands := make([]byte, 14, 14)
	droped := make([]byte, 9, 9)
	cache := make(map[int64]bool, 100)
	count := 0
	log.Printf("pairs: %v \n gt: %v \n, gt3: %v ", pairs, gt, gt3)
	for pi, p := range pairs {
		for i0, g0 := range gt3 {
			for i1 := 0; i1 < lenOfGt-2; i1++ {
				g1 := gt[i1]
				if groupRepeated(g0, g1) {
					continue
				}
				for i2 := 1; i2 < lenOfGt-1; i2++ {
					g2 := gt[i2]
					if groupRepeated(g0, g2) || &g1 == &g2 { // g1 g2 地址相等，则完全相等
						continue
					}
					for i3 := 2; i3 < lenOfGt; i3++ {
						g3 := gt[i3]
						if groupRepeated(g0, g3) || &g1 == &g3 || &g2 == &g3 {
							continue
						}
						// 暴力拆解
						breakTheHands5(cache, hands, droped, p, g0, g1, g2, g3)
						count++
						if count%10000 == 0 {
							log.Printf("count1 = %d, cache = %d, pi=%d, i0=%d\n",
								count, len(cache), pi, i0)
							printMj(droped)
						}
					}
				}
			}
		}
	}
	for pi, p := range pairs {
		for i0 := 0; i0 < lenOfGt-3; i0++ {
			g0 := gt[i0]
			for i1 := 0; i1 < lenOfGt-2; i1++ {
				g1 := gt[i1]
				if &g0 == &g1 {
					continue
				}
				for i2 := 1; i2 < lenOfGt-1; i2++ {
					g2 := gt[i2]
					if &g0 == &g2 || &g1 == &g2 { // g1 g2 地址相等，则完全相等
						continue
					}
					for i3 := 2; i3 < lenOfGt; i3++ {
						g3 := gt[i3]
						if &g0 == &g3 || &g1 == &g3 || &g2 == &g3 {
							continue
						}
						// 暴力拆解
						breakTheHands5(cache, hands, droped, p, g0, g1, g2, g3)
						count++
						if count%10000 == 0 {
							log.Printf("count2 = %d, cache = %d, pi=%d, i0=%d\n",
								count, len(cache), pi, i0)
							printMj(droped)
						}
					}
				}
			}
		}
	}
	log.Printf("count = %d, cache = %d \n", count, len(cache))
	log.Println("making lz5_9.data")
	f, err := os.Create("data/lz5_9.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriterSize(f, 2048)
	for k := range cache {
		buf.WriteString(strconv.FormatInt(k, 10))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

// 5个癞子，6张手牌, 牌型中只有一个对子,没有3张，否则可以组成大对子
func FactorCreateLz5_6() {
	pairs := findPairs2()
	log.Printf("pairs = %v \n", pairs)
	hands := make([]byte, 11, 11)
	droped := make([]byte, 6, 6)
	cache := make(map[int64]bool, 100)
	// 3 - 1 搭配,
	lenOfTiles := len(tiles6_8_1)
	count := 0
	for _, p := range pairs {
		for i := 0; i < lenOfTiles; i++ {
			g0 := tiles6_8_1[i]
			//log.Printf("i = %d, g0 = %v \n", i, g0)
			if pairRepeated(p, g0) {
				continue
			}
			for j := i + 1; j < lenOfTiles; j++ {
				g1 := tiles6_8_1[j]
				if &g0 == &g1 {
					continue
				}
				for k := j + 1; k < lenOfTiles; k++ {
					g2 := tiles6_8_2[k]
					count++
					breakTheHands5(cache, hands, droped, p, g0, g1, g2)
					if count%100 == 0 {
						log.Printf("count = %d, cache.size = %d \n", count, len(cache))
					}
				}
			}
		}
	}
	for _, p := range pairs {
		for i := 0; i < lenOfTiles; i++ {
			g0 := tiles6_8_1[i]
			//log.Printf("i = %d, g0 = %v \n", i, g0)
			if pairRepeated(p, g0) {
				continue
			}
			for j := i + 1; j < lenOfTiles; j++ {
				g1 := tiles6_8_2[j]
				for k := j + 1; k < lenOfTiles; k++ {
					g2 := tiles6_8_2[k]
					if &g1 == &g2 {
						continue
					}
					count++
					breakTheHands5(cache, hands, droped, p, g0, g1, g2)
					if count%100 == 0 {
						log.Printf("count = %d, cache.size = %d \n", count, len(cache))
					}
				}
			}
		}
	}
	log.Printf("count = %d, cache = %d \n", count, len(cache))
	log.Println("making lz5_6.data")
	f, err := os.Create("data/lz5_6.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriterSize(f, 2048)
	for k := range cache {
		buf.WriteString(strconv.FormatInt(k, 10))
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

func breakTheHands6(cache map[int64]bool, hands []byte, droped []byte, p []byte,
	g0 []byte, g1 []byte, g2 []byte, g3 []byte) {
	// 暴力拆解,6个癞子，
	for x0 := 0; x0 < 9; x0++ {
		for x1 := x0 + 1; x1 < 10; x1++ {
			for x2 := x1 + 1; x2 < 11; x2++ {
				for x3 := x2 + 1; x3 < 12; x3++ {
					for x4 := x3 + 1; x4 < 13; x4++ {
						for x5 := x4 + 1; x5 < 14; x5++ {
							fillHands(hands, p, g0, g1, g2, g3)
							hands[x0] = 0x00
							hands[x1] = 0x00
							hands[x2] = 0x00
							hands[x3] = 0x00
							hands[x4] = 0x00
							hands[x5] = 0x00
							copyNotEmpty(hands, droped)
							sort.Sort(byteSlice(droped))
							//printMj(droped)
							v := byteToInt64(droped)
							cache[v] = true
						}
					}
				}
			}
		}
	}
}

// , g1 []byte, g2 []byte, g3 []byte
func breakTheHands5(cache map[int64]bool, hands []byte, droped []byte, p []byte, g0 ...[]byte) {
	lenOfHands := len(hands)
	// 暴力拆解,5个癞子，
	for x0 := 0; x0 < lenOfHands-4; x0++ {
		for x1 := x0 + 1; x1 < lenOfHands-3; x1++ {
			for x2 := x1 + 1; x2 < lenOfHands-2; x2++ {
				for x3 := x2 + 1; x3 < lenOfHands-1; x3++ {
					for x4 := x3 + 1; x4 < lenOfHands; x4++ {
						fillHands(hands, p, g0...) // , g1, g2, g3
						hands[x0] = 0x00
						hands[x1] = 0x00
						hands[x2] = 0x00
						hands[x3] = 0x00
						hands[x4] = 0x00
						copyNotEmpty(hands, droped)
						sort.Sort(byteSlice(droped))
						//printMj(droped)
						v := byteToInt64(droped)
						cache[v] = true
					}
				}
			}
		}
	}
}

func FactorCreateYaoJiu() {
	pairs := [][]byte{
		{0x1, 0x1}, {0x9, 0x9}, {0x11, 0x11}, {0x19, 0x19},
	}
	groups := [][]byte{
		{0x1, 0x1, 0x1}, {0x1, 0x2, 0x3}, {0x9, 0x9, 0x9}, {0x7, 0x8, 0x9},
		{0x11, 0x11, 0x11}, {0x11, 0x12, 0x13}, {0x19, 0x19, 0x19}, {0x17, 0x18, 0x19},
	}
	cache := make(map[string]bool, 2000)
	hands5 := make([]byte, 5, 5)
	hands8 := make([]byte, 8, 8)
	hands11 := make([]byte, 11, 11)
	hands14 := make([]byte, 14, 14)
	for _, p := range pairs {
		for _, g0 := range groups {
			fillHands(hands5, p, g0)
			if checkValid(hands5) {
				v5 := bytesToStr(hands5)
				cache[v5] = true
			}
			for _, g1 := range groups {
				fillHands(hands8, p, g0, g1)
				if checkValid(hands8) {
					v8 := bytesToStr(hands8)
					cache[v8] = true
				}
				for _, g2 := range groups {
					fillHands(hands11, p, g0, g1, g2)
					if checkValid(hands11) {
						v11 := bytesToStr(hands11)
						cache[v11] = true
					}
					for _, g3 := range groups {
						fillHands(hands14, p, g0, g1, g2, g3)
						if checkValid(hands14) {
							v14 := bytesToStr(hands14)
							cache[v14] = true
						}
					}
				}
			}
		}
	}
	log.Println("making yj.data")
	f, err := os.Create("data/yj.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriterSize(f, 2048)
	for k := range cache {
		buf.WriteString(k)
		buf.WriteString("\r\n")
	}
	buf.Flush()
}

func FactorCreateYiBanGao() {

}
