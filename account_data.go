package rebrandly

import "time"

// Account represents all the metadata related to the api key account
type Account struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	FullName     string    `json:"fullName"`
	AvatarURL    string    `json:"avatarUrl"`
	MyCname      string    `json:"myCname"`
	Subscription struct {
		Category  string    `json:"category"`
		CreatedAt time.Time `json:"createdAt"`
		Billing   struct {
			Cycle struct {
				Price struct {
					Full int `json:"full"`
					Net  int `json:"net"`
					Vat  int `json:"vat"`
				} `json:"price"`
				ResetsAt time.Time `json:"resetsAt"`
			} `json:"cycle"`
			Extra struct {
				Cycle struct {
				} `json:"cycle"`
			} `json:"extra"`
		} `json:"billing"`
		Limits struct {
			Links struct {
				Used     int `json:"used"`
				Included int `json:"included"`
				Blocked  int `json:"blocked"`
			} `json:"links"`
			Domains struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"domains"`
			Teammates struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"teammates"`
			Workspaces struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"workspaces"`
			Teams struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"teams"`
			Cycle struct {
				Clicks struct {
					Used     int `json:"used"`
					Included int `json:"included"`
				} `json:"clicks"`
			} `json:"cycle"`
			Tags struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"tags"`
			Scripts struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"scripts"`
		} `json:"limits"`
		Features struct {
			Links struct {
				Scripts bool `json:"scripts"`
				Tags    bool `json:"tags"`
				Qrcode  bool `json:"qrcode"`
				Emoji   bool `json:"emoji"`
				Utm     bool `json:"utm"`
				Notes   bool `json:"notes"`
				Rules   bool `json:"rules"`
				Import  bool `json:"import"`
				Export  bool `json:"export"`
				Jobs    bool `json:"jobs"`
			} `json:"links"`
			Workspaces bool `json:"workspaces"`
			Teams      bool `json:"teams"`
			Teammates  bool `json:"teammates"`
			Domains    struct {
				Whitelabeled bool `json:"whitelabeled"`
			} `json:"domains"`
			Clicks struct {
				Total bool `json:"total"`
				Last  bool `json:"last"`
			} `json:"clicks"`
			Reports struct {
				Links  bool `json:"links"`
				Custom bool `json:"custom"`
				Public bool `json:"public"`
			} `json:"reports"`
			TwoFactorAuth bool `json:"twoFactorAuth"`
		} `json:"features"`
	} `json:"subscription"`
	Clicks      int       `json:"clicks"`
	Sessions    int       `json:"sessions"`
	LastClickAt time.Time `json:"lastClickAt"`
}
