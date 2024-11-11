package parser

import (
	"cw-q3arena/entities"
	"cw-q3arena/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserEventMatcher(t *testing.T) {
	t.Run("should match kill event including names with spaces", func(t *testing.T) {
		parser := New()
		event, result, err := parser.Parse("2:22 Kill: 3 2 10: Isgalamido killed Dono da Bola by MOD_RAILGUN")
		assert.Equal(t, events.EventKill, event)
		assert.NoError(t, err)

		assert.Equal(t, entities.Kill{
			Timestamp:  "2:22",
			KillerId:   3,
			KillerName: "Isgalamido",
			VictimId:   2,
			VictimName: "Dono da Bola",
			MethodId:   10,
			MethodName: "MOD_RAILGUN",
		}, result)
	})

	t.Run("should match kill event including world", func(t *testing.T) {
		parser := New()
		event, result, err := parser.Parse("21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT")
		assert.Equal(t, events.EventKill, event)
		assert.NoError(t, err)

		assert.Equal(t, entities.Kill{
			Timestamp:  "21:42",
			KillerId:   1022,
			KillerName: "<world>",
			VictimId:   2,
			VictimName: "Isgalamido",
			MethodId:   22,
			MethodName: "MOD_TRIGGER_HURT",
		}, result)
	})

	t.Run("should not match kill when wrong event", func(t *testing.T) {
		parser := New()
		event, _, err := parser.Parse("Fake event")
		assert.Equal(t, events.EventUnknown, event)
		assert.NoError(t, err)

	})

	t.Run("should return unknown when an error occurred while parsing", func(t *testing.T) {
		parser := New()
		event, _, err := parser.Parse("")
		assert.Equal(t, events.EventUnknown, event)
		assert.NoError(t, err)

	})
}
