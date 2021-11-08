package scenarios

import "github.com/go-kratos/kratos/v2/log"

type dummyLogger struct {
}

func (d *dummyLogger) Log(level log.Level, keyvals ...interface{}) error {
	return nil
}
