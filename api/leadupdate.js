const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.put('/name',(req,res)=>{
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

routes.put('/jobtitle',(req,res)=>{
    network.contract("invoke",["UpdateJobTitle",req.headers.id,req.body.leadJobtitle,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/phone',(req,res)=>{
    network.contract("invoke",["UpdatePhone",req.headers.id,req.body.lead_phone,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/email',(req,res)=>{
    network.contract("invoke",["UpdateEmail",req.headers.id,req.body.lead_email,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/mobile',(req,res)=>{
    network.contract("invoke",["UpdateMobile",req.headers.id,req.body.lead_mobile,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/secondaryemail',(req,res)=>{
    network.contract("invoke",["UpdateSecondaryEmail",req.headers.id,req.body.lead_secondary_email,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
// from here
routes.put('/skypid',(req,res)=>{
    network.contract("invoke",["UpdateSkypid",req.headers.id,req.body.lead_skypid,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/address',(req,res)=>{
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

routes.put('/emailoptout',(req,res)=>{
    network.contract("invoke",["UpdateEmailoptout",req.headers.id,req.body.lead_emailopt_out,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/fax',(req,res)=>{
    network.contract("invoke",["UpdateFax",req.headers.id,req.body.lead_fax,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/company',(req,res)=>{
    network.contract("invoke",["UpdateComanpy",req.headers.id,req.body.lead_company,req.body.lead_industry,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/annualrevenue',(req,res)=>{
    network.contract("invoke",["UpdateAnnualRevenue",req.headers.id,req.body.lead_annual_revenue,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/website',(req,res)=>{
    network.contract("invoke",["UpdateWebsite",req.headers.id,req.body.lead_website,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/noofemp',(req,res)=>{
    network.contract("invoke",["UpdateNoofEmployees",req.headers.id,req.body.lead_noof_employees,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/source',(req,res)=>{
    network.contract("invoke",["UpdateSource",req.headers.id,req.body.lead_source,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/status',(req,res)=>{
    network.contract("invoke",["UpdateStatus",req.headers.id,req.body.lead_status,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/rating',(req,res)=>{
    network.contract("invoke",["UpdateRating",req.headers.id,req.body.lead_rating,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/contactid',(req,res)=>{
    network.contract("invoke",["UpdateContactID",req.headers.id,req.body.contact_id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/owner',(req,res)=>{
    network.contract("invoke",["UpdateOwner",req.headers.id,req.body.lead_owner,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/leaddescription',(req,res)=>{
    network.contract("invoke",["EditLeadDescription",req.headers.id,req.body.lead_description,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

module.exports = routes