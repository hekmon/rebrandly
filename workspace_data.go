package rebrandly

import "time"

type Workspace struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Links     int    `json:"links"`
	Teammates int    `json:"teammates"`
	Domains   int    `json:"domains"`
	Owner     struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"owner"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Type         string    `json:"type"`
	Default      bool      `json:"default"`
	Subscription struct {
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
			Workspaces struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"workspaces"`
			Cycle struct {
				Clicks struct {
					Used     int `json:"used"`
					Included int `json:"included"`
				} `json:"clicks"`
				Reports struct {
					Used     int `json:"used"`
					Included int `json:"included"`
				} `json:"reports"`
				Teammates struct {
					Used     int `json:"used"`
					Included int `json:"included"`
				} `json:"teammates"`
			} `json:"cycle"`
			Tags struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"tags"`
			Scripts struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"scripts"`
			Apps struct {
				Used     int `json:"used"`
				Included int `json:"included"`
			} `json:"apps"`
		} `json:"limits"`
	} `json:"subscription"`
	Role        string    `json:"role"`
	Clicks      int       `json:"clicks"`
	Sessions    int       `json:"sessions,omitempty"`
	LastClickAt time.Time `json:"lastClickAt,omitempty"`
	Correlation struct {
		Status string `json:"status"`
	} `json:"correlation"`
}
