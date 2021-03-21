package generator

import (
	"KafkaMessagesGenerator/model"
	"time"
)
import "math/rand"

var Matches = [] string { "AR - US", "BR - FR", "SP - PR", "EN - DE", "PO - CH" }
var EventToScore = make(map[string]model.Score)
func GetEvent() model.Event {
	var matchIndex = rand.Intn(len(Matches))
	var match = Matches[matchIndex]
	var score model.Score
	if contains, ok := EventToScore[match]; ok {
		if rand.Intn(2) == 1 {
			contains.Home++
		} else {
			contains.Away++
		}
		score = contains
	} else {
		score = model.Score{Home: 0, Away: 0}
		EventToScore[match] = score
	}
	return model.Event{Match: match, Score: score, Time: time.Now()}
}
