// Package config loads application configuration.
package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/thmeitz/ksqldb-go/net"

	"github.com/openmeterio/openmeter/pkg/models"
)

// Configuration holds any kind of Configuration that comes from the outside world and
// is necessary for running the application.
// TODO: improve Configuration options
type Configuration struct {
	Address string

	Environment string

	// Telemetry configuration
	Telemetry TelemetryConfig

	// Namespace configuration
	Namespace NamespaceConfiguration

	// Ingest configuration
	Ingest struct {
		Kafka IngestKafkaConfiguration
	}

	// Dedupe configuration
	Dedupe DedupeConfiguration

	// SchemaRegistry configuration
	SchemaRegistry struct {
		URL      string
		Username string
		Password string
	}

	// Processor configuration
	Processor struct {
		KSQLDB     ProcessorKSQLDBConfiguration
		ClickHouse ProcessorClickhouseConfiguration
	}

	// Sink configuration
	Sink struct {
		KafkaConnect SinkKafkaConnectConfiguration
	}

	Meters []*models.Meter
}

// Validate validates the configuration.
func (c Configuration) Validate() error {
	if c.Address == "" {
		return errors.New("server address is required")
	}

	if err := c.Namespace.Validate(); err != nil {
		return err
	}

	if err := c.Ingest.Kafka.Validate(); err != nil {
		return err
	}

	if err := c.Processor.KSQLDB.Validate(); err != nil {
		return err
	}

	if err := c.Processor.ClickHouse.Validate(); err != nil {
		return err
	}

	if err := c.Sink.KafkaConnect.Validate(); err != nil {
		return err
	}

	if err := c.Dedupe.Validate(); err != nil {
		return fmt.Errorf("dedupe: %w", err)
	}

	if err := c.Telemetry.Validate(); err != nil {
		return fmt.Errorf("telemetry: %w", err)
	}

	for _, m := range c.Meters {
		// set default window size
		if m.WindowSize == "" {
			m.WindowSize = models.WindowSizeMinute
		}

		if err := m.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Namespace configuration
type NamespaceConfiguration struct {
	Default           string
	DisableManagement bool
}

func (c NamespaceConfiguration) Validate() error {
	if c.Default == "" {
		return errors.New("default namespace is required")
	}

	return nil
}

// Ingest Kafka configuration
type IngestKafkaConfiguration struct {
	Broker              string
	SecurityProtocol    string
	SaslMechanisms      string
	SaslUsername        string
	SaslPassword        string
	Partitions          int
	EventsTopicTemplate string
}

// CreateKafkaConfig creates a Kafka config map.
func (c IngestKafkaConfiguration) CreateKafkaConfig() kafka.ConfigMap {
	config := kafka.ConfigMap{
		"bootstrap.servers": c.Broker,

		// Required for logging
		"go.logs.channel.enable": true,
	}

	if c.SecurityProtocol != "" {
		config["security.protocol"] = c.SecurityProtocol
	}

	if c.SaslMechanisms != "" {
		config["sasl.mechanism"] = c.SaslMechanisms
	}

	if c.SaslUsername != "" {
		config["sasl.username"] = c.SaslUsername
	}

	if c.SaslPassword != "" {
		config["sasl.password"] = c.SaslPassword
	}

	return config
}

// Validate validates the configuration.
func (c IngestKafkaConfiguration) Validate() error {
	if c.Broker == "" {
		return errors.New("kafka broker is required")
	}

	if c.EventsTopicTemplate == "" {
		return errors.New("events topic template is required")
	}

	return nil
}

// KSQLDB Processor configuration
type ProcessorKSQLDBConfiguration struct {
	Enabled                     bool
	URL                         string
	Username                    string
	Password                    string
	DetectedEventsTopicTemplate string
}

// CreateKafkaConfig creates a Kafka config map.
func (c ProcessorKSQLDBConfiguration) CreateKSQLDBConfig() net.Options {
	config := net.Options{
		BaseUrl:   c.URL,
		AllowHTTP: true,
	}

	if strings.HasPrefix(c.URL, "https://") {
		config.AllowHTTP = false
	}

	if c.Username != "" || c.Password != "" {
		config.Credentials = net.Credentials{
			Username: c.Username,
			Password: c.Password,
		}
	}

	return config
}

// Validate validates the configuration.
func (c ProcessorKSQLDBConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.URL == "" {
		return errors.New("ksqldb URL is required")
	}

	if c.DetectedEventsTopicTemplate == "" {
		return errors.New("namespace detected events topic template is required")
	}

	return nil
}

// Clickhouse configuration
type ProcessorClickhouseConfiguration struct {
	Enabled  bool
	Address  string
	TLS      bool
	Database string
	Username string
	Password string
}

func (c ProcessorClickhouseConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Address == "" {
		return errors.New("clickhouse address is required")
	}

	return nil
}

// Sink configuration
type SinkKafkaConnectConfiguration struct {
	Enabled    bool
	URL        string
	ClickHouse KafkaSinkClickhouseConfiguration
}

func (c SinkKafkaConnectConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.URL == "" {
		return errors.New("kafka connect url is required")
	}

	if err := c.ClickHouse.Validate(); err != nil {
		return err
	}

	return nil
}

// Clickhouse configuration
// This may feel repetative but clikhouse sink and processor configs can be different,
// for example Kafka Connect ClickHouse plugin uses 8123 HTTP port while client uses native protocol's 9000 port.
// Hostname can be also different, as Kafka Connect and ClickHouse communicates inside the docker compose network.
// This why we default hostname in config to `clickhouse`.
type KafkaSinkClickhouseConfiguration struct {
	Hostname string
	Port     int
	SSL      bool
	Database string
	Username string
	Password string
}

func (c KafkaSinkClickhouseConfiguration) Validate() error {
	if c.Hostname == "" {
		return errors.New("kafka sink clickhouse hostname is required")
	}
	if c.Port == 0 {
		return errors.New("kafka sink clickhouse port is required")
	}
	if c.Database == "" {
		return errors.New("kafka sink clickhouse database is required")
	}
	if c.Username == "" {
		return errors.New("kafka sink clickhouse username is required")
	}

	return nil
}

// Configure configures some defaults in the Viper instance.
func Configure(v *viper.Viper, flags *pflag.FlagSet) {
	// Viper settings
	v.AddConfigPath(".")

	// Environment variable settings
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Server configuration
	flags.String("address", ":8888", "Server address")
	_ = v.BindPFlag("address", flags.Lookup("address"))
	v.SetDefault("address", ":8888")

	// Environment used for identifying the service environment
	v.SetDefault("environment", "unknown")

	configureTelemetry(v, flags)

	// Namespace configuration
	v.SetDefault("namespace.default", "default")
	v.SetDefault("namespace.disableManagement", false)

	// Ingest configuration
	v.SetDefault("ingest.kafka.broker", "127.0.0.1:29092")
	v.SetDefault("ingest.kafka.securityProtocol", "")
	v.SetDefault("ingest.kafka.saslMechanisms", "")
	v.SetDefault("ingest.kafka.saslUsername", "")
	v.SetDefault("ingest.kafka.saslPassword", "")
	v.SetDefault("ingest.kafka.partitions", 1)
	v.SetDefault("ingest.kafka.eventsTopicTemplate", "om_%s_events")

	// Schema Registry configuration
	v.SetDefault("schemaRegistry.url", "")
	v.SetDefault("schemaRegistry.username", "")
	v.SetDefault("schemaRegistry.password", "")

	// Processor ksqlDB configuration
	v.SetDefault("processor.ksqldb.enabled", true)
	v.SetDefault("processor.ksqldb.url", "http://127.0.0.1:8088")
	v.SetDefault("processor.ksqldb.username", "")
	v.SetDefault("processor.ksqldb.password", "")
	v.SetDefault("processor.ksqldb.detectedEventsTopicTemplate", "om_%s_detected_events")

	// Processor Clickhouse configuration
	v.SetDefault("processor.clickhouse.enabled", false)
	v.SetDefault("processor.clickhouse.address", "127.0.0.1:9000")
	v.SetDefault("processor.clickhouse.tls", false)
	v.SetDefault("processor.clickhouse.database", "default")
	v.SetDefault("processor.clickhouse.username", "default")
	v.SetDefault("processor.clickhouse.password", "")

	// Sink Kafka Connect configuration
	v.SetDefault("sink.kafkaConnect.enabled", false)
	v.SetDefault("sink.kafkaConnect.url", "http://127.0.0.1:8083")
	v.SetDefault("sink.kafkaConnect.clickhouse.hostname", "clickhouse")
	v.SetDefault("sink.kafkaConnect.clickhouse.port", 8123)
	v.SetDefault("sink.kafkaConnect.clickhouse.ssl", false)
	v.SetDefault("sink.kafkaConnect.clickhouse.database", "default")
	v.SetDefault("sink.kafkaConnect.clickhouse.username", "default")
	v.SetDefault("sink.kafkaConnect.clickhouse.password", "")

	configureDedupe(v)
}
