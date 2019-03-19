package burrow_exporter

import "github.com/prometheus/client_golang/prometheus"

// If we are missing a status, it will return 0
var Status = map[string]int{
	"NOTFOUND": 1,
	"OK":       2,
	"WARN":     3,
	"ERR":      4,
	"STOP":     5,
	"STALL":    6,
	"REWIND":   7,
}

var (
	KafkaConsumerTopicLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_topic_lag",
			Help: "The sum lag of the latest offset commit on topics as reported by burrow.",
		},
		[]string{"cluster", "group", "topic"},
	)
	KafkaConsumerPartitionLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_lag",
			Help: "The lag of the latest offset commit on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition"},
	)
	KafkaConsumerPartitionCurrentOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_current_offset",
			Help: "The latest offset commit on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition"},
	)
	KafkaConsumerPartitionCurrentStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_status",
			Help: "The status of a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition"},
	)
	KafkaConsumerPartitionMaxOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_max_offset",
			Help: "The log end offset on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition"},
	)
	KafkaConsumerTotalLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_total_lag",
			Help: "The total amount of lag for the consumer group as reported by burrow",
		},
		[]string{"cluster", "group"},
	)
	KafkaConsumerStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_status",
			Help: "The status of a partition as reported by burrow.",
		},
		[]string{"cluster", "group"},
	)
	KafkaTopicSumOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_topic_offset",
			Help: "The sum of latest offset on a topic's partition as reported by burrow.",
		},
		[]string{"cluster", "topic"},
	)
	KafkaTopicPartitionOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_topic_partition_offset",
			Help: "The latest offset on a topic's partition as reported by burrow.",
		},
		[]string{"cluster", "topic", "partition"},
	)
)

func init() {
	prometheus.MustRegister(KafkaTopicSumOffset)
	prometheus.MustRegister(KafkaConsumerTopicLag)
	prometheus.MustRegister(KafkaConsumerPartitionLag)
	prometheus.MustRegister(KafkaConsumerPartitionCurrentOffset)
	prometheus.MustRegister(KafkaConsumerPartitionCurrentStatus)
	prometheus.MustRegister(KafkaConsumerPartitionMaxOffset)
	prometheus.MustRegister(KafkaConsumerTotalLag)
	prometheus.MustRegister(KafkaConsumerStatus)
	prometheus.MustRegister(KafkaTopicPartitionOffset)
}
