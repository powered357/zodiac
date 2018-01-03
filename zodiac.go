// Package zodiac implements zodiac sign parse on birthdate
package zodiac

import (
	"fmt"
	"time"
)

// SignEn represent the zodiac signs in english
type SignEn int

const (
	// Aries zodiac sign
	Aries SignEn = iota
	// Taurus zodiac sign
	Taurus
	// Gemini zodiac sign
	Gemini
	// Cancer zodiac sign
	Cancer
	// Leo zodiac sign
	Leo
	// Virgo zodiac sign
	Virgo
	// Libra zodiac sign
	Libra
	// Scorpio zodiac sign
	Scorpio
	// Sagittarius zodiac sign
	Sagittarius
	// Capricorn zodiac sign
	Capricorn
	// Aquarius zodiac sign
	Aquarius
	// Pisces zodiac sign
	Pisces
)

// ZodiacEn2Zh represent the map from SignEn to chinese sign name
var ZodiacEn2Zh = map[SignEn]string{
	Aquarius:    "水瓶座",
	Aries:       "牡羊座",
	Cancer:      "巨蟹座",
	Capricorn:   "摩羯座",
	Gemini:      "双子座",
	Leo:         "狮子座",
	Libra:       "天秤座",
	Pisces:      "双鱼座",
	Sagittarius: "射手座",
	Scorpio:     "天蝎座",
	Taurus:      "金牛座",
	Virgo:       "处女座",
}

// GetChineseSign implements parse time to a chinese zodiac sign
func GetChineseSign(date time.Time) string {
	zodiacStr := CalculateZodiac(date)
	switch zodiacStr {
	case "aries":
		return ZodiacEn2Zh[Aries]
	case "taurus":
		return ZodiacEn2Zh[Taurus]
	case "gemini":
		return ZodiacEn2Zh[Gemini]
	case "cancer":
		return ZodiacEn2Zh[Cancer]
	case "leo":
		return ZodiacEn2Zh[Leo]
	case "virgo":
		return ZodiacEn2Zh[Virgo]
	case "libra":
		return ZodiacEn2Zh[Libra]
	case "scorpio":
		return ZodiacEn2Zh[Scorpio]
	case "sagittarius":
		return ZodiacEn2Zh[Sagittarius]
	case "capricorn":
		return ZodiacEn2Zh[Capricorn]
	case "aquarius":
		return ZodiacEn2Zh[Aquarius]
	case "pisces":
		return ZodiacEn2Zh[Pisces]
	}
	return "" // never run this based on util.CalculateZodiac
}

type zd struct {
	StartMonth time.Month
	StartDay   int
	EndMonth   time.Month
	EndDay     int
}

// ref http://en.wikipedia.org/wiki/Zodiac "Tropical Zodiac 2011"
var dateToZodiac = map[zd]string{
	zd{StartMonth: 3, StartDay: 21, EndMonth: 4, EndDay: 19}:   "aries",
	zd{StartMonth: 4, StartDay: 20, EndMonth: 5, EndDay: 20}:   "taurus",
	zd{StartMonth: 5, StartDay: 21, EndMonth: 6, EndDay: 21}:   "gemini",
	zd{StartMonth: 6, StartDay: 22, EndMonth: 7, EndDay: 22}:   "cancer",
	zd{StartMonth: 7, StartDay: 23, EndMonth: 8, EndDay: 22}:   "leo",
	zd{StartMonth: 8, StartDay: 23, EndMonth: 9, EndDay: 22}:   "virgo",
	zd{StartMonth: 9, StartDay: 23, EndMonth: 10, EndDay: 23}:  "libra",
	zd{StartMonth: 10, StartDay: 24, EndMonth: 11, EndDay: 22}: "scorpio",
	zd{StartMonth: 11, StartDay: 23, EndMonth: 12, EndDay: 21}: "sagittarius",
	zd{StartMonth: 12, StartDay: 22, EndMonth: 12, EndDay: 31}: "capricorn",
	zd{StartMonth: 1, StartDay: 1, EndMonth: 1, EndDay: 19}:    "capricorn",
	zd{StartMonth: 1, StartDay: 20, EndMonth: 2, EndDay: 18}:   "aquarius",
	zd{StartMonth: 2, StartDay: 19, EndMonth: 3, EndDay: 20}:   "pisces",
}

func CalculateZodiac(birthdate time.Time) string {
	for k, v := range dateToZodiac {
		start := time.Date(birthdate.Year(), k.StartMonth, k.StartDay, 0, 0, 0, 0, time.UTC)
		end := time.Date(birthdate.Year(), k.EndMonth, k.EndDay, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)

		utcBD := time.Date(birthdate.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.UTC)

		if !utcBD.Before(start) && !(utcBD.After(end) || utcBD.Equal(end)) {
			return v
		}
	}
	panic(fmt.Sprintf("Birthdate doesnt have zodiac! %v", birthdate))
}
