package sarama

type serializer struct {
	config *Config
}

type Serializer interface {
	Serialize(msg *ProducerMessage) (*ProducerMessage, error)
}

func NewSerializer(config *Config) (*serializer, error) {
	serializer := &serializer{
		config:    config,
	}
	return serializer, nil
}

func (serializer *serializer) Serialize(msg *ProducerMessage) (*ProducerMessage, error) {
	return msg, nil
}