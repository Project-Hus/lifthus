package dto

import "time"

// SessionUserInfo is SessionHandler's response
type SessionUserInfo struct {
	UID        uint64  `json:"uid"`
	Registered bool    `json:"registered"`
	Username   *string `json:"username"`
	Usercode   string  `json:"usercode"`
}

type HusSessionCheckBody struct {
	Sid             string `json:"sid"`
	Uid             string `json:"uid"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
	EmailVerified   bool   `json:"email_verified"`
	Name            string `json:"name"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	Birthdate       string `json:"birthdate"`
}

type QueryUserDto struct {
	ID              uint64     `json:"uid,omitempty"`
	Registered      bool       `json:"registered,omitempty"`
	RegisteredAt    *time.Time `json:"registered_at,omitempty"`
	Username        *string    `json:"username,omitempty"`
	Email           *string    `json:"email,omitempty"`
	EmailVerified   *bool      `json:"email_verified,omitempty"`
	Name            *string    `json:"name,omitempty"`
	GivenName       *string    `json:"given_name,omitempty"`
	FamilyName      *string    `json:"family_name,omitempty"`
	Birthdate       *time.Time `json:"birthdate,omitempty"`
	ProfileImageURL string     `json:"profile_image_url,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`

	Usercode string  `json:"usercode,omitempty"`
	Company  *string `json:"company,omitempty"`
	Location *string `json:"location,omitempty"`
	Contact  *string `json:"contact,omitempty"`
}

type UpdateUserInfoDto struct {
	Uid       uint64     `json:"uid,omitempty"`
	Username  *string    `json:"username,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`

	Company  *string `json:"company,omitempty"`
	Location *string `json:"location,omitempty"`
	Contact  *string `json:"contact,omitempty"`
}
