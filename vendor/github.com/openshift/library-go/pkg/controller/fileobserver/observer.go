package fileobserver

import (
	"os"

	"github.com/golang/glog"
)

type Observer interface {
	Run(stopChan <-chan struct{})
	AddReactor(reaction reactorFn, files ...string) Observer
}

// ActionType define a type of action observed on the file
type ActionType int

const (
	// FileModified means the file content was modified.
	FileModified ActionType = iota

	// FileCreated means the file was just created.
	FileCreated

	// FileDeleted means the file was deleted.
	FileDeleted
)

// reactorFn define a reaction function called when an observed file is modified.
type reactorFn func(file string, action ActionType) error

// ExitOnChangeReactor provides reactor function that causes the process to exit when the change is detected.
var ExitOnChangeReactor reactorFn = func(filename string, action ActionType) error {
	glog.Infof("exiting because %q changed", filename)
	os.Exit(0)
	return nil
}

// NewObserver provides a platform specific file observer controller capable of observing changes to reactors on disk and reacting to those
// changes.
var NewObserver func() (Observer, error)
