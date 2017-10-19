package main

import (
	"mjhp"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

const TEST_MODE = false

func main() {
	mjhp.CheckDataFiles()
	mjhp.LoadConfig()
	mjhp.LoadData()
	if !TEST_MODE {
		mjhp.StartComputeWork()
		mjhp.StartKafka()
	} else {
		bin, _ := ioutil.ReadFile("test.json")
		req := &mjhp.JudgeReqBatch{}
		err := json.Unmarshal(bin, req)
		if err != nil {
			panic(err)
		}
		mjhp.JudgeHuBatch(req, false)
	}

	//str := "1t,1t,1t,3w,3w,5t,5t,6t,6t,7t,7t"
	//str := "2万, 2万, 5万, 5万, 7万, 7万, 2筒, 2筒, 2筒, 2筒, 4筒, 4筒, 4筒, 4筒"
	//str := "3万, 3万, 3万, 9万, 9万, 9万, 2筒, 2筒, 2筒, 4筒, 4筒, 7筒, 7筒"
	// Mask 第0位 	是否略过翻数判断 （翻数判断，必定缺一门）1
	// Mask 第1位	宜宾麻将 2
	// Mask 第2位	血战麻将	4
	// Mask 第3位	是否为查叫（查叫会添加一个癞子，算4癞子，8癞子会减1）8
	//req := &mjhp.JudgeReq{Hands: parse(str), LzCount: 3, LzTotal: 3, JudgeMj: -1, Mask: 0x0a,
	//CMask: mjhp.CMASK_YJJD | mjhp.CMASK_MENQING | mjhp.CMASK_ZHONGZHANG,
	//CMask: 0,
	//DingQue: 2, MaxRate: 9,
	//}
	//mjhp.TestJudgeHu(req)

	//mjhp.FactorCreateYaoJiu()
	//mjhp.FactorCreateLz5_6()
	//mjhp.FactorCreateLz5_9()
	//mjhp.FactorCreateLz6_8()
	//mjhp.FactorCreateLz6_8()
	//mjhp.FactorSimpleToSet()
	//mjhp.FactorCreateSet14()
	//mjhp.FactorCreateSet11()
	//mjhp.FactorCreateSet8()
	//mjhp.FactorCreateSet5()

	//mjhp.LoadMap()
	//mjhp.BenchmarkWinEx2(1000)
	//mjhp.MakeTest()
	//time.Sleep(100 * time.Second)
	//benchmarkWin(10000000, wins)
	//benchmarkWinEx(1000)
	//benchmarkWin3(10000000, wins)
}

func parse(str string) []int {
	ss := strings.Split(str, ", ")
	if len(ss) < 2 {
		ss = strings.Split(str, "，")
	}
	if len(ss) < 2 {
		ss = strings.Split(str, ",")
	}
	arr := make([]int, len(ss), len(ss))
	for idx, v := range ss {
		sa := strings.Split(v, "")
		switch sa[1] {
		case "万":
			append(arr, idx, 0, sa)
		case "w":
			append(arr, idx, 0, sa)
		case "筒":
			append(arr, idx, 9, sa)
		case "t":
			append(arr, idx, 9, sa)
		case "条":
			append(arr, idx, 18, sa)
		case "a":
			append(arr, idx, 18, sa)
		}
	}
	return arr
}

func append(arr []int, idx, n int, sa []string) {
	x, err := strconv.Atoi(sa[0])
	if err != nil {
		panic(err)
	}
	arr[idx] = n + x - 1
}
