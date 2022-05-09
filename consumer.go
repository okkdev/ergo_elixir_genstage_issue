package main

import (
	"fmt"
	"time"

	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

type Consumer struct {
	gen.Stage
}

func (c *Consumer) InitStage(process *gen.StageProcess, args ...etf.Term) (gen.StageOptions, error) {
	var opts = gen.StageSubscribeOptions{
		MinDemand: 0,
		MaxDemand: 5,
	}

	fmt.Println("Subscribe consumer", process.Name(), "[", process.Self(), "]",
		"with min events =", opts.MinDemand,
		"and max events", opts.MaxDemand)

	process.Subscribe(gen.ProcessID{Name: "Elixir.Producer", Node: "producer@localhost"}, opts)

	return gen.StageOptions{}, nil
}

func (c *Consumer) HandleEvents(process *gen.StageProcess, subscription gen.StageSubscription, events etf.List) gen.StageStatus {
	time.Sleep(1 * time.Second)
	fmt.Printf("Consumer '%s' got events: %v\n", process.Name(), events)
	return gen.StageStatusOK
}

func main() {
	// create nodes for producer and consumers
	fmt.Println("Starting nodes 'consumer@localhost'")
	node, _ := ergo.StartNode("consumer@localhost", "cookiesecret", node.Options{})

	// create producer and consumer objects
	consumer := &Consumer{}

	fmt.Println("Spawn consumer on 'consumer@localhost'")
	_, err := node.Spawn("consumer", gen.ProcessOptions{}, consumer)
	if err != nil {
		panic(err)
	}

	node.Wait()
}
