package main

import (
	"context"
	"fmt"
	"kafka-bet-generator/effect"
	"strings"
)

func Handler(ctx context.Context, question Question) (*ResponseEntity, error) {
	fmt.Println(ctx)

	var text = question.Request.OriginalUtterance
	var voice = text
	if strings.Contains(strings.ToLower(question.Request.OriginalUtterance), "олю") {
		text = "Привет, Оля"
		voice = effect.AddRandom(text)
	}
	return &ResponseEntity{
			Version: "1.0",
			Response: Response{
				EndSession: false,
				Text:       text,
				Tts:        voice,
			},
		},
		nil
}
