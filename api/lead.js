const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.post('/createNewLead',(req,res)=>{
    rbody = req.body
    input=rbody.input
    requester=rbody.requester
    network.contract("invoke",["CreateNewLead",input,requester],(err,payload)=>{
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
        input=rbody.input
        requester = rbody.requester

        network.contract("invoke",["CreateLeadFromContact",input,requester],(err,payload)=>{
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


    