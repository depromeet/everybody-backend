package dto

import (
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_getLatestPart(t *testing.T) {
	baseCreatedAt := time.Now().Add(-time.Hour)
	album := &ent.Album{Edges: ent.AlbumEdges{
		Picture: []*ent.Picture{
			{
				BodyPart:  "whole",
				CreatedAt: baseCreatedAt,
			},
			{
				BodyPart:  "upper",
				CreatedAt: baseCreatedAt.Add(30 * time.Minute),
			},
			{
				BodyPart:  "lower",
				CreatedAt: baseCreatedAt.Add(10 * time.Minute),
			},
		},
	}}

	assert.Equal(t, "upper", getLatestPart(album))
}
