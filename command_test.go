package main

import (
    "fmt"
    "reflect"
    "testing"
)

type commandTestCase struct {
    name     string
    command  string
    handler  func(*state, command) error
    wantErr  error
}

func setupCommandTest() commands {
    return commands{registeredCommands: make(map[string]func(*state, command) error)}
}

func TestRegisterCommand(t *testing.T) {
    tests := []commandTestCase{
        {
            name:    "register successful command",
            command: "test",
            handler: func(*state, command) error {
                return nil
            },
            wantErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cmds := setupCommandTest()
            cmds.register(tt.command, tt.handler)

            registerFunc, ok := cmds.registeredCommands[tt.command]
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
            handler: func(*state, command) error {
                return testError
            },
            wantErr: testError,
        },
        {
            name:    "run command successfully",
            command: "success",
            handler: func(*state, command) error {
                return nil
            },
            wantErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cmds := setupCommandTest()
            cmds.register(tt.command, tt.handler)

            gotErr := cmds.run(&state{}, command{Name: tt.command})
            if gotErr != tt.wantErr {
                t.Errorf("Want error %v, got %v", tt.wantErr, gotErr)
            }
        })
    }
}