package mpris2

import dbus "github.com/guelfey/go.dbus"

type Player struct {
	name string // name of the player, e.g "spotify"
}

func Player(name string) player *Player {
	player := make(Player)
	player.name = "org.mpris.MediaPlayer2." + name
}

func (Player *player) Play() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}

	obj := conn.Object(player.name, "/org/mpris/MediaPlayer2")
	obj.Call("org.mpris.MediaPlayer2.Player.Play")
}
