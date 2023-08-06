package main

import (
	"fmt"
	"golang-utils/age-utils"
	"golang-utils/channelutils"
	"golang-utils/gadget"
	"golang-utils/name_utils"
	"golang-utils/timeutils"
	"math"
	"strconv"
)

// 验证包名是否可以带中横线和下划线
func verifyGolangPackageName() {
	fmt.Println("Current Year is : " + strconv.Itoa(timeutils.GetYear()))
	fmt.Println("NickName is : " + name_utils.GetNickName())
	fmt.Println("Age is : " + strconv.Itoa(age_utils.GetAge()))
}

// 遍历数组
func iterateArray() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}
	for i, n := range notes {
		fmt.Printf("%d: %s\n", i, n)
	}
}

// 计算数组的平均值
func calculateAverage() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)

}

type Player interface {
	Play(song string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	// verifyGolangPackageName()
	// iterateArray()
	// calculateAverage()

	// fmt.Println(maximumNumber(71.8, 56.2, 89.5))
	// fmt.Println(maximumNumber(90.7, 89.7, 98.5, 92.3))

	// mixedTape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	// var player Player = gadget.TapePlayer{}
	// playList(player, mixedTape)
	// player = gadget.TapeRecorder{}
	// playList(player, mixedTape)

	// TryOut(gadget.TapeRecorder{})

	// fmt.Println("逗号在\"码神来了,码神bye bye\"中的位置是: ", stringutils.SearchString("码神来了,码神bye bye", ","))

	// stringNumber := "5566"
	// pureNumber, _ := stringutils.Switch2Int(stringNumber)
	// fmt.Println("字符串\"", stringNumber, "\"转换为整型数字: ", pureNumber)

	// num := 7788
	// numString := numberutils.Int2String(num)
	// fmt.Printf("数字%d转换为字符串: %s\n", num, numString)

	// arr := [...]int{1, 2, 3, 7: 113}
	// slice := arr[2:2]
	// fmt.Println("测试切片起始位置和结束位置相同是否会编译报错: ", slice)

	// var strList []string
	// strList = append(strList, "码神之路")
	// fmt.Println(strList)

	// a := make([]int, 2)
	// b := make([]int, 2, 10)
	// fmt.Println(a, b)

	// a, b, _ := namedReturnValueFunction()
	// fmt.Println(a, b)

	// var whatever = [5]int{1, 2, 3, 4, 5}
	// for _, v := range whatever {
	// 	// v := v
	// 	defer func() {
	// 		fmt.Println("Variable v in defer: ", v)
	// 	}()
	// 	fmt.Println("Variable v in loop: ", v)
	// }

	// vec21 := Vec2{1.0, 2.0}
	// vec22 := Vec2{3.0, 4.0}
	// vec21.Add(vec22)

	// go func(s string) {
	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s)
	// 	}
	// }("world")
	// // 主协程
	// for i := 0; i < 2; i++ {
	// 	// 切一下，再次分配任务
	// 	runtime.Gosched()
	// 	fmt.Println("hello")
	// }

	channelutils.TestChannel1()

}

type Vec2 struct {
	X, Y float32
}

// Add 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func maximumNumber(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func namedReturnValueFunction() (names []string, m map[string]string, i int) {
	names = make([]string, 10)
	names = append(names, "First item in slice")

	m = make(map[string]string, 20)
	m["lkl"] = "handsome"

	return
}
