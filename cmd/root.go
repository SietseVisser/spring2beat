package cmd

import (
	"github.com/SietseVisser/spring2beat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "spring2beat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
