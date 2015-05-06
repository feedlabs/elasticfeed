package model

type EventManager interface {

	Init()

	InstallSchedule(string, string, func() error) error
}
