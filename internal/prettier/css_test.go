package prettier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSSPrettier_Pretty(t *testing.T) {
	tests := []struct {
		css     string
		want    string
		wantErr bool
	}{
		{
			css: `body{background-color:#fff;color:#000;}`,
			want: `body {
  background-color: #fff;
  color: #000;
}
`,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			p := NewCSSPrettier()
			got, err := p.Pretty(tt.css)

			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
