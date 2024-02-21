package apis

import (
	"encoding/json"
	"fmt"
	"time"
)

// Duration extends time.Duration with additional methods for (de-)serialization to/from
// JSON, allowing it to be embedded in custom API objects.
type Duration time.Duration

var (
	_ fmt.Stringer     = (*Duration)(nil)
	_ json.Marshaler   = (*Duration)(nil)
	_ json.Unmarshaler = (*Duration)(nil)
)

// String implements the fmt.Stringer interface.
func (d Duration) String() string {
	return time.Duration(d).String()
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Duration) UnmarshalJSON(data []byte) error {
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	dur, err := time.ParseDuration(dataStr)
	if err != nil {
		return fmt.Errorf("failed to parse duration %q: %w", dataStr, err)
	}

	*d = Duration(dur)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Duration(d).String() + `"`), nil
}
