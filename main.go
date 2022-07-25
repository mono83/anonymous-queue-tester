package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"os"
)

var (
	mainCmdAmqpUrl    string
	mainCmdDurable    bool
	mainCmdAutoDelete bool
	mainCmdExclusive  bool
)

var mainCMD = cobra.Command{
	Use:   "test",
	Short: "RabbitMQ anonymous queue tester",
	RunE: func(_ *cobra.Command, args []string) error {
		fmt.Print("Connecting to RabbitMQ ... ")
		connection, err := amqp.Dial(mainCmdAmqpUrl)
		if err != nil {
			return err
		}
		defer connection.Close()
		fmt.Println("OK")

		fmt.Print("Creating communication channel ... ")
		channel, err := connection.Channel()
		if err != nil {
			return err
		}
		defer channel.Close()
		fmt.Println("OK")

		fmt.Println("Creating queue")
		fmt.Println("Durable     ...", mainCmdDurable)
		fmt.Println("Auto Delete ...", mainCmdAutoDelete)
		fmt.Println("Exclusive   ...", mainCmdExclusive)
		q, err := channel.QueueDeclare("", mainCmdDurable, mainCmdAutoDelete, mainCmdExclusive, false, nil)
		if err != nil {
			return err
		}
		fmt.Println("Creating    ...", q.Name)
		for {
		}
	},
}

func init() {
	mainCMD.Flags().StringVar(&mainCmdAmqpUrl, "addr", "amqp://guest:guest@localhost:5672/", "")
	mainCMD.Flags().BoolVarP(&mainCmdDurable, "durable", "d", false, "")
	mainCMD.Flags().BoolVarP(&mainCmdAutoDelete, "autodelete", "a", false, "")
	mainCMD.Flags().BoolVarP(&mainCmdExclusive, "exclusive", "e", false, "")
}

func main() {
	if err := mainCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
