package main

import (
	"fmt"
	"math"
)

/*
实现 atoi，将字符串转为整数。

该函数首先根据需要丢弃任意多的空格字符，直到找到第一个非空格字符为止。如果第一个非空字符是正号或负号，选取该符号，并将其与后面尽可能多的连续的数字组合起来，这部分字符即为整数的值。
如果第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

字符串可以在形成整数的字符后面包括多余的字符，这些字符可以被忽略，它们对于函数没有影响。

当字符串中的第一个非空字符序列不是个有效的整数；或字符串为空；或字符串仅包含空白字符时，则不进行转换。

若函数不能执行有效的转换，返回 0。

说明：

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。如果数值超过可表示的范围，则返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
*/
func myAtoi(str string) int {
	isMinus := false
	sum := 0
	hasSymbol := false
	hasSpace := false
	hasBeginZero := false
	var s []byte
	for i := 0; i < len(str); i++ {
		if str[i] != 32 {
			hasSpace = true
			if (str[i] == '-' || str[i] == '+') && !hasSymbol {
				s = append(s, str[i])
				hasSymbol = true
			} else if isNumber(str[i]) {
				s = append(s, str[i])
				hasSymbol = true
			} else {
				break
			}
		} else {
			if hasSpace {
				break
			}
		}
	}
	lenS := len(s)
	if lenS == 0 {
		return 0
	}
	if !isNumber(s[0]) {
		if s[0] == '-' {
			isMinus = true
		}
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v != '0' {
			hasBeginZero = true
			temp := math.Pow10(len(s) - 1 - i)
			sum += int(temp) * int(v-'0')
		} else {
			if !hasBeginZero {
				lenS--
			}
		}
	}
	if lenS >= 19 || sum > math.MaxInt32 {
		if isMinus {
			sum = math.MinInt32
		} else {
			sum = math.MaxInt32
		}
		return sum
	}
	if isMinus {
		sum = -sum
	}
	return sum
}

func myAtoiFast(str string) int {
	length := len(str)
	hasSymbol := false
	isStartZero := false
	MAX := 2 << 30
	var sum int64
	for i := 0; i < length; i++ {
		if str[i] == ' ' && !isStartZero {
			continue
		}
		if str[i] == '-' {
			if isStartZero {
				break
			}
			hasSymbol = true
			isStartZero = true
			continue
		}
		if str[i] == '+' {
			if isStartZero {
				break
			}
			isStartZero = true
			continue
		}
		if str[i] >= '0' && str[i] <= '9' {
			sum = sum*10 + int64(str[i]-'0')
			if sum > int64(MAX) {
				break
			}
			isStartZero = true
			continue
		}
		break

	}
	if hasSymbol && sum > int64(MAX) {
		return -MAX
	}
	if !hasSymbol && sum >= int64(MAX) {
		return MAX - 1
	}

	if hasSymbol {
		return int(-sum)
	}
	return int(sum)
}

func testMyAtoi() {
	s := "  -10000000001234567890"
	//s := "200"
	fmt.Println(myAtoi(s))
}
