package r

import "log"

// Go 统一处理panic错误
func Go(f func(val any), val any) {
	go func(val any) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		f(val)
	}(val)
}
