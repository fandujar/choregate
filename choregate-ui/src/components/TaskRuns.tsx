import { useEffect } from 'react'
import { getTaskRuns, getTaskRunStatus, getTaskRunLogs } from '@/services/taskApi'
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { useRecoilState } from 'recoil';
import { TaskUpdateAtom } from '@/atoms/Update';
import { TaskRunsAtom } from '@/atoms/Tasks';
import { TaskRunLogsAtom } from '@/atoms/Tasks';
import { TaskRun } from '@/types/Task';
import { Sheet, SheetClose, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger } from './ui/sheet';

type TaskRunListProps = {
    taskID: string
}

export function TaskRuns(props: TaskRunListProps) {
    const { taskID } = props
    const [taskRuns, setTaskRuns] = useRecoilState(TaskRunsAtom)
    const [update, setUpdate] = useRecoilState(TaskUpdateAtom)

    useEffect(() => {
        const response = getTaskRuns(taskID)
        response.then((tasksRuns) => {
            setTaskRuns(tasksRuns)
        })
        setUpdate(false)
    }, [update])
    
    const handleTaskRunStatus = (taskRunID: string) => {
        let response = getTaskRunStatus(taskID, taskRunID)
        response.then(({conditions}) => {
            return conditions
        })
    }

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
                        <TableCell>{handleTaskRunStatus(taskRun.ID)}</TableCell>
                        <TableCell><TaskRunsLogs taskID={taskID} taskRunID={taskRun.ID}/></TableCell>
                    </TableRow>
                ))}
                </TableBody>
                </Table>
                </CardContent>
            </Card>
        </div>
    )
}

type TaskRunsLogsProps = {
    taskID: string
    taskRunID: string
}

const TaskRunsLogs = (props: TaskRunsLogsProps) => {
    const { taskID, taskRunID } = props
    const [taskRunLogs, setTaskRunLogs] = useRecoilState(TaskRunLogsAtom)
    const [update, setUpdate] = useRecoilState(TaskUpdateAtom)

    useEffect(() => {
        let response = getTaskRunLogs(taskID, taskRunID)
        response.then((logs) => {
            setTaskRunLogs(logs)
        })
        setUpdate(false)
    }, [])

    const handleViewLogs = () => {
        let response = getTaskRunLogs(taskID, taskRunID)
        response.then((logs) => {
            setTaskRunLogs(logs)
        })
        setUpdate(false)
    }

    return (
        <Sheet>
            <SheetTrigger>
                <Button onClick={handleViewLogs}>View Logs</Button>
            </SheetTrigger>
            <SheetContent>
                <SheetHeader>
                    <SheetTitle>Task Run Logs</SheetTitle>
                    <SheetDescription>logs from the task execution</SheetDescription>
                </SheetHeader>

                <pre>
                    {taskRunLogs}
                </pre>


                <SheetFooter>
                    <SheetClose asChild>
                        <Button type="submit">Close</Button>
                    </SheetClose>
                </SheetFooter>
            </SheetContent>
        </Sheet>
    )
}