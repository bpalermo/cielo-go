package client

type TestLogger struct {
}

func (l *TestLogger) Error(_ ...interface{}) {
}

func (l *TestLogger) Errorf(_ string, _ ...interface{}) {
}
