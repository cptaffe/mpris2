package main

import "os"
import dbus "github.com/guelfey/go.dbus"
import "fmt"

// This program interacts with any Media Player implementing
// the MPRIS D-Bus Inteface Spcecification.

// currently only use this player
var mediaPlayer string = "spotify"

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

func sendCommand(player string, command string) error {
	// Connect to player and send command
	conn, err := dbus.SessionBus();
	if err != nil {
		return err;
	}
	defer conn.Close();

	// Emits signal on message bus
	conn.Emit("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2." + command)

	return nil
}

func main() {
	err := sendCommand(mediaPlayer, "Player.PlayPause")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
