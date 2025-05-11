package cli

import (
	"fmt"
	"reflect"
	"testing"
)

type commandTestCase struct {
	name    string
	command string
	handler func(*State, Command) error
	wantErr error
}

func setupCommandTest() Commands {
	return Commands{RegisteredCommands: make(map[string]func(*State, Command) error)}
}

func TestRegisterCommand(t *testing.T) {
	tests := []commandTestCase{
		{
			name:    "register successful command",
			command: "test",
			handler: func(*State, Command) error {
				return nil
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmds := setupCommandTest()
			cmds.Register(tt.command, tt.handler)

			registerFunc, ok := cmds.RegisteredCommands[tt.command]
			if !ok {
				t.Fatalf("Failed to register command %s", tt.command)
			}

			if reflect.ValueOf(tt.handler).Pointer() != reflect.ValueOf(registerFunc).Pointer() {
				t.Errorf("Want function %p, got %p", tt.handler, registerFunc)
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	testError := fmt.Errorf("test error")
	tests := []commandTestCase{
		{
			name:    "run command with error",
			command: "test",
			handler: func(*State, Command) error {
				return testError
			},
			wantErr: testError,
		},
		{
			name:    "run command successfully",
			command: "success",
			handler: func(*State, Command) error {
				return nil
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmds := setupCommandTest()
			cmds.Register(tt.command, tt.handler)

			gotErr := cmds.Run(&State{}, Command{Name: tt.command})
			if gotErr != tt.wantErr {
				t.Errorf("Want error %v, got %v", tt.wantErr, gotErr)
			}
		})
	}
}
