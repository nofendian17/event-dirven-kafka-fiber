package cmd

import (
	"github.com/spf13/cobra"
	"kafka-fibre/config"
	"kafka-fibre/worker"
)


func init()  {
	command = &cobra.Command{
		Use:   "consume",
		Short: "consume event from topic",
		Long:  `Command to consume topic`,
		Run: func(cmd *cobra.Command, args []string) {
			configuration := config.New()
			kafkaReader := config.NewKafkaReader(configuration)
			consumer := worker.NewKafkaConsumer(kafkaReader)
			consumer.Consume()
		},
	}
	rootCmd.AddCommand(command)
}