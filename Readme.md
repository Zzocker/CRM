# CRM (Customer Relationship Management Software)

>## Table of Contents

- [Installation](#installation)
- [Status](#status)
- [Documentation](#documentation)

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
- :heavy_check_mark: Deal model
- :heavy_check_mark: Lead and Deal Function 
- :white_large_square: Access control for Admin, salesperson,manger on deals and leads 

:heavy_check_mark: Network configured

:white_large_square: Middleware (nodejs server)

- :heavy_check_mark: Lead APIs
- :white_large_square: Deal APIs
- :white_large_square: Containerization of Node server

:white_large_square: Deployment on AWS
- :white_large_square: Tested
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