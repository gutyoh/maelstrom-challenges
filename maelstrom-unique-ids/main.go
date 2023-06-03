package main

import (
	"encoding/json"
	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		var req map[string]any
		err := json.Unmarshal(msg.Body, &req)
		if err != nil {
			log.Fatal(err)
		}

		if req["type"] == "generate" {
			res := map[string]any{
				"type": "generate_ok",
				"id":   uuid.New().String(),
			}
			return n.Reply(msg, res)
		}
		return nil
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
