import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm'

@Entity()
export class TransactionalState {
    @PrimaryGeneratedColumn()
    id: number

    @Column()
    operationId: number

    @Column()
    order: number

    @Column()
    currentEventName: string

    @Column()
    previousEventName: string
}
