import { useEffect } from 'react'
import { getTaskRuns, getTaskRunLogs, runTask } from '@/services/taskApi'
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { useRecoilState } from 'recoil';
import { TaskRunsUpdateAtom } from '@/atoms/Update';
import { TaskRunsAtom } from '@/atoms/Tasks';
import { TaskRunLogsAtom } from '@/atoms/Tasks';
import { TaskRunType } from '@/types/Task';
import { Sheet, SheetClose, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger } from './ui/sheet';
import { ScrollArea } from './ui/scroll-area';
import { toast } from 'sonner';

type TaskRunListProps = {
    taskID: string
}

export function TaskRuns(props: TaskRunListProps) {
    const { taskID } = props
    const [taskRuns, setTaskRuns] = useRecoilState<TaskRunType[]>(TaskRunsAtom)
    const [update, setUpdate] = useRecoilState(TaskRunsUpdateAtom)

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
                {taskRuns?.map((taskRun: TaskRunType, index) => (
                    <TableRow key={index}>
                        <TableCell>{taskRun.ID}</TableCell>
                        <TableCell>{taskRun.CreatedAt}</TableCell>
                        <TableCell>{taskRun.Status}</TableCell>
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

type RunTaskProps = {
    taskID: string
}

export const RunTask = (props: RunTaskProps) => {
    const { taskID } = props
    const [_, setUpdate] = useRecoilState(TaskRunsUpdateAtom)

    const handleRunTask = (e: any) => {
        e.preventDefault()
        runTask(taskID).then(() => {
            setUpdate(true)
        }).catch((err) => {
            console.log(err)
            toast.error(`${err.message}: ${err.response.data}`)
            setUpdate(true)
        })
    }

    return (
        <Button className="bg-pink-700 text-white" onClick={handleRunTask}>
            Run Task
        </Button>
    )
}

type TaskRunsLogsProps = {
    taskID: string
    taskRunID: string
}

const TaskRunsLogs = (props: TaskRunsLogsProps) => {
    const { taskID, taskRunID } = props
    const [taskRunLogs, setTaskRunLogs] = useRecoilState(TaskRunLogsAtom)

    const handleViewLogs = () => {
        let response = getTaskRunLogs(taskID, taskRunID)
        response.then((logs) => {
            setTaskRunLogs(logs)
        })
    }

    return (
        <Sheet>
            <SheetTrigger>
                <Button onClick={handleViewLogs} variant={"ghost"}>View Logs</Button>
            </SheetTrigger>
            <SheetContent className="w-[540px] sm:w-[940px] sm:max-w-[940px]">
                <SheetHeader>
                    <SheetTitle>Task Run Logs</SheetTitle>
                    <SheetDescription>logs from the task execution</SheetDescription>
                </SheetHeader>
                <ScrollArea className='h-full w-full'>
                    <pre className='p-10'>
                        {taskRunLogs}
                    </pre>
                </ScrollArea>
                <SheetFooter>
                    <SheetClose asChild>
                        <Button type="submit">Close</Button>
                    </SheetClose>
                </SheetFooter>
            </SheetContent>
        </Sheet>
    )
}