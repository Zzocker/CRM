const express = require('express')
const network = require('./contract')

const routes = express.Router()

routes.post('/createNewLead',async(req,res)=>{
    try {
        const rbody = req.body
        sal = rbody.sal
        lead_firstname = rbody.lead_firstname
        lead_lastname = rbody.lead_lastname
        requester = rbody.requester
        lead_mobile = rbody.lead_mobile

        const payload = await network.contract("invoke",["CreateNewLead"," ",sal,lead_firstname,lead_lastname,lead_mobile,requester])
        id = payload.toString()
        res.status(200).send(id)
    } catch (error) {
        res.status(500).json({error})
    }
})

routes.get('/getallmylead',async (req,res)=>{
    try {
        const payload = await network.contract("query",["GetAllMyLeads","created","}}",req.headers.requester])
        res.status(200).json(JSON.parse(payload))
    } catch (error) {
        res.status(500).json(error)
    }
})
routes.get('/getlead',async(req,res)=>{
    try {

        const payload = await network.contract("query",["GetMyLead",req.headers.id,req.headers.requester])
        res.status(200).json(JSON.parse(payload))
    } catch (error) {
        res.status(500).json("No Lead found")
    }
})
routes.delete('/deletelead',async(req,res)=>{
    try {

        const payload = await network.contract("invoke",["DeleteLead",req.headers.id,req.headers.requester])
        res.status(200).json("DONE")
    } catch (error) {
        res.status(500).json(error)
    }
})
module.exports = routes
