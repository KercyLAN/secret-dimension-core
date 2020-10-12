package extra

// 如果满足表达式返回satisfy，否则返回dissatisfaction
func If(expression bool, satisfy, dissatisfaction interface{}) interface{} {
	if expression {
		return satisfy
	}
	return dissatisfaction
}