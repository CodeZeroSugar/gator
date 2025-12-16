package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	runCmd, exists := c.commandMap[cmd.name]
	if !exists {
		return fmt.Errorf("input command does not exist")
	}

	if err := runCmd(s, cmd); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	_, exists := c.commandMap[name]
	if exists {
		fmt.Println("This command already exists")
		return
	}
	c.commandMap[name] = f
}
