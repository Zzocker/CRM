package main

const (
	DEAL = "DEAL"
)

// Deal type
type Deal struct {
	DocType string `json:"docTyp"`

	DealID         string `json:"deal_id"`
	OrganizationID int `json:"organization_id"`
	DealOwner      string `json:"deal_owner"`
	DealLeadID     string `json:"deal_lead_id"`
	Description    string `json:"deal_description"`
	DealAccountID  string `json:"deal_account_id"`

	DealName          string `json:"deal_name"` // missing 
	DealDate          string `json:"deal_date"`
	DealPotentialName string `json:"deal_potential_name"`

	DealAccountName string `json:"deal_account_name"`

	DealType string `json:"deal_type"`

	DealNextStep     string `json:"deal_next_step"`
	DealNextStepDate string `json:"deal_next_stepdate"`

	DealLeadSource string `json:"deal_lead_source"`

	DealCampaignSource  string  `json:"deal_campaign_source"` // missing
	DealContactName     string  `json:"deal_contact_name"`
	DealCurrencyCode    string  `json:"deal_currency_code"`
	DealAmount          float64 `json:"deal_amount"`
	DealClosingDate     string  `json:"deal_closing_date"`
	DealStage           string  `json:"deal_stage"`
	DealProbility       string  `json:"deal_probability"`
	DealExpectedRevenue string  `json:"deal_expected_revence"`

	CreatedBy   string `json:"created_by"`
	CreatedDate int64  `json:"created_date"`

	UpdatedBy   string `json:"updated_by"`
	UpdatedDate int64  `json:"updated_date"`
}
