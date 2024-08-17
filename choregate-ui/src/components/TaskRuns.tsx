import { useEffect, useState } from 'react'
import { getTaskRuns } from '@/services/taskApi'
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

type TaskRunListProps = {
    taskID: string
    update: boolean;
    setUpdate: Function;
}

type TaskRun = {
    ID: string
}

export function TaskRuns(props: TaskRunListProps) {
    const { taskID } = props
    const [taskRuns, setTaskRuns] = useState([])
    const { update, setUpdate } = props

    useEffect(() => {
        const tasksRuns = getTaskRuns(taskID)
        tasksRuns.then((tasksRuns) => {
            setTaskRuns(tasksRuns)
        })
        setUpdate(false)
    }, [update])

    return (
        <div className="space-y-4">
            <Card className="mb-4">
                <CardContent>
                <Table className='mt-5'>
                <TableHeader>
                    <TableRow>
                        <TableHead>ID</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                {taskRuns?.map((taskRun: TaskRun, index) => (
                    <TableRow key={index} role="button">
                        <TableCell>{taskRun.ID}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
                </Table>    

                </CardContent>
            </Card>
        </div>
    )
}