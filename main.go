package main

import (
	"KafkaMessagesGenerator/generator"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

func main() {
	topicEvent := "event-topic"
	topicBet := "bet-topic"
	partition := 1

	connBet, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topicBet, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	connEvent, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topicEvent,
		partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go writeBet(connBet)
	go writeEvent(connEvent)
	wg.Wait()

	if err := connBet.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	if err := connEvent.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func writeBet(conn *kafka.Conn) {
	for {
		var bet = generator.GetBet()
		b, err := json.Marshal(bet)
		if checkErr(err) {
			return
		}
		_, err = conn.WriteMessages(
			kafka.Message{Value: b},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func writeEvent(conn *kafka.Conn) {
	for {
		var event = generator.GetEvent()
		e, err := json.Marshal(event)
		if checkErr(err) {
			return
		}
		_, err = conn.WriteMessages(
			kafka.Message{Value: e},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		time.Sleep(10000 * time.Millisecond)
	}
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}