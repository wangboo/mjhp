package main

import (
	"mjhp"
)

func main() {
	//mjhp.CheckDataFiles()
	mjhp.LoadConfig()
	mjhp.LoadData()
	mjhp.StartComputeWork()
	mjhp.StartKafka()
	//mjhp.TestJudgeHu([]int{
	//	0,0,1,1,2,2,3,3,4,5,6,6,7,8,
	//}, 0, 3)
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
