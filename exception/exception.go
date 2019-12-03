package recover

/**
ThrowException
@author : Bill
异常统一抛出处理
*/

func ThrowException(execCall func()) {
	defer func() {
		if err := recover(); err != nil {
			//Do Something
		}
	}()
	execCall()
}
