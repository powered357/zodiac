package zodiac_test

import (
	"backend/app/smspush/zodiac"
	"testing"
	"time"
)

const shortForm = "2006-01-02"

func TestGetChineseZodiacSign(t *testing.T) {
	tests := []struct {
		time           string
		expectedSignZh string
	}{
		{"1963-12-18", "射手座"},
		{"1963-11-11", "天蝎座"},
	}

	for _, test := range tests {
		testTime, _ := time.Parse(shortForm, test.time)
		sign := zodiac.GetChineseSign(testTime)
		if sign != test.expectedSignZh {
			t.Errorf("got value(%v) != (%v)", sign, test.expectedSignZh)
		}
	}
}
