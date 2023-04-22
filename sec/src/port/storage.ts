import { DataSource } from 'typeorm'

export const EventDataSource = new DataSource({
    type: 'postgres',
    host: process.env['EVENTS_DB_HOST'],
    port: parseInt(process.env['EVENTS_DB_PORT']!),
    username: process.env['EVENTS_DB_USER'],
    password: process.env['EVENTS_DB_PASSWORD'],
    database: process.env['EVENTS_DB_NAME'],
})
