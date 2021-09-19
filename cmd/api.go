package cmd

import (
	"fmt"
	healthcheck "github.com/aschenmaker/fiber-health-check"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"kafka-fibre/config"
	"kafka-fibre/controller"
	"kafka-fibre/exception"
	"kafka-fibre/repository"
	"kafka-fibre/service"
)

var command *cobra.Command

func init()  {
	command = &cobra.Command{
		Use:   "serve",
		Short: "serve api",
		Long:  `Command to serve application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running serve command")

			configuration := config.New()

			kafkaWriter := config.NewKafkaWriter(configuration)

			defer func(kafkaWriter *kafka.Writer) {
				err := kafkaWriter.Close()
				if err != nil {
					exception.PanicIfNeeded(err)
				}
			}(kafkaWriter)

			eventRepository := repository.NewProduceRepository(kafkaWriter)

			eventService := service.NewProduceService(&eventRepository)

			eventController := controller.NewEventController(&eventService)

			app := fiber.New(config.NewFiberConfig())
			app.Use(
				healthcheck.New(),
				recover.New(),
				logger.New(logger.Config{
					Format:     "[${time}] ${status} - ${latency} ${method} ${path} \n",
					TimeFormat: "02-Jan-2006",
					TimeZone:   "Asia/Jakarta",
				}),
				requestid.New(),
			)
			app.Get("/dashboard", monitor.New())

			eventController.Route(app)

			err := app.Listen(":3000")
			exception.PanicIfNeeded(err)
		},
	}
	rootCmd.AddCommand(command)
}
