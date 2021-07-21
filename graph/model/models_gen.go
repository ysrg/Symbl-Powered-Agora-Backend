// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Pstn struct {
	Number string `json:"number"`
	Dtmf   string `json:"dtmf"`
}

type Passphrase struct {
	Host         *string `json:"host"`
	View         string  `json:"view"`
	SymblToken   string  `json:"SymblToken"`
	SymblExpires float64 `json:"SymblExpires"`
}

type Session struct {
	Channel     string           `json:"channel"`
	Title       string           `json:"title"`
	IsHost      bool             `json:"isHost"`
	Secret      string           `json:"secret"`
	MainUser    *UserCredentials `json:"mainUser"`
	ScreenShare *UserCredentials `json:"screenShare"`
}

type ShareResponse struct {
	Passphrase *Passphrase `json:"passphrase"`
	Channel    string      `json:"channel"`
	Title      string      `json:"title"`
	Pstn       *Pstn       `json:"pstn"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCredentials struct {
	Rtc          string  `json:"rtc"`
	Symt         string  `json:"symt"`
	SymblTExpire float64 `json:"symblTExpire"`
	Rtm          *string `json:"rtm"`
	UID          int     `json:"uid"`
}
