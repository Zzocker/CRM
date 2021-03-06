const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const lead = require('./lead')
const deal = require('./deal')

const PORT = "3000"

const api = express()
api.use(cors())

const logger = (req,res,next)=>{
    console.log(`${req.protocol}://${req.get('host')}${req.originalUrl}`)
    next()
}

api.use(bodyParser.urlencoded({ extended: false }))
api.use(bodyParser.json())

api.use('/lead',lead)
api.use(cors())
api.use('/deal',deal)

api.use(logger)

api.listen(PORT,()=>{
    console.log(`Listening on port : ${PORT}`)
})