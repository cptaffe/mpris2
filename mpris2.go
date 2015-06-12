package mpris2

import dbus "github.com/guelfey/go.dbus"

type Player struct {
	name string // name of the player, e.g "spotify"
}

func NewPlayer(name string) *Player {
	player := new(Player)
	player.name = "org.mpris.MediaPlayer2." + name
	return player
}

func (player *Player) object() (obj *dbus.Object, err error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}

	obj = conn.Object(player.name, "/org/mpris/MediaPlayer2")
	return
}

func (player *Player) Next() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Next", 0)
	return
}

func (player *Player) Previous() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Previous", 0)
	return
}

func (player *Player) Pause() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Pause", 0)
	return
}

func (player *Player) PlayPause() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
	return
}

func (player *Player) Stop() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
	return
}

func (player *Player) Play() (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Play", 0)
	return
}

// TODO: fix this function
func (player *Player) Seek(offset uint32) (err error) {
	obj, err := player.object()
	if err != nil {
		return
	}

	obj.Call("org.mpris.MediaPlayer2.Player.Seek", 0, offset)
	return
}
