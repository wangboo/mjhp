package mjhp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
	"strconv"
	"bufio"
	"io/ioutil"
	"encoding/base64"
)

var tiles = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
}

var full = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	//0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91,             // East South West North Red Green White
}

// 不带癞子
var wins = make(map[int]bool, 9000)
// 幺九2色
var lys = make(map[string]bool, 780)
// 6癞子8手牌
var lz6_8 = make(map[int64]bool, 410861)
// 5癞子9手牌
var lz5_9 = make(map[int64]bool, 2124056)
// 5癞子4手牌
var lz5_6 = make(map[int64]bool, 47132)

func findPairs() [][]byte {
	pairs := make([][]byte, 0, len(tiles))
	for _, v := range tiles {
		pair := []byte{v, v}
		pairs = append(pairs, pair)
	}
	return pairs
}

func findGroups() [][]byte {
	groups := make([][]byte, 0, len(tiles)+(9-2)*3)

	// find three identical tiles
	for _, v := range tiles {
		group := []byte{v, v, v}
		groups = append(groups, group)
	}

	// find three sequence tiles
	for i := 2; i < len(tiles); i++ {
		if tiles[i-2]+1 == tiles[i-1] && tiles[i-1] == tiles[i]-1 {
			group := []byte{tiles[i-2], tiles[i-1], tiles[i]}
			groups = append(groups, group)
		}
	}

	return groups
}

type byteSlice []byte

func (b byteSlice) Len() int {
	return len(b)
}

func (b byteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func checkValid(win []byte) bool {
	lenOfWin := len(win)
	if lenOfWin < 4 {
		return true
	}
	sort.Sort(byteSlice(win))
	for i := 4; i < lenOfWin; i++ {
		if win[i] == win[i-4] {
			return false
		}
	}
	return true
}

func printMj(hand []byte) {
	str := ""
	for _, b := range hand {
		if b < 0x10 {
			str = fmt.Sprintf("%s, %d万", str, b)
		} else if b < 0x20 {
			str = fmt.Sprintf("%s, %d筒", str, b&0x0f)
		} else if b < 0x30 {
			str = fmt.Sprintf("%s, %d条", str, b&0x0f)
		} else {
			switch b {
			case 0x31:
				str = fmt.Sprintf("%s, 东风,", str)
			case 0x41:
				str = fmt.Sprintf("%s, 西风,", str)
			case 0x51:
				str = fmt.Sprintf("%s, 南风,", str)
			case 0x61:
				str = fmt.Sprintf("%s, 北风,", str)
			case 0x71:
				str = fmt.Sprintf("%s, 红中,", str)
			case 0x81:
				str = fmt.Sprintf("%s, 绿板,", str)
			case 0x91:
				str = fmt.Sprintf("%s, 白板,", str)
			}
		}
	}
	log.Println(str)
}

type simpleData struct {
	K int    // Key
	V []byte // Value
}

// 癞子将牌型生成特征码，每个麻将不超过0x30,所以取6位
func byteToInt64(win []byte) int64 {
	lenOfWin := len(win)
	if lenOfWin > 10 {
		log.Printf("byteToInt64 not accept %v \n", win)
		return -1
	}
	var mask int64
	for idx, v := range win {
		shift := int64(v)
		shift <<= uint(idx * 6)
		mask |= shift
	}
	return mask
}

// 5张癞子，9张手牌是否可胡，已经确定hands是两色牌
func isCanHuLz5_9(hands []byte) bool {
	return isCanHuWithLz(lz5_9, hands)
}

// 5张癞子，9张手牌是否可胡，已经确定hands是两色牌
func isCanHuLz5_6(hands []byte) bool {
	return isCanHuWithLz(lz5_6, hands)
}

// 6张癞子，8张手牌是否可胡，已经确定hands是两色牌
func isCanHuLz6_8(hands []byte) bool {
	return isCanHuWithLz(lz6_8, hands)
}

func isCanHuWithLz(cache map[int64]bool, hands []byte) bool {
	arr := transFrom2Color(hands)
	v := byteToInt64(arr)
	_, ok := cache[v]
	if cfg.IsDebugOn {
		log.Println("isCanHuWithLz v = ", v)
	}
	return ok
}

func transFrom2Color(hands []byte) []byte {
	count := make([]byte, 3, 3)
	for _, v := range hands {
		if v < 0x10 {
			count[0]++
		} else if v < 0x20 {
			count[1]++
		} else {
			count[2]++
		}
	}
	//printMj(hands)
	var arr []byte
	if count[0] > 0 && count[1] > 0 {
		arr = hands
	} else if count[0] > 0 && count[2] > 0 { // 万，条 -> 万，筒
		arr = make([]byte, len(hands), len(hands))
		copy(arr, hands)
		for idx, v := range arr {
			if v > 0x20 {
				arr[idx] -= 0x10
			}
		}
	} else if count[1] > 0 && count[2] > 0 { // 筒，条 -> 万，筒
		arr = make([]byte, len(hands), len(hands))
		copy(arr, hands)
		for idx, v := range arr {
			if v > 0x20 {
				arr[idx] -= 0x20
			}
		}
		// 条变成了万，重新整理顺序
		sort.Sort(byteSlice(arr))
	}
	return arr
}

func bytesToInt(win []byte) int {
	tmp := make([]byte, 0, len(win))
	tmp = append(tmp, 1)
	for i, pos := 1, 0; i < len(win); i++ {
		if win[i-1] == win[i] {
			tmp[pos]++
		} else if win[i-1]+1 == win[i] {
			tmp = append(tmp, 1)
			pos++
		} else {
			tmp = append(tmp, 1)
			tmp[pos] += 0x0A
			pos++
		}
	}
	res := 1
	for _, v := range tmp {
		switch v {
		case 0x01:
			res <<= 1 // 0
		case 0x02:
			res <<= 3
			res |= 0x06 // 110
		case 0x03:
			res <<= 5
			res |= 0x1E // 11110
		case 0x04:
			res <<= 7
			res |= 0x7E // 1111110
		case 0x0B:
			res <<= 2
			res |= 0x02 // 10
		case 0x0C:
			res <<= 4
			res |= 0x0E // 1110
		case 0x0D:
			res <<= 6
			res |= 0x3E // 111110
		case 0x0E:
			res <<= 8
			res |= 0xFE // 11111110
		}
	}
	//log.Printf("res = %b\n", res)
	return res
}

func bytesToStr(bin []byte) string {
	return base64.StdEncoding.EncodeToString(bin)
}

func printMask(mask []byte) {
	var str string
	for _, m := range mask {
		str = fmt.Sprintf("%s0x%02x,", str, m)
	}
	log.Println("mask = ", str)
}

func isCanHu(hands []byte) bool {
	sort.Sort(byteSlice(hands))
	if len(hands) == 14 && judge7Dui(hands) {
		return true
	}
	//printMj(hands)
	v := bytesToInt(hands)
	//log.Println("v: ", v)
	_, ok := wins[v]
	return ok
}

func LoadData() {
	if len(wins) > 0 {
		log.Println("wins data is already laoded")
		return
	}
	if _, err := os.Stat("data"); err != nil {
		os.Mkdir("data", os.ModePerm)
	}
	loadDataFile(wins, "data/set14.data")
	loadDataFile(wins, "data/set11.data")
	loadDataFile(wins, "data/set8.data")
	loadDataFile(wins, "data/set5.data")
	loadDataStrFile(lys, "data/yj.data")
	loadDataFile64(lz6_8, "data/lz6_8.data")
	loadDataFile64(lz5_9, "data/lz5_9.data")
	loadDataFile64(lz5_6, "data/lz5_6.data")
}

func loadDataFile(cache map[int]bool, file string) {
	beginAt := time.Now()
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("Open", err)
	}
	// 一次性载入内存
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(bytes.NewReader(data))
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		str := string(line)
		if str == "" {
			continue
		}
		v, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(fmt.Sprintf("str(%s) atoi failed", str), err)
		}
		cache[v] = true
	}
	log.Printf("loaded %s, cost: %v, cache size: %d\n", file, time.Since(beginAt), len(cache))
}

func loadDataFile64(cache map[int64]bool, file string) {
	beginAt := time.Now()
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("Open", err)
	}
	// 一次性载入内存
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(bytes.NewReader(data))
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		str := string(line)
		if str == "" {
			continue
		}
		v, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			log.Fatal(fmt.Sprintf("str(%s) atoi failed", str), err)
		}
		cache[v] = true
	}
	log.Printf("loaded %s, cost: %v, cache size: %d\n", file, time.Since(beginAt), len(cache))
}

func loadDataStrFile(cache map[string]bool, file string) {
	beginAt := time.Now()
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("Open", err)
	}
	// 一次性载入内存
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(bytes.NewReader(data))
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		str := string(line)
		if str == "" {
			continue
		}
		cache[str] = true
	}
	log.Printf("loaded %s, cost: %v, cache size: %d\n", file, time.Since(beginAt), len(cache))
}
