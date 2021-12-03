package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ConvertIntToTime(t *testing.T) {
	// RFC3339 포맷 이용
	expected, err := time.Parse(time.RFC3339, "2021-08-17T12:00:00+09:00")
	assert.NoError(t, err)
	converted, err := ConvertIntToTime(2021, 8, 17)
	assert.Equal(t, expected, converted)

	_, err = ConvertIntToTime(0, 0, 0)
	assert.Error(t, err)
}

func Test_ConvertTimeToStr(t *testing.T) {
	// RFC3339 포맷 이용
	expected, err := time.Parse(time.RFC3339, "2021-08-17T12:03:00+09:00")
	assert.NoError(t, err)
	year, month, day, hour, min := ConvertTimeToStr(expected)
	assert.Equal(t, 2021, year)
	assert.Equal(t, 8, month)
	assert.Equal(t, 17, day)
	assert.Equal(t, 12, hour)
	assert.Equal(t, 03, min)
}

//func Test_ConvertTimeToStr(t *testing.T) {
//	// RFC3339 포맷 이용
//	expected, err := time.Parse(time.RFC3339, "2021-08-17T12:03:00+09:00")
//	assert.NoError(t, err)
//	year, month, day, hour, min := ConvertTimeToStr(expected)
//	assert.Equal(t, 2021, year)
//	assert.Equal(t, 8, month)
//	assert.Equal(t, 17, day)
//	assert.Equal(t, 12, hour)
//	assert.Equal(t, 03, min)
//}
