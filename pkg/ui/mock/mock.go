package mock

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mgutz/ansi"
)

// ProgressBar displays a simple text spinner-style
// progress indicator with a label. Fakes long-running operations
// to increase realism
func ProgressBar(label string, duration int, finalmsg bool) {
	// FIXME: (JD) This is totally unrealistic and inadequate
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	if finalmsg {
		s.FinalMSG = label + "… Done.\n" // Newline is important!
	}
	s.Prefix = label + "… " // Leave a little space after the label
	s.Start()
	time.Sleep(time.Duration(duration) * time.Millisecond)
	s.Stop()
}

// Success prints "Success!" plus optional message
func Success(params ...string) {
	var msg = ""
	if len(params) > 0 {
		msg = params[0]
		fmt.Println(ansi.Green + "▸ Success! " + ansi.Reset + msg)
	}
}

// Failure tries to summarise what went wrong
// and provide guidance for recovery
func Failure(params ...string) {
	var msg = ""
	if len(params) > 0 {
		msg = params[0]
		fmt.Println(ansi.Red + "▸ Operation failed! " + ansi.Reset + msg)
	}
}

// Notice displays the result of a notice() function call
func Notice(message string) {
	fmt.Println(ansi.LightCyan + "[notice]\t" + ansi.Reset + message)
}

// DryRunBanner is some visual sugar to indicate simulation operations
func DryRunBanner() {
	fmt.Println(ansi.Blue + "==== DRY RUN MODE ====" + ansi.Reset)
}

// DescribeStep is the first thing that gets spit out after
// calling a command; this is sort of like the instruction text
// in a dialog box
func DescribeStep(msg string) {
	fmt.Println("\n" + ansi.Cyan + "▸ " + msg + ansi.Reset + "\n")
}

// DescribeStepWithField is the first thing that gets spit out after
// calling a command; this is sort of like the instruction text
// in a dialog box
func DescribeStepWithField(field string, msg string) {
	fmt.Println("\n" + ansi.Cyan + "▸ " + field + " " + ansi.Reset + msg + "\n")
}

// DiffAdd show that a thing will be added
func DiffAdd(message string) {
	fmt.Println(ansi.Green + "+ " + message + ansi.Reset)
}

// DiffRemove show that a thing will be removed
func DiffRemove(message string) {
	fmt.Println(ansi.Magenta + "- " + message + ansi.Reset)
}

// DiffChange show that a thing will be changed
func DiffChange(message string) {
	fmt.Println(ansi.Yellow + "~ " + message + ansi.Reset)
}

// DiffUnchanged shows that a thing will be not be changed
func DiffUnchanged(message string) {
	fmt.Println(ansi.LightBlack + "  " + message + ansi.Reset)
}

// DiffConflict shows a resource property where what's in the manifest/stored state doesn't match the running state
func DiffConflict(message string) {
	fmt.Println(ansi.LightRed + "! " + message + ansi.Reset)
}

// ExecOut passes through a line of stdout from an ssh-like thing
func ExecOut(target string, task string, message string) {
	fmt.Println(ansi.Green + "[" + target + "](" + task + "): " + ansi.Reset + message)
}
