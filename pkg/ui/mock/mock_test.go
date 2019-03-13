package mock

import (
	"testing"
)

// TODO: Make these real tests.

func Test_ProgressBar(t *testing.T) {
	ProgressBar("message", 10, true)
}

func Test_Delay(t *testing.T) {
	Delay(1000)
}
