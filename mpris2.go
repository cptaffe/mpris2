package mpris2

import dbus "github.com/guelfey/go.dbus"

type Player struct {
	Name string // name of the player
	obj *dbus.Object // associated dbus object
}

func NewPlayer(name string) (*Player, error) {
	player := new(Player)
	player.Name = "org.mpris.MediaPlayer2." + name

	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	player.obj = conn.Object(player.Name, "/org/mpris/MediaPlayer2")

	return player, err
}

// MediaPlayer2 Methods

// Brings the media player's user interface to the front
// using any appropriate mechanism available
func (player *Player) Raise() {
	player.obj.Call("org.mpris.MediaPlayer2.Raise", 0)
}

// Causes the media player to stop running
func (player *Player) Quit() {
	player.obj.Call("org.mpris.MediaPlayer2.Quit", 0)
}

// MediaPlayer2.Player Methods

// Skips to the next track in the tracklist
func (player *Player) Next() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Next", 0)
}

// Skips to the previous track in the tracklist
func (player *Player) Previous() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Previous", 0)
}

// Pauses playback
func (player *Player) Pause() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Pause", 0)
}

// Pauses playback
// If playback is already paused, resumes playback
func (player *Player) PlayPause() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
}

// Stops playback
func (player *Player) Stop() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
}

// Starts or resumes playback
func (player *Player) Play() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Play", 0)
}

// Seeks forward in the current track by the specified
// number of microseconds
func (player *Player) Seek(offset uint32) {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Seek", 0, offset)
}
