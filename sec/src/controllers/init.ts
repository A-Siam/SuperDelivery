import express from 'express'
import { managerRouter } from './manager'

const app = express()

export function initializeFunction() {
    // middlewares
    app.use(express.json())

    // routers
    app.use('/manage', managerRouter)
}
