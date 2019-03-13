package ui

import (
	"testing"
)

// TODO: Make these real tests.

func Test_Message(t *testing.T) {
	Message("resource", "I am a resource message")
}

func Test_ShowMessage(t *testing.T) {
	ShowMessage("Hiya, I'm a message")
}

func Test_AskForConfirmation(t *testing.T) {
	AskForConfirmation("Are you sure?")
}

func Test_Delay(t *testing.T) {
	Delay(1000)
}
