package main

import (
	"github.com/Songmu/flextime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetThisMondayBySunday(t *testing.T) {
	t.Run("monday test", func(t *testing.T) {
		mockTime := time.Date(2020, 6, 7, 1, 2, 3, 123456000, time.UTC)
		restore := flextime.Set(mockTime)
		defer restore()
		expected := getThisMonday(Option{
			Ago: 0,
		})
		assert.Equal(t, expected.Format("2006/01/02"), "2020/06/01")
	})
}

func TestGetThisMondayBySaturday(t *testing.T) {
	t.Run("monday test", func(t *testing.T) {
		mockTime := time.Date(2020, 6, 6, 1, 2, 3, 123456000, time.UTC)
		restore := flextime.Set(mockTime)
		defer restore()
		expected := getThisMonday(Option{
			Ago: 0,
		})
		assert.Equal(t, expected.Format("2006/01/02"), "2020/06/01")
	})
}

func TestGetThisMondayByMonday(t *testing.T) {
	t.Run("monday test", func(t *testing.T) {
		mockTime := time.Date(2020, 6, 1, 1, 2, 3, 123456000, time.UTC)
		restore := flextime.Set(mockTime)
		defer restore()
		expected := getThisMonday(Option{
			Ago: 0,
		})
		assert.Equal(t, expected.Format("2006/01/02"), "2020/06/01")
	})
}

func TestGetThisMondayByMondayAgoOne(t *testing.T) {
	t.Run("monday test", func(t *testing.T) {
		mockTime := time.Date(2020, 6, 1, 1, 2, 3, 123456000, time.UTC)
		restore := flextime.Set(mockTime)
		defer restore()
		expected := getThisMonday(Option{
			Ago: 1,
		})
		assert.Equal(t, expected.Format("2006/01/02"), "2020/05/25")
	})
}
