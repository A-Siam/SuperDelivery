-- CreateEnum
CREATE TYPE "ActionType" AS ENUM ('TRANSACTION', 'COMPENSATION');

-- CreateTable
CREATE TABLE "TransactionLog" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "transactionId" TEXT NOT NULL,
    "currentEventName" TEXT NOT NULL,
    "previousEventName" TEXT,
    "order" INTEGER NOT NULL DEFAULT 1,
    "actionType" "ActionType" NOT NULL,

    CONSTRAINT "TransactionLog_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "TransactionLog_transactionId_idx" ON "TransactionLog"("transactionId");
