package model

type OuputData struct {
	Event            string               `json:"event"`
	Event_Type       string               `json:"event_type"`
	App_Id           string               `json:"app_id"`
	User_Id          string               `json:"user_id"`
	Message_Id       string               `json:"message_id"`
	Page_Title       string               `json:"page_title"`
	Page_Url         string               `json:"page_url"`
	Browser_Language string               `json:"browser_language"`
	Screen_Size      string               `json:"screen_size"`
	Attributes       map[string]Attribute `json:"attributes"`
	Traits           map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
