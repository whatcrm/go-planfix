package models

type Contact struct {
	ID                       int              `json:"id,omitempty"`
	Template                 *ContactTemplate `json:"template,omitempty"`
	SourceObjectId           string           `json:"sourceObjectId,omitempty"`
	SourceDataVersion        string           `json:"sourceDataVersion,omitempty"`
	Name                     string           `json:"name,omitempty"`
	Midname                  string           `json:"midname,omitempty"`
	Lastname                 string           `json:"lastname,omitempty"`
	Gender                   string           `json:"gender,omitempty"`
	Description              string           `json:"description,omitempty"`
	Address                  string           `json:"address,omitempty"`
	Site                     string           `json:"site,omitempty"`
	Email                    string           `json:"email,omitempty"`
	AdditionalEmailAddresses []string         `json:"additionalEmailAddresses,omitempty"`
	Skype                    string           `json:"skype,omitempty"`
	TelegramId               string           `json:"telegramId,omitempty"`
	Telegram                 string           `json:"telegram,omitempty"`
	Facebook                 string           `json:"facebook,omitempty"`
	Instagram                string           `json:"instagram,omitempty"`
	Vk                       string           `json:"vk,omitempty"`
	Position                 string           `json:"position,omitempty"`
	Group                    *ContactGroup    `json:"group,omitempty"`
	IsCompany                bool             `json:"isCompany,omitempty"`
	IsDeleted                bool             `json:"isDeleted,omitempty"`
	BirthDate                *TimePoint       `json:"birthDate,omitempty"`
	CreatedDate              *TimePoint       `json:"createdDate,omitempty"`
	DateOfLastUpdate         *TimePoint       `json:"dateOfLastUpdate,omitempty"`
	Supervisors              interface{}      `json:"supervisors,omitempty"`
	Phones                   []ContactPhone   `json:"phones,omitempty"`
	Companies                []interface{}    `json:"companies,omitempty"`
	Contacts                 []interface{}    `json:"contacts,omitempty"`
	Files                    []interface{}    `json:"files,omitempty"`
	DataTags                 []interface{}    `json:"dataTags,omitempty"`
	CustomFieldData          []interface{}    `json:"customFieldData,omitempty"`
}

type ContactListRequest struct {
	CommonListParams
}

type ContactListResponse struct {
	SourceId string    `json:"sourceId"`
	Result   string    `json:"result"`
	Contacts []Contact `json:"contacts"`
}

type ContactTemplate struct {
	Id int `json:"id"`
}

type ContactGroup struct {
	Id int `json:"id"`
}

type ContactBirthDate struct {
	Date string `json:"date"`
}

type ContactPhone struct {
	Number       string `json:"number"`
	MaskedNumber string `json:"maskedNumber,omitempty"`
	Type         int    `json:"type"`
}

type ContactRequest struct {
	ID          int             `json:"id"`
	Template    ContactTemplate `json:"template,omitempty"`
	Name        string          `json:"name,omitempty"`
	Midname     string          `json:"midname,omitempty"`
	Lastname    string          `json:"lastname,omitempty"`
	Gender      string          `json:"gender,omitempty"`
	Description string          `json:"description,omitempty"`
	Address     string          `json:"address,omitempty"`
	Site        string          `json:"site,omitempty"`
	Email       string          `json:"email,omitempty"`
	Skype       string          `json:"skype,omitempty"`
	TelegramId  string          `json:"telegramId"`
	Telegram    string          `json:"telegram"`
	Facebook    string          `json:"facebook"`
	Vk          string          `json:"vk"`
	Instagram   string          `json:"instagram"`
	IsCompany   bool            `json:"isCompany"`
	IsDeleted   bool            `json:"isDeleted"`
	Position    string          `json:"position"`
	CreatedDate struct {
		Date               string `json:"date"`
		Time               string `json:"time"`
		Datetime           string `json:"datetime"`
		DateTimeUtcSeconds string `json:"dateTimeUtcSeconds"`
	} `json:"createdDate"`
	DateOfLastUpdate struct {
		Date               string `json:"date"`
		Time               string `json:"time"`
		Datetime           string `json:"datetime"`
		DateTimeUtcSeconds string `json:"dateTimeUtcSeconds"`
	} `json:"dateOfLastUpdate"`
	Phones    []interface{} `json:"phones"`
	Companies []interface{} `json:"companies"`
	Contacts  []interface{} `json:"contacts"`
	Files     []interface{} `json:"files"`
	DataTags  []interface{} `json:"dataTags"`
}

type ContactGetResponse struct {
	Result  string  `json:"result"`
	Contact Contact `json:"contact"`
}
