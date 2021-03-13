package polling

import (
	"github.com/bwmarrin/discordgo"
)

type Poll struct {
	channel        string
	question       string
	pollMessage    *discordgo.Message
	runoffMessage  *discordgo.Message
	lastLetter     int
	emotes         map[int]string
	entries        map[string]string
	entriesReverse map[string]string
	isMovie        bool
}

var emotes map[int]string = map[int]string{
	0:  "🇦",
	1:  "🇧",
	2:  "🇨",
	3:  "🇩",
	4:  "🇪",
	5:  "🇫",
	6:  "🇬",
	7:  "🇭",
	8:  "🇮",
	9:  "🇯",
	10: "🇰",
	11: "🇱",
	12: "🇲",
	13: "🇳",
	14: "🇴",
	15: "🇵",
	16: "🇶",
	17: "🇷",
	18: "🇸",
	19: "🇹",
}
var entries map[string]string = map[string]string{
	"🇦": "🇦",
	"🇧": "🇧",
	"🇨": "🇨",
	"🇩": "🇩",
	"🇪": "🇪",
	"🇫": "🇫",
	"🇬": "🇬",
	"🇭": "🇭",
	"🇮": "🇮",
	"🇯": "🇯",
	"🇰": "🇰",
	"🇱": "🇱",
	"🇲": "🇲",
	"🇳": "🇳",
	"🇴": "🇴",
	"🇵": "🇵",
	"🇶": "🇶",
	"🇷": "🇷",
	"🇸": "🇸",
	"🇹": "🇹",
}

func New() Poll {
	poll := Poll{lastLetter: 0, emotes: copyIntMap(emotes), entries: copyMap(entries), entriesReverse: map[string]string{}}
	//create emotes
	return poll
}

func copyMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
func copyIntMap(m map[int]string) map[int]string {
	cp := make(map[int]string)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
