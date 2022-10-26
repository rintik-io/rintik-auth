package app

import "flag"

var (
	ConfigPath string
	LogPath    string
	getVersion bool
)

func getFlags() {
	flag.StringVar(&ConfigPath, "config", "", "Set the Path of Configuration File (Optional)")
	flag.StringVar(&LogPath, "log", "", "Set the Path & Filename of Log File (Optional)")
	flag.BoolVar(&getVersion, "version", false, "Show current app version")

	flag.Parse()
}
