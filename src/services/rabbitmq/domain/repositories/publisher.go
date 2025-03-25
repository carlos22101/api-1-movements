package repositories

type MovementPublisher interface {
    Publish(message interface{}) error
}
