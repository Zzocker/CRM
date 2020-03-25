const express = require('express')
const lead = require('./lead')
const leadupdate = require('./leadupdate')
const PORT = "3000"

const api = express()

const logger = (req,res,next)=>{
    console.log(`${req.protocol}://${req.get('host')}${req.originalUrl}`)
    next()
}

api.use(express.json())


api.use('/lead',lead)
api.use('/leadupdate',leadupdate)

api.use(logger)

api.listen(PORT,()=>{
    console.log(`Listening on port : ${PORT}`)
})