/*
  Warnings:

  - Added the required column `serviceName` to the `TransactionLog` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "TransactionLog" ADD COLUMN     "serviceName" TEXT NOT NULL;
