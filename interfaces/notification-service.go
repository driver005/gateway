package interfaces

type ReturnedData struct {
	To     string
	Status string
	Data   map[string]interface{}
}

type INotificationService interface {
	SendNotification(event string, data interface{}, attachmentGenerator interface{}) (ReturnedData, error)
	ResendNotification(notification interface{}, config interface{}, attachmentGenerator interface{}) (ReturnedData, error)
}

func IsNotificationService(object interface{}) bool {
	_, ok := object.(INotificationService)
	return ok
}
