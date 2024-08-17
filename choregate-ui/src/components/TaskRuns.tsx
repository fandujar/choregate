import { useEffect, useState } from 'react'
import { getTaskRuns, getTaskRunLogs, getTaskRunStatus } from '@/services/taskApi'
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'

type TaskRunListProps = {
    taskID: string
    update: boolean;
    setUpdate: Function;
}

type TaskRun = {
    ID: string,
    CreatedAt: string,
}

export function TaskRuns(props: TaskRunListProps) {
    const { taskID } = props
    const [taskRuns, setTaskRuns] = useState([])
    const [taskRunStatus, setTaskRunStatus] = useState({})

    const { update, setUpdate } = props

    useEffect(() => {
        const tasksRuns = getTaskRuns(taskID)
        tasksRuns.then((tasksRuns) => {
            setTaskRuns(tasksRuns)
        })
        setUpdate(false)
    }, [update])

    useEffect(() => {
        taskRuns?.map((taskRun: TaskRun) => {
            const taskRunStatus = getTaskRunStatus(taskID, taskRun.ID)
            taskRunStatus.then((taskRunStatus) => {
                setTaskRunStatus({taskRunID: taskRunStatus})
            })
        })
    }, [taskRuns])

    return (
        <div className="space-y-4">
            <Card className="mb-4">
                <CardContent>
                <Table className='mt-5'>
                <TableHeader>
                    <TableRow>
                        <TableHead>ID</TableHead>
                        <TableHead>Created At</TableHead>
                        <TableHead>Status</TableHead>
                        <TableHead>Logs</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                {taskRuns?.map((taskRun: TaskRun, index) => (
                    <TableRow key={index}>
                        <TableCell>{taskRun.ID}</TableCell>
                        <TableCell>{taskRun.CreatedAt}</TableCell>
                        <TableCell>{taskRunStatus[taskRun.ID] || "Pending"}</TableCell>
                        <TableCell><Button>View Logs</Button></TableCell>
                    </TableRow>
                ))}
                </TableBody>
                </Table>    

                </CardContent>
            </Card>
        </div>
    )
}