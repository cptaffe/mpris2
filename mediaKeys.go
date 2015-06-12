package main

import "os"
import "os/exec"
import "fmt"

var usageMsg string = `mediaKeys [play | prev | next | vol-up | vol-down | mute]

play     signals play/pause keypress
prev     signals previous track keypress
next     signals next track keypress
vol-up   signals volume up keypress
vol-down signals voludme down keypress
mute     signals mute keypress`

func usage() {
	fmt.Printf("%s\n", usageMsg)
	os.Exit(1)
}

func main() {
	keys := make(map[string]string)
	keys["play"] = "XF86AudioPlay"
	keys["prev"] = "XF86AudioPrev"
	keys["next"] = "XF86AudioNext"
	keys["vol-up"] = "XF86AudioRaiseVolume"
	keys["vol-down"] = "XF86AudioLowerVolume"
	keys["mute"] = "XF86AudioMute"

	if len(os.Args) != 2 {
		usage()
	} else {
		key := keys[os.Args[1]]
		if key == "" {
			usage()
		} else {
			cmd := exec.Command("xdotool", "key", key)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	}
}
