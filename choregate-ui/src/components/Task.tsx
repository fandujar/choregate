import { getTask } from "@/services/taskApi"
import { useEffect } from "react"
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

import { Steps } from "./Steps"
import { StepAdd } from "./StepAdd"
import { TaskRuns } from "./TaskRuns"
import { TaskRunCreate } from "./TaskRunCreate"
import { TaskAtom } from "@/atoms/Tasks";
import { useRecoilState } from "recoil";

import { TaskSettings } from "./TaskSettings";
import { useQuery } from "react-query";

type TaskProps = {
    taskID: string
}

export const Task = (props: TaskProps) => {
    const { taskID } = props
    const [task, setTask] = useRecoilState(TaskAtom)
    const { data, isLoading } = useQuery('task', () => getTask(taskID))

    useEffect(() => {
        if (data) {
            setTask(data)
        }
    }, [data])

    if (isLoading) {
        return <div>Loading...</div>
    }

    return (
        <section className="flex-auto m-5">
            <div className="flex">
                <h2 className="text-xl font-semibold mb-4">Task</h2>
                <div className="ml-auto">
                    <StepAdd taskID={taskID}/>
                </div>
                <div className="ml-2">
                    <TaskRunCreate taskID={taskID}/>
                </div>
                <div className="ml-2">
                    <TaskSettings/>
                </div>
            </div>
            <Card className="mb-4">
                <CardContent className="flex justify-between items-center">
                    <Table className='mt-5'>
                        <TableHeader>
                            <TableRow>
                                <TableHead>ID</TableHead>
                                <TableHead>Task name</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            <TableRow key={taskID}>
                                <TableCell>{taskID}</TableCell>
                                <TableCell>{task.name}</TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </CardContent>
            </Card>
            <h2 className="text-xl font-semibold mb-4">Steps</h2>
            <Steps taskID={taskID}/>
            <h2 className="text-xl font-semibold mb-4">Runs</h2>
            <TaskRuns taskID={taskID}/>
        </section>
    )
}