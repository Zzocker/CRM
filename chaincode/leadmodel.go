package main
const (
	LEAD = "LEAD"
)
// Lead : lead detials
type Lead struct {
	DocType  string              `json:"docTyp"`
	ID       string              `json:"lead_id"`
	Update   map[int64]string    `json:"updates"`
	General  LeadGeneralDetails  `json:"lead_general_details"`
	Company  LeadCompanyDetails  `json:"lead_company_details"`
	Personal LeadPersonalDetails `json:"lead_personal_details"`
}

// LeadGeneralDetails : lead general detials
type LeadGeneralDetails struct {
	Source     string `json:"lead_source"`
	Status     string `json:"lead_status"`
	Rating     string `json:"lead_rating"`
	CreateBy   string `json:"created_by"`
	CreateDate string `json:"created_date"`
	ContactID  string `json:"contact_id"`
	Owner      string `json:"lead_owner"`
}

// LeadCompanyDetails : Company details of leads
type LeadCompanyDetails struct {
	Company       string `json:"lead_company"`
	Industry      string `json:"lead_industry"`
	AnnualRevenue int64 `json:"lead_annual_revenue"`
	Website       string `json:"lead_website"`
	NoofEmployees uint   `json:"lead_noof_employees"`
}

// LeadPersonalDetails : personal details of lead
type LeadPersonalDetails struct {
	Saluation      string `json:"lead_saluation"`
	FirstName      string `json:"lead_firstname"`
	LastName       string `json:"lead_last_name"`
	JobTitle       string `json:"leadJobtitle"`
	Phone          string `json:"lead_phone"`
	Mobile         string `json:"lead_mobile"`
	Email          string `json:"lead_email"`
	SecondaryEmail string `json:"lead_secondary_email"`
	Skypid         string `json:"lead_skypid"`
	Street         string `json:"lead_street"`
	Arealoction    string `json:"lead_arealocation"`
	City           string `json:"lead_city"`
	State          string `json:"lead_state"`
	Zipcode        string `json:"lead_zipcode"`
	Country        string `json:"lead_country"`
	Emailoptout    string `json:"lead_emailopt_out"`
	Fax            string `json:"lead_fax"`
}
