package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// CheckIDCard 验证身份证号码
func CheckIDCard(id string) bool {
	// 身份证号码为18位，最后一位可以是数字或字母X
	match, _ := regexp.MatchString(`^[1-9]\d{16}[\dXx]$`, id)
	return match
}

// CalculateAge 从身份证计算年龄
func CalculateAge(idCard string) (int, error) {
	if len(idCard) != 18 {
		return 0, fmt.Errorf("身份证号码长度不正确")
	}

	// 提取出生年月日
	birthYear, err := strconv.Atoi(idCard[6:10])
	if err != nil {
		return 0, fmt.Errorf("无法解析出生年份")
	}
	birthMonth, err := strconv.Atoi(idCard[10:12])
	if err != nil {
		return 0, fmt.Errorf("无法解析出生月份")
	}
	birthDay, err := strconv.Atoi(idCard[12:14])
	if err != nil {
		return 0, fmt.Errorf("无法解析出生日")
	}

	// 当前日期
	now := time.Now()

	// 计算年龄
	age := now.Year() - birthYear
	if now.Month() < time.Month(birthMonth) || (now.Month() == time.Month(birthMonth) && now.Day() < birthDay) {
		age--
	}

	return age, nil
}

// ExtractBirthYearMonth 从身份证上提取出生日期
func ExtractBirthYearMonth(id string) (int, int, error) {
	if len(id) != 18 {
		return 0, 0, fmt.Errorf("invalid ID length")
	}

	year, err := strconv.Atoi(id[6:10])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing year: %v", err)
	}

	month, err := strconv.Atoi(id[10:12])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing month: %v", err)
	}

	return year, month, nil
}
