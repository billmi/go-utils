package floatformat

import (
	"fmt"
	"strconv"
)

/**
  浮点这里需要仔细看!!!
  @author Bill
*/

func Float64ToString(f float64) string {
	return fmt.Sprintf("%g", f)
}

func StringToFloat64(val string) float64 {
	currVal, _ := strconv.ParseFloat(val, 64)
	return currVal
}

//四舍五入
func Float64Rand(v float64, dig int) float64 {
	cDig := strconv.Itoa(dig)
	val := fmt.Sprintf("%0."+cDig+"f", v)
	return StringToFloat64(val)
}

//浮点数串化(左边是整数位置,右边是小数位,dig参数控制)
func FloatToFDig(floVal float64, dig int) string {
	return fmt.Sprintf("%10."+strconv.Itoa(dig)+"f", floVal) //十位整数，8位小数
}
