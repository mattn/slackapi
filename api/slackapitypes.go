package api

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type ChannelListResponse struct {
	Response
	Channels []Channel `json:"channels"`
}
type UserListResponse struct {
	Response
	Users []User `json:"users"`
}

type RtmStartResponse struct {
	Response
	Url string `json:"url"`
	//Self     User      `json:"self"`
	Users    []User    `json:"users"`
	Channels []Channel `json:"channels"`

	// Groups
	// Ims
	// Team
	// Bots
}

type Channel struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	IsChannel  bool     `json:"is_channel"`
	Created    int      `json:"created"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	IsGeneral  bool     `json:"is_general"`
	Members    []string `json:"members"`
	Topic      struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"topic"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"purpose"`
	IsMember    bool   `json:"is_member"`
	LastRead    string `json:"last_read"`
	UnreadCount int    `json:"unread_count"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
	Color   string `json:"color"`
	Profile struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RealName  string `json:"real_name"`
		Email     string `json:"email"`
		Skype     string `json:"skype"`
		Phone     string `json:"phone"`
		Image24   string `json:"image_24"`
		Image32   string `json:"image_32"`
		Image48   string `json:"image_48"`
		Image72   string `json:"image_72"`
		Image192  string `json:"image_192"`
	} `json:"profile"`
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	HasFiles          bool `json:"has_files"`
}
