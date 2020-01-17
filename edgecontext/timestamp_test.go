package edgecontext_test

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"testing/quick"
	"time"

	"github.com/reddit/baseplate.go/edgecontext"
)

type jsonTestType struct {
	Timestamp edgecontext.TimestampMillisecond `json:"timestamp,omitempty"`
}

func TestTimestampQuick(t *testing.T) {
	f := func(ms int64) bool {
		time := edgecontext.MillisecondsToTime(ms)
		actual := edgecontext.TimeToMilliseconds(time)
		if actual != ms {
			t.Errorf(
				"For timestamp %d we got %v and %d", ms, time, actual,
			)
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTimestampJSONQuick(t *testing.T) {
	f := func(ms int64) bool {
		var v1, v2 jsonTestType

		v1 = jsonTestType{
			Timestamp: edgecontext.TimestampMillisecond(edgecontext.MillisecondsToTime(ms)),
		}
		s, err := json.Marshal(v1)
		if err != nil {
			t.Error(err)
			return false
		}
		substr := strconv.FormatInt(ms, 10)
		if !strings.Contains(string(s), substr) {
			t.Errorf("Encoded json expected to contain %q, got %q", substr, s)
		}

		err = json.Unmarshal(s, &v2)
		if err != nil {
			t.Error(err)
			return false
		}
		if !v2.Timestamp.ToTime().Equal(v1.Timestamp.ToTime()) {
			t.Errorf("Expected timestamp %v, got %v", v1.Timestamp, v2.Timestamp)
		}
		return !t.Failed()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTimestampJSON(t *testing.T) {
	var v jsonTestType

	t.Run(
		"null",
		func(t *testing.T) {
			s, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}
			substr := "null"
			if !strings.Contains(string(s), substr) {
				t.Errorf("Encoded json expected to contain %q, got %q", substr, s)
			}
			expectedStr := `{"timestamp":null}`
			if string(s) != expectedStr {
				t.Errorf("Encoded json expected to be %q, got %q", expectedStr, s)
			}

			v.Timestamp = edgecontext.TimestampMillisecond(time.Now())
			err = json.Unmarshal(s, &v)
			if err != nil {
				t.Fatal(err)
			}
			if !v.Timestamp.ToTime().IsZero() {
				t.Errorf("Timestamp expected to be zero, got %v", v.Timestamp)
			}
		},
	)

	label := "2019-12-31T23:59:59.123Z"
	t.Run(
		label,
		func(t *testing.T) {
			const substr = "1577836799123"

			ts, err := time.Parse(time.RFC3339Nano, label)
			if err != nil {
				t.Fatal(err)
			}
			v.Timestamp = edgecontext.TimestampMillisecond(ts)
			s, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(s), substr) {
				t.Errorf("Encoded json expected to contain %q, got %q", substr, s)
			}
			expectedStr := `{"timestamp":` + substr + `}`
			if string(s) != expectedStr {
				t.Errorf("Encoded json expected to be %q, got %q", expectedStr, s)
			}

			v.Timestamp = edgecontext.TimestampMillisecond(time.Now())
			err = json.Unmarshal(s, &v)
			if err != nil {
				t.Fatal(err)
			}
			if !v.Timestamp.ToTime().Equal(ts) {
				t.Errorf("Timestamp expected %v, got %v", ts, v.Timestamp)
			}
		},
	)
}
