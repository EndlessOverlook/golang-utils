package main

import (
	"fmt"
	"golang-utils/age-utils"
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
