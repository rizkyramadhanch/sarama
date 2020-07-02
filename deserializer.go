package sarama

type deserializer struct {
	config *Config
}

type Deserializer interface {
	Deserialize(msg *ConsumerMessage) (*ConsumerMessage, error)
}

func NewDeserializer(config *Config) (*deserializer, error) {
	deserializer := &deserializer{
		config:    config,
	}
	return deserializer, nil
}

func (deserializer *deserializer) Deserialize(msg *ConsumerMessage) (*ConsumerMessage, error) {
	return msg, nil
}
