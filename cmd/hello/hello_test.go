package hello

import (
	"bytes"
	"testing"

	"github.com/urfave/cli"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantErr    bool
		wantOutput string
	}{
		{
			name:       "print help",
			args:       nil,
			wantErr:    true,
			wantOutput: "",
		},
		{
			name:       "print un-quoted",
			args:       []string{"world"},
			wantErr:    false,
			wantOutput: "Hello world.\n",
		},
		{
			name:       "print quoted",
			args:       []string{"--quote-name", "world"},
			wantErr:    false,
			wantOutput: "Hello \"world\".\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var output bytes.Buffer
			testApp := cli.NewApp()
			testApp.Flags = Command.Flags
			testApp.Action = Command.Action
			testApp.Writer = &output

			if err := testApp.Run(append([]string{"test"}, tt.args...)); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantOutput != "" && output.String() != tt.wantOutput {
				t.Errorf("Run() output = %s, wantOutput %s", output.String(), tt.wantOutput)
			}
		})
	}
}
