package sarama

type serde struct {
   config *Config;
   serializer *serializer;
   deserializer *deserializer;
}

type Serde interface {
	Serializer() *serializer
	Deserializer() *deserializer
}

func NewSerde(config *Config) (*serde, error) {
	serde := &serde{
		config:     config,
	}
	serializer, err := NewSerializer(config);
	if (err == nil) {
        serde.serializer = serializer
	}
	deserializer, err := NewDeserializer(config);
	if (err == nil) {
        serde.deserializer = deserializer
	}
	return serde, err
}

func (serde *serde) Serializer() (*serializer) {
	return serde.serializer
}

func (serde *serde) Deserializer() (*deserializer) {
	return serde.deserializer
}
