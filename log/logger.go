package log

//go:generate mockgen -destination=mocks/mocklogger.go -package=mocks . Logger
type Logger interface {
	Info(string)
	Warn(string)
	Error(string)
}
