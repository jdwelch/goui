package mock

import (
	"time"

	"github.com/briandowns/spinner"
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
