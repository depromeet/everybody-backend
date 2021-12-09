package util

import (
	"github.com/depromeet/everybody-backend/rest-api/er"
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

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	// ct에 저장
	t.Run("성공", func(t *testing.T) {
		tmp := &CustomTime{}
		err := tmp.UnmarshalJSON([]byte("\"2021-12-02 20:30:00\""))
		assert.NoError(t, err)
	})

	t.Run("실패) 빈 값", func(t *testing.T) {
		tmp := &CustomTime{}
		target := &er.WrongTimeFormatErr{}
		err := tmp.UnmarshalJSON([]byte(""))
		assert.ErrorAs(t, err, &target)
		wrongTimeFormatErr := err.(*er.WrongTimeFormatErr)
		// 참고: Error 인터페이스는 %s 시에String()메소드가 아니라 Error() 메소드를 이용하네
		t.Log(wrongTimeFormatErr)
		t.Logf("%+v\n", wrongTimeFormatErr)
	})

	t.Run("실패) 이상한 값", func(t *testing.T) {
		tmp := &CustomTime{}
		target := &er.WrongTimeFormatErr{}
		err := tmp.UnmarshalJSON([]byte("\"hello, world\""))
		assert.ErrorAs(t, err, &target)
		wrongTimeFormatErr := err.(*er.WrongTimeFormatErr)
		// 참고: Error 인터페이스는 %s 시에String()메소드가 아니라 Error() 메소드를 이용하네
		t.Log(wrongTimeFormatErr)
		t.Logf("%+v\n", err)
	})

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
