package main

import (
	"encoding/json"
	."fmt"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) CreateNewLead(ctx CustomTransactionContextInterface, id, salutation, fname, lname, mobile, requester string) string {
	id = uuid.New().String()
	createTime := time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID:      id,
	}

	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime

	lead.Saluation = salutation
	lead.FirstName = fname
	lead.LastName = lname
	lead.Mobile = mobile

	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id
}
func (c *Chaincode) CreateLeadFromContact(ctx CustomTransactionContextInterface, id, contactid, title, salutation, fname, lname, mobile, country, state, city, pincode, email, requester string) string {
	id = uuid.New().String()
	createTime := time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID:      id,
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime

	lead.JobTitle = title
	lead.Saluation = salutation
	lead.FirstName = fname
	lead.LastName = lname
	lead.Mobile = mobile
	lead.Country = country
	lead.State = state
	lead.City = city
	lead.Zipcode = pincode
	lead.Email = email

	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester
	lead.ContactID = contactid

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id

}

func (c *Chaincode) DeleteLead(ctx CustomTransactionContextInterface, id, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester {
		Errorf("Unauthorized to Delete this lead ")
	}
	return ctx.GetStub().DelState(id)
}

