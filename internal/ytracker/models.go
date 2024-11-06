package ytracker

type User struct {
	Self                 string   `json:"self"`
	Uid                  int64    `json:"uid"`
	Login                string   `json:"login"`
	TrackerUid           int64    `json:"trackerUid"`
	PassportUid          int      `json:"passportUid"`
	CloudUid             string   `json:"cloudUid"`
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	Display              string   `json:"display"`
	Email                string   `json:"email"`
	External             bool     `json:"external"`
	HasLicense           bool     `json:"hasLicense"`
	Dismissed            bool     `json:"dismissed"`
	UseNewFilters        bool     `json:"useNewFilters"`
	DisableNotifications bool     `json:"disableNotifications"`
	FirstLoginDate       string   `json:"firstLoginDate"`
	LastLoginDate        string   `json:"lastLoginDate"`
	WelcomeMailSent      bool     `json:"welcomeMailSent"`
	Sources              []string `json:"sources"`
}

type Issue struct {
	Self                 string `json:"self"`
	Id                   string `json:"id"`
	Key                  string `json:"key"`
	Version              int    `json:"version"`
	LastCommentUpdatedAt string `json:"lastCommentUpdatedAt"`
	Summary              string `json:"summary"`
	Complexity           string `json:"complexity"`
	Spent                string `json:"spent"`
	StatusStartTime      string `json:"statusStartTime"`
	UpdatedBy            struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"updatedBy"`
	StatusType struct {
		Value string `json:"value"`
		Order int    `json:"order"`
	} `json:"statusType"`
	Description string `json:"description"`
	Boards      []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"boards"`
	Type struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"type"`
	Priority struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"priority"`
	PreviousStatusLastAssignee struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"previousStatusLastAssignee"`
	CreatedAt string `json:"createdAt"`
	Followers []struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"followers"`
	CreatedBy struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"createdBy"`
	CommentWithoutExternalMessageCount int `json:"commentWithoutExternalMessageCount"`
	Votes                              int `json:"votes"`
	CommentWithExternalMessageCount    int `json:"commentWithExternalMessageCount"`
	Assignee                           struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"assignee"`
	Queue struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"queue"`
	UpdatedAt string `json:"updatedAt"`
	Status    struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"status"`
	PreviousStatus struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"previousStatus"`
	Favorite bool `json:"favorite"`
}

type IssueLog struct {
	Id    string `json:"id"`
	Self  string `json:"self"`
	Issue struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"issue"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy struct {
		Self        string `json:"self"`
		Id          int64  `json:"id,string"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"updatedBy"`
	Type      string `json:"type"`
	Transport string `json:"transport"`
	Fields    []struct {
		Field struct {
			Self    string `json:"self"`
			Id      string `json:"id"`
			Display string `json:"display"`
		} `json:"field"`
		From interface{} `json:"from"`
		To   interface{} `json:"to"`
	} `json:"fields,omitempty"`
	Links []struct {
		From interface{} `json:"from"`
		To   struct {
			Type struct {
				Self    string `json:"self"`
				Id      string `json:"id"`
				Inward  string `json:"inward"`
				Outward string `json:"outward"`
			} `json:"type"`
			Direction string `json:"direction"`
			Object    struct {
				Self    string `json:"self"`
				Id      string `json:"id"`
				Key     string `json:"key"`
				Display string `json:"display"`
			} `json:"object"`
		} `json:"to"`
	} `json:"links,omitempty"`
	Attachments struct {
		Added []struct {
			Self    string `json:"self"`
			Id      string `json:"id"`
			Display string `json:"display"`
		} `json:"added"`
	} `json:"attachments,omitempty"`
}
