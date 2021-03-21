package generator

import (
	"KafkaMessagesGenerator/model"
	"math/rand"
	"time"
)

var Names = []string{"Adams", "Baker", "Clark", "Davis", "Evans", "Frank", "Ghosh", "Hills", "Irwin", "Jones", "Klein",
	"Lopez", "Mason", "Nalty", "Ochoa", "Patel", "Quinn", "Reily", "Smith", "Trott", "Usman", "Valdo", "White",
	"Xiang", "Yakub", "Zafar"}

func GetBet() model.Bet {
	var nameIndex = rand.Intn(len(Names))
	var matchIndex  = rand.Intn(len(Matches))
	var amount = rand.Int()
	return model.Bet{
		Bettor: Names[nameIndex],
		Match:  Matches[matchIndex],
		Amount: amount,
		Time:   time.Now(),
	}

}
