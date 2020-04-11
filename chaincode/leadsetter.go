package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) UpdateName(ctx CustomTransactionContextInterface, id, Saluation, fname, lname, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Saluation = Saluation
	lead.FirstName = fname
	lead.LastName = lname

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateJobTitle(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.JobTitle = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdatePhone(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Phone = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateMobile(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Mobile = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateEmail(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Email = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateSecondaryEmail(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.SecondaryEmail = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateSkypid(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Skypid = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateAddress(ctx CustomTransactionContextInterface, id, street, arealoction, city, state, zipcode, country, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Street = street
	lead.Arealoction = arealoction
	lead.City = city
	lead.State = state
	lead.Zipcode = zipcode
	lead.Country = country

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateEmailoptout(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Emailoptout = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateFax(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Fax = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateComanpy(ctx CustomTransactionContextInterface, id, cname, industry, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Company = cname
	lead.Industry = industry
	lead.AnnualRevenue = 0
	lead.Website = ""
	lead.NoofEmployees = 0

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateAnnualRevenue(ctx CustomTransactionContextInterface, id string, NewUp int64, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.AnnualRevenue = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateWebsite(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Website = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateNoofEmployees(ctx CustomTransactionContextInterface, id string, NewUp uint, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.NoofEmployees = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateSource(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Source = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateStatus(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Status = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateRating(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Rating = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateContactID(ctx CustomTransactionContextInterface, id string, NewUp int, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.ContactID = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateOwner(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Owner = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) EditLeadDescription(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()

	lead.Description = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, existing)
}
