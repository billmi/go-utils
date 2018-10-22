package exception


/**
	异常捕获(暂时还没完善)
	@author Bill
 */
func Try(fun func(), handler func()) {
	defer func() {
		if err := recover(); err != nil {
			handler()
		}
	}()
	fun()
}