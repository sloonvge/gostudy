package operation


func CreatOperate(opExpr string) (op Operate) {
	switch opExpr {
	case "+":
		op = NewOperationAdd()
	case "-":
		op = NewOperationSub()
	case "*":
		op = NewOperationMul()
	case "/":
		op = NewOperationDiv()
	default:
		panic("不允许的操作符!")
	}
	return
}
