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


    