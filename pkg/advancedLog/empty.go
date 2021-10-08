package advancedLog

type EmptyLogger struct{}

func NewEmptyLogger() *EmptyLogger { return &EmptyLogger{} }

func (l EmptyLogger) Debug(v ...interface{}) {}
func (l EmptyLogger) Info(v ...interface{})  {}
func (l EmptyLogger) Warn(v ...interface{})  {}
