package main
package main
const (
	LEAD = "LEAD"
)
// Lead : lead detials
type Lead struct {
	DocType  string              `json:"docTyp"`
	ID       string              `json:"lead_id"`
	UpdatedBy string `json:"updated_by"`
	UpdatedDate int64 `json:"updated_date"`
	Description string `json:"lead_description"`
	
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

	Company       string `json:"lead_company"`
	Industry      string `json:"lead_industry"`
	AnnualRevenue int64 `json:"lead_annual_revenue"`
	Website       string `json:"lead_website"`
	NoofEmployees uint   `json:"lead_noof_employees"`
	
	Source     string `json:"lead_source"`
	Status     string `json:"lead_status"`
	Rating     string `json:"lead_rating"`
	CreateBy   string `json:"created_by"`
	CreateDate int64 `json:"created_date"`
	ContactID  string `json:"contact_id"`
	Owner      string `json:"lead_owner"`
}


