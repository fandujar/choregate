import { useEffect } from 'react'
import { getTaskRuns } from '@/services/taskApi'
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { useRecoilState } from 'recoil';
import { TaskRunsUpdateAtom } from '@/atoms/Update';
import { TaskRunsAtom } from '@/atoms/Tasks';
import { TaskRunType } from '@/types/Task';
import { TaskRunLogs } from './TaskRunLogs';


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
                        <TableCell><TaskRunLogs taskID={taskID} taskRunID={taskRun.ID}/></TableCell>
                    </TableRow>
                ))}
                </TableBody>
                </Table>
                </CardContent>
            </Card>
        </div>
    )
}