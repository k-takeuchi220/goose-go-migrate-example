package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	ProfileInfo []byte    `json:"profile_info"`
	CreatedAt   time.Time `json:"created_at"`

	Profile ProfileInfo
}

func (u *User) LoadProfile() error {
	if len(u.ProfileInfo) == 0 {
		return nil
	}
	return u.Profile.Decode(u.ProfileInfo)
}

func (u *User) SaveProfile() error {
	data, err := u.Profile.Encode()
	if err != nil {
		return err
	}
	u.ProfileInfo = data
	return nil
}

type ProfileInfo struct {
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
}

func (pi *ProfileInfo) Encode() ([]byte, error) {
	return json.Marshal(pi)
}

func (pi *ProfileInfo) Decode(data []byte) error {
	return json.Unmarshal(data, pi)
}
