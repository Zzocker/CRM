# CRM (Customer Relationship Management Software)

>## Table of Contents

- [Installation](#installation)
- [Status](#status)
- [Documentation](#documentation)
- [Tech Stack](#Stack)
- [Contributors](#Contributors)

## Installation

### Clone
- Clone this repo to your local machine using `https://github.com/Zzocker/CRM.git`

### Install docker (Optional)
- chmod +x ./docker.sh
- ./docker.sh
- usermod -a -G docker ${USER}

### Fabric network

#### Configuration of fabric network (v2.0)
> 2 Orderer using **raft** consensus

> 1 org named **Peepaltree** with one peer name devpeer. devpeer is using **Couchdb** as worldstate.

> Chaincode written in **Golang**

#### Start Docker Containers and etup the peers

    > cd network/deployment/
    > docker-compose -f docker-compose-cli.yaml -f docker-compose-orderer.yaml -f docker-compose-peerca.yaml up -d
    > docker exec -it cli bash
    > cd channel-artifacts/
    > chmod +x chaincodetest.sh
    > ./chaincodetest.sh

#####  Install chaincode on peer

    > export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
    > peer lifecycle chaincode package crm.tar.gz --label crm -p 
    > peer lifecycle chaincode install crm.tar.gz
    > export CC_PACKAGE=$(peer lifecycle chaincode queryinstalled -O json | jq .installed_chaincodes | jq .[0] | jq .package_id)
    > peer lifecycle chaincode approveformyorg -C crmchannel -n crm --package-id $CC_PACKAGE -v 1.0 -o orderer:7050 --tls --cafile $ORDERER
    > peer lifecycle chaincode commit -C crmchannel -n crm -v 1.0 -o orderer:7050 --tls --cafile $ORDERER

### Fire up the APIs
    
    > cd api
    > npm i
    > npm run dev

## Status

:white_large_square: Chaincode development
- :heavy_check_mark: [Lead model](#leadmodel)
- :heavy_check_mark: [Deal model](#dealmodel)
- :heavy_check_mark: Lead and Deal Function 
- :white_large_square: Access control for Admin, salesperson,manger on deals and leads 

:heavy_check_mark: Network configured

:heavy_check_mark: Middleware (nodejs server)

- :heavy_check_mark: [Lead APIs](#APIDocs)
- :heavy_check_mark: [Deal APIs](#APIDocs)
- :heavy_check_mark: Containerization of Node server

:white_large_square: Deployment on AWS
- :heavy_check_mark: Tested
- :white_large_square: Dev running

## Documentation

### leadmodel
```golang
type Lead struct {
	DocType     string `json:"docTyp"`
	ID          string `json:"lead_id"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedDate int64  `json:"updated_date"`
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
	AnnualRevenue int64  `json:"lead_annual_revenue"`
	Website       string `json:"lead_website"`
	NoofEmployees uint   `json:"lead_noof_employees"`

	Source     string `json:"lead_source"`
	Status     string `json:"lead_status"`
	Rating     string `json:"lead_rating"`
	CreateBy   string `json:"created_by"`
	CreateDate int64  `json:"created_date"`
	ContactID  string `json:"contact_id"`
	Owner      string `json:"lead_owner"`
}
```
### dealmodel

```golang
    type Deal struct {
	DocType string `json:"docTyp"`

	DealID         string `json:"deal_id"`
	OrganizationID string `json:"organization_id"`
	DealOwner      string `json:"deal_owner"`
	DealLeadID     string `json:"deal_lead_id"`
	Description    string `json:"deal_description"`
	DealAccountID  string `json:"deal_aacount_id"`

	DealName          string `json:"deal_name"`
	DealDate          string `json:"deal_date"`
	DealPotentialName string `json:"deal_potential_name"`

	DealAccountName string `json:"deal_account_name"`

	DealType string `json:"deal_type"`

	DealNextStep     string `json:"deal_next_step"`
	DealNextStepDate string `json:"deal_next_stepdate"`

	DealLeadSource string `json:"deal_lead_source"`

	DealCampaignSource  string  `json:"deal_campaign_source"`
	DealContactName     string  `json:"deal_contact_name"`
	DealCurrencyCode    string  `json:"deal_currency_code"`
	DealAmount          float64 `json:"deal_amount"`
	DealClosingDate     string  `json:"deal_closing_date"`
	DealStage           string  `json:"deal_stage"`
	DealProbility       string  `json:"deal_probility"`
	DealExpectedRevenue string  `json:"deal_expected_revenue"`

	CreatedBy   string `json:"created_by"`
	CreatedDate int64  `json:"created_date"`

	UpdatedBy   string `json:"updated_by"`
	UpdatedDate int64  `json:"updated_date"`
}
```

### APIDocs

Postman Documentation <a href="https://documenter.getpostman.com/view/7262970/SzYT6hqk" target="_blank">`Link`</a>

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/f7a051370809fb91e7a8)

## Stack

---
## Contributors

 <a href="https://github.com/Zzocker" target="_blank">**Pritam Singh**</a> 

[![Pritam Singh](https://avatars1.githubusercontent.com/u/43764373?s=200&u=6a3ef280e24c5ffe3b5e108338e028ca4e0745e4&v=4)](https://www.linkedin.com/in/pritam-singh-b1807617b/)   

---


