package effect

import "math/rand"

type Effect string

const (
	Hamster       Effect = "hamster"
	BehindTheWall Effect = "behind_the_wall"
	Megaphone     Effect = "megaphone"
	PitchDown     Effect = "pitch_down"
	Psychodelic   Effect = "psychodelic"
	Pulse         Effect = "pulse"
	TrainAnnounce Effect = "train_announce"
)

var voices = []Effect{
	Hamster,
	BehindTheWall,
	Megaphone,
	PitchDown,
	Psychodelic,
	Pulse,
	TrainAnnounce,
}

// Add wraps text with proper tags using given effect.
//
// Example:
//
//	...
//	"tts": "Hello, " + effects.Add(effects.Megaphone, "world!"),
//	...
func Add(e Effect, s string) string {
	return `<speaker effect="` + string(e) + `">` + s + `<speaker effect="-">`
}

func AddRandom(s string) string {

	return `<speaker effect="` + string(voices[rand.Intn(len(voices))]) + `">` + s + `<speaker effect="-">`
}
