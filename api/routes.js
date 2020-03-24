const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.post('/createNewLead',(req,res)=>{
    rbody = req.body
    lead_lastname=rbody.lead_lastname
    company=rbody.company
    requester=rbody.requester
    network.contract("invoke",["CreateNewLead",lead_lastname,company,requester],(err,payload)=>{
        if (err){
            res.status(500).json({error})
        }else{
            id = payload.toString()
        res.status(200).send(id)
        }
    })
   
})

routes.post('/createNewLeadFromContact',(req,res)=>{
   
        const rbody = req.body
        contactid=rbody.contact_id
        first_name=rbody.lead_firstname
        last_name=rbody.lead_lastname
        mobile = rbody.lead_mobile
        country=rbody.lead_country
        state = rbody.lead_state
        city = rbody.lead_city
        pincode = rbody.lead_pincode
        email = rbody.lead_email
        requester = rbody.requester

        network.contract("invoke",["CreateLeadFromContact",contactid,first_name,last_name,mobile,country,state,city,pincode,email,requester],(err,payload)=>{
            if (err){
                res.status(500).json(err)
            }else{
                id = payload.toString()
                res.status(200).send(id)
            }
        })
        
})

routes.delete('/deletelead',(req,res)=>{
    network.contract("invoke",["DeleteLead",req.headers.id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }
        else {res.status(200).json(`DONE`)}
    })
})


routes.put('/updatename',(req,res)=>{
    rbody = req.body
    id=rbody.lead_id
    sal=rbody.lead_saluation 
    fname=rbody.lead_firstname 
    lname=rbody.lead_last_name 
    requester=rbody.requester
    network.contract("invoke",["UpdateName",id,sal,fname,lname,requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updatejobtitle',(req,res)=>{
    network.contract("invoke",["UpdateJobTitle",req.headers.id,req.body.leadJobtitle,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatephone',(req,res)=>{
    network.contract("invoke",["UpdatePhone",req.headers.id,req.body.lead_phone,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updateemail',(req,res)=>{
    network.contract("invoke",["UpdateEmail",req.headers.id,req.body.lead_email,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatemobile',(req,res)=>{
    network.contract("invoke",["UpdateMobile",req.headers.id,req.body.lead_mobile,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatesecondaryemail',(req,res)=>{
    network.contract("invoke",["UpdateSecondaryEmail",req.headers.id,req.body.lead_secondary_email,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
// from here
routes.put('/updateskypid',(req,res)=>{
    network.contract("invoke",["UpdateSkypid",req.headers.id,req.body.lead_skypid,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updateaddress',(req,res)=>{
    rbody= req.body
    street = rbody.lead_street
    areal=rbody.lead_arealocation
    city=rbody.lead_city
    state=rbody.lead_state
    zipcode = rbody.lead_zipcode
    country = rbody.lead_country
    network.contract("invoke",["UpdateAddress",req.headers.id,street,areal,city,state,zipcode,country,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
})

routes.put('/updateemailoptout',(req,res)=>{
    network.contract("invoke",["UpdateEmailoptout",req.headers.id,req.body.lead_emailopt_out,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatefax',(req,res)=>{
    network.contract("invoke",["UpdateFax",req.headers.id,req.body.lead_fax,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updatecompany',(req,res)=>{
    network.contract("invoke",["UpdateComanpy",req.headers.id,req.body.lead_company,req.body.lead_industry,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updateannualrevenue',(req,res)=>{
    network.contract("invoke",["UpdateAnnualRevenue",req.headers.id,req.body.lead_annual_revenue,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updatewebsite',(req,res)=>{
    network.contract("invoke",["UpdateWebsite",req.headers.id,req.body.lead_website,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatenoofemp',(req,res)=>{
    network.contract("invoke",["UpdateNoofEmployees",req.headers.id,req.body.lead_noof_employees,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatesource',(req,res)=>{
    network.contract("invoke",["UpdateSource",req.headers.id,req.body.lead_source,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatestatus',(req,res)=>{
    network.contract("invoke",["UpdateStatus",req.headers.id,req.body.lead_status,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updaterating',(req,res)=>{
    network.contract("invoke",["UpdateRating",req.headers.id,req.body.lead_rating,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updatecontactid',(req,res)=>{
    network.contract("invoke",["UpdateContactID",req.headers.id,req.body.contact_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/updateowner',(req,res)=>{
    network.contract("invoke",["UpdateOwner",req.headers.id,req.body.lead_owner,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/updateleaddescription',(req,res)=>{
    network.contract("invoke",["EditLeadDescription",req.headers.id,req.body.lead_description,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})



routes.get('/getallmylead',(req,res)=>{
    network.contract("query",["GetAllMyLeads",req.headers.type,"}}",req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(`No lead`)
        }else{res.status(200).json(JSON.parse(payload))}
    })
})
routes.get('/getlead',(req,res)=>{
    
    network.contract("query",["GetMyLead",req.headers.id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.message)
        }
        else {res.status(200).json(JSON.parse(payload))}
    })
})  




module.exports = routes


    