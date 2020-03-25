const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.post('/createnewdeal',(req,res)=>{
    rbody = req.body
    leadSourceID = rbody.deal_lead_source
    dname = rbody.deal_name
    accoutName = rbody.deal_account_name
    accountID = rbody.deal_aacount_id
    dealtype = rbody.deal_type
    amount  = rbody.deal_amount
    closingDate = rbody.deal_closing_date
    stage = rbody.deal_stage
    probability = rbody.deal_probility
    requester = rbody.requester

    network.contract("invoke",["CreatNewDeal",leadSourceID,dname,accoutName,accountID,dealtype,amount,closingDate,stage,probability,requester],(err,payload)=>{
        if (err){
            res.status(500).json({error})
        }else{
            id = payload.toString()
        res.status(200).send(id)
        }
    })
   
})

routes.delete('/delete',(req,res)=>{
    network.contract("invoke",["DeleteDeal",req.headers.id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.responses[0].response.message)
        }
        else {res.status(200).json(`DONE`)}
    })
})


routes.get('/getalldeal',(req,res)=>{
    network.contract("query",["GetAllMyDeals",req.headers.type,"}}",req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(`No deal`)
        }else{res.status(200).json(JSON.parse(payload))}
    })
})

routes.get('/getdeal',(req,res)=>{
    
    network.contract("query",["GetMyDeal",req.headers.id,req.headers.requester],(err,payload)=>{
        if (err){
            res.status(500).json(err.message)
        }
        else {res.status(200).json(JSON.parse(payload))}
    })
})  
module.exports = routes



    