package rabbitmq

func StartConsumers() {
	go func() {
		StartProjectsBriefConsumer()
	}()
}

func Start() {
	Connect()
	StartConsumers()
}
