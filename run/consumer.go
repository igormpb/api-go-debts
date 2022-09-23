package run

import "github.com/igormpb/api-go-debts/modules/rabbitmq"

func Consumer() {
	rabbitmq.Sub()
}
