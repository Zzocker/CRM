const express = require('express')
const network = require('./contract')

const routes = express.Router()


routes.put('/owner',(req,res)=>{
    network.contract("invoke",["UpdateDealOwner",req.headers.id,req.body.deal_owner,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/dealname',(req,res)=>{
    network.contract("invoke",["UpdateDealName",req.headers.id,req.body.deal_name,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/dealdate',(req,res)=>{
    network.contract("invoke",["UpdateDealDate",req.headers.id,req.body.deal_date,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/potentialname',(req,res)=>{
    network.contract("invoke",["UpdateDealPotentialName",req.headers.id,req.body.deal_potential_name,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/accountname',(req,res)=>{
    network.contract("invoke",["UpdateDealAccountName",req.headers.id,req.body.deal_account_name,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/type',(req,res)=>{
    network.contract("invoke",["UpdateDealType",req.headers.id,req.body.deal_type,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/nextstep',(req,res)=>{
    network.contract("invoke",["UpdateDealNextStep",req.headers.id,req.body.deal_next_step,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/nextstepdate',(req,res)=>{
    network.contract("invoke",["UpdateDealNextStepDate",req.headers.id,req.body.deal_next_stepdate,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/leadsource',(req,res)=>{
    network.contract("invoke",["UpdateDealLeadSource",req.headers.id,req.body.deal_lead_source,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/campaignsource',(req,res)=>{
    network.contract("invoke",["UpdateCampaignSource",req.headers.id,req.body.deal_campaign_source,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/contactname',(req,res)=>{
    network.contract("invoke",["UpdateDealContactName",req.headers.id,req.body.deal_contact_name,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/currencycode',(req,res)=>{
    network.contract("invoke",["UpdateDealCurrencyCode",req.headers.id,req.body.deal_currency_code,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/amount',(req,res)=>{
    network.contract("invoke",["UpdateDealAmount",req.headers.id,req.body.deal_amount,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/closingdate',(req,res)=>{
    network.contract("invoke",["UpdateClosingDate",req.headers.id,req.body.deal_closing_date,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/stage',(req,res)=>{
    network.contract("invoke",["UpdateDealStage",req.headers.id,req.body.deal_stage,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/probility',(req,res)=>{
    network.contract("invoke",["UpdateDealProbility",req.headers.id,req.body.deal_probility,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})

routes.put('/expectedrevencue',(req,res)=>{
    network.contract("invoke",["UpdateExpectedRevenue",req.headers.id,req.body.deal_expected_revenue,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
routes.put('/description',(req,res)=>{
    network.contract("invoke",["EditDealDescription",req.headers.id,req.body.deal_description,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }else{
            res.status(200).send(`DONE`)
        }
    })
   
})
module.exports = routes