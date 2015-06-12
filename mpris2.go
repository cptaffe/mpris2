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

func (player *Player) Next() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Next", 0)
}

func (player *Player) Previous() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Previous", 0)
}

func (player *Player) Pause() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Pause", 0)
}

func (player *Player) PlayPause() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
}

func (player *Player) Stop() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Stop", 0)
}

func (player *Player) Play() {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Play", 0)
}

// TODO: fix this function
func (player *Player) Seek(offset uint32) {
	player.obj.Call("org.mpris.MediaPlayer2.Player.Seek", 0, offset)
}
