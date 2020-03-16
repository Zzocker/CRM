package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) CreateNewLead(ctx CustomTransactionContextInterface,id,salutation,fname,lname,mobile,requester string) Lead  {
	id = uuid.New().String()
	createTime:= time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID: id,
		Update: make(map[int64]string),
		General: LeadGeneralDetails{},
		Company: LeadCompanyDetails{},
		Personal: LeadPersonalDetails{},
	}

	lead.Update[createTime]=requester

	lead.Personal.Saluation=salutation
	lead.Personal.FirstName =fname
	lead.Personal.Mobile=mobile

	lead.General.CreateBy=requester
	lead.General.CreateDate=createTime
	lead.General.Owner=requester

	leadAsByte,_:=json.Marshal(lead)

	ctx.GetStub().PutState(id,leadAsByte)


	// return id
	return lead
}
func (c *Chaincode) CreateLeadFromContact(ctx CustomTransactionContextInterface,id,contactid ,title,salutation,fname,lname,mobile,country,state,city,pincode,email,requester string) Lead {
	id = uuid.New().String()
	createTime:= time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID: id,
		Update: make(map[int64]string),
		General: LeadGeneralDetails{},
		Company: LeadCompanyDetails{},
		Personal: LeadPersonalDetails{},
	}
	lead.Update[createTime]=requester

	lead.Personal.JobTitle= title
	lead.Personal.Saluation=salutation
	lead.Personal.FirstName=fname
	lead.Personal.LastName=lname
	lead.Personal.Mobile=mobile
	lead.Personal.Country=country
	lead.Personal.State=state
	lead.Personal.City=city
	lead.Personal.Zipcode=pincode
	lead.Personal.Email=email


	lead.General.CreateBy=requester
	lead.General.CreateDate=createTime
	lead.General.Owner=requester
	lead.General.ContactID=contactid

	leadAsByte,_:=json.Marshal(lead)

	ctx.GetStub().PutState(id,leadAsByte)


	// return id
	return lead

}
