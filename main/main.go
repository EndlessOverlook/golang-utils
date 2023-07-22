package main

import (
	"fmt"
	"golang-utils/age-utils"
	"golang-utils/name_utils"
	"golang-utils/timeutils"
	"strconv"
)

func verifyGolangPackageName() {
	// 验证包名是否可以带中横线和下划线
	fmt.Println("Current Year is : " + strconv.Itoa(timeutils.GetYear()))
	fmt.Println("NickName is : " + name_utils.GetNickName())
	fmt.Println("Age is : " + strconv.Itoa(age_utils.GetAge()))
}

func main() {
	verifyGolangPackageName()
}
