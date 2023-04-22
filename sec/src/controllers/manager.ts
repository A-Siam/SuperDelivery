import * as express from 'express'
import { toRollbackRequestBody } from '../mappers/rollback-mapper'
import { rollbackTranasction } from '../services/transaction-storage'

export const managerRouter = express.Router()

managerRouter.post('/rollback', async (req, res) => {
    const body = req.body()
    const rollbackRequest = toRollbackRequestBody(body)
    await rollbackTranasction(rollbackRequest)
    res.status(200)
})
