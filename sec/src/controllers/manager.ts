import * as express from 'express'
import { toRollbackRequestBody } from '../mappers/rollback-mapper'
import TransactionStorageService from '../services/transaction-storage'
import TransactionCompensationService from '../services/compensation'

export const managerRouter = express.Router()

managerRouter.post('/rollback', async (req, res) => {
    const transactionCompensationService = new TransactionCompensationService()
    const transactionStorageService = new TransactionStorageService(
        transactionCompensationService
    )
    rollbackHandler(req, res, transactionStorageService)
})

export async function rollbackHandler(
    req: express.Request,
    res: express.Response,
    transactionStorageService: TransactionStorageService
) {
    const body = req.body()
    const rollbackRequest = toRollbackRequestBody(body)
    await transactionStorageService.rollback(rollbackRequest)
    res.status(200)
}
