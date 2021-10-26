package direction_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"robot/internal/direction"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name     string
		src      direction.Direction
		expected string
	}{
		{
			name:     "should return string represenation of East",
			src:      direction.East,
			expected: "EAST",
		},
		{
			name:     "should return string represenation of West",
			src:      direction.West,
			expected: "WEST",
		},
		{
			name:     "should return string represenation of South",
			src:      direction.South,
			expected: "SOUTH",
		},
		{
			name:     "should return string represenation of North",
			src:      direction.North,
			expected: "NORTH",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.src.String()
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestRotateLeft(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name     string
		src      direction.Direction
		expected direction.Direction
	}{
		{
			name:     "should rotate left from East",
			src:      direction.East,
			expected: direction.North,
		},
		{
			name:     "should rotate left from West",
			src:      direction.West,
			expected: direction.South,
		},
		{
			name:     "should rotate left from South",
			src:      direction.South,
			expected: direction.East,
		},
		{
			name:     "should rotate left from North",
			src:      direction.North,
			expected: direction.West,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.src
			actual.RotateLeft()
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestRotateRight(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name     string
		src      direction.Direction
		expected direction.Direction
	}{
		{
			name:     "should rotate right from East",
			src:      direction.East,
			expected: direction.South,
		},
		{
			name:     "should rotate right from West",
			src:      direction.West,
			expected: direction.North,
		},
		{
			name:     "should rotate right from South",
			src:      direction.South,
			expected: direction.West,
		},
		{
			name:     "should rotate right from North",
			src:      direction.North,
			expected: direction.East,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.src
			actual.RotateRight()
			require.Equal(t, tt.expected, actual)
		})
	}
}
