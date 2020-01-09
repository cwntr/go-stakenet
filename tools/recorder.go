package tools

import "github.com/dnaeon/go-vcr/recorder"

//CreateRecorder creates a HTTP request and response recorder used for external packages
func CreateRecorder(casettePath string) (*recorder.Recorder, error) {
	return recorder.New(casettePath)
}
