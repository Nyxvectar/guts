/**
 * Author:  Nyxvectar Yan
 * Repo:    go-zju-formulas
 * Created: 07/23/2025
 */

// Package algebran 提供代数运算的基本功能
// 导入路径: "go-zju-formulas/algebran"
package algebran

import (
	"errors"
	"math"
)

// valid 验证对数运算的底数和真数是否合法
// 参数要求：
//   - 底数必须大于0且不等于1
//   - 真数必须大于0
func valid(base, x float64) (bool, error) {
	if base <= 0 || base == 1 {
		return false, errors.New("对数的底须大于0且不为1")
	}
	if x <= 0 {
		return false, errors.New("对数的真数须大于0")
	}
	return true, nil
}

// LogBaseConvert 计算以base为底x的对数
// 返回值为 logₐx，其中a是底数，x是真数
// 计算公式：logₐx = ln(x) / ln(a)
func LogBaseConvert(base, x float64) (float64, error) {
	if ok, err := valid(base, x); !ok {
		return 0, err
	}
	return math.Log(x) / math.Log(base), nil
}

// LogEquation 求解对数方程 logₐx = y 中的y值
// 返回值为对数方程的解 y
// 注意：原函数实现有误，已修正为返回正确的对数值
func LogEquation(base, x float64) (float64, error) {
	if ok, err := valid(base, x); !ok {
		return 0, err
	}
	return math.Log(x) / math.Log(base), nil
}

// GrowthAvg 计算增长率
// 返回值为增长率：(现期值-基期值)/基期值
// 参数要求：
//   - 基期值(previous)不能为0
func GrowthAvg(present, previous float64) (float64, error) {
	if previous == 0 {
		return 0, errors.New("基期值不得为零")
	}
	return (present - previous) / previous, nil
}
