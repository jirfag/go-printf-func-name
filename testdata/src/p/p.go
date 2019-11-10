package p

func notPrintfFuncAtAll() {}

func funcWithEllipsis(args ...interface{}) {}

func printfLikeButWithStrings(format string, args ...string) {}

func printfLikeButWithBadFormat(format int, args ...interface{}) {}

func secondArgIsNotEllipsis(arg1 string, arg2 int) {}

func printfLikeButWithExtraInterfaceMethods(format string, args ...interface {
	String() string
}) {
}

func prinfLikeFuncf(format string, args ...interface{}) {}

func prinfLikeFunc(format string, args ...interface{}) {} // want "printf-like formatting function"

func prinfLikeFuncWithExtraArgs1(extraArg, format string, args ...interface{}) {} // want "printf-like formatting function"

func prinfLikeFuncWithExtraArgs2(extraArg int, format string, args ...interface{}) {} // want "printf-like formatting function"

func prinfLikeFuncWithReturnValue(format string, args ...interface{}) string { // want "printf-like formatting function"
	return ""
}
