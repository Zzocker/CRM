const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.post('/createNewLead',(req,res)=>{
    rbody = req.body
    input = JSON.stringify(rbody)
    requester=req.headers.requester
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
        input = JSON.stringify(rbody)
        requester = rbody.requester

        network.contract("invoke",["CreateLeadFromContact",input,requester],(err,payload)=>{
            if (err){
                res.status(500).json(err)
            }else{
                id = payload.toString()
                res.status(200).send(payload)
            }
        })
        
})

routes.put('/update',(req,res)=>{
    input = JSON.stringify(req.body)
    lead_id = req.headers.lead_id
    requester=req.headers.requester
    // console.log(input)
    network.contract("invoke",["UpdateLead",lead_id,input,requester],(err,payload)=>{
        if (err){
            res.status(500).json({err})
        }else{
            res.status(200).json("DONE")
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


    


