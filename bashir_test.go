package bashir

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TODO: use bashir to test elim!
func TestElim(t *testing.T) {
	df1 := "Mon 02 January 2006 03:04:05 PM MST"
	df2 := "2006-01-02"

	tests := map[string]struct {
		input []string
		want  []string
	}{
		"date local": {
			input: []string{"date"},
			want:  []string{df(time.Now(), df1)},
		},
		"date utc": {
			input: []string{"date -u"},
			want:  []string{df(time.Now().UTC(), df1)},
		},
		"date ymd": {input: []string{`date +"%Y-%m-%d"`},
			want: []string{df(time.Now(), df2)},
		},
		"date three": {
			input: []string{
				"date",
				"date -u",
				`date +"%Y-%m-%d"`,
			},
			want: []string{
				df(time.Now(), df1),
				df(time.Now().UTC(), df1),
				df(time.Now(), df2),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Bash(tc.input...)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func df(d time.Time, f string) string {
	return d.Format(f)
}
