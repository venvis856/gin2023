package handlePanic

import "gin/internal/global"

func Panic(errs string) {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Daily("run", "error", err)
		}
	}()
	panic(errs)
}
