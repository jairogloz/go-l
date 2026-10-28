package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDate_JSON(t *testing.T) {
	d := Date{
		Day:   25,
		Month: 12,
		Year:  2024,
	}

	data, err := json.Marshal(d)
	require.NoError(t, err)

	var unmarshaled Date
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	assert.Equal(t, d, unmarshaled)
}

func TestTimeRange_JSON(t *testing.T) {
	tr := TimeRange{
		Start: TimeOfDay{Hour: 9, Minute: 30},
		End:   TimeOfDay{Hour: 11, Minute: 0},
	}

	data, err := json.Marshal(tr)
	require.NoError(t, err)

	var unmarshaled TimeRange
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	assert.Equal(t, tr, unmarshaled)
}
