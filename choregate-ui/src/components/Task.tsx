import { getTask } from "@/services/taskApi"
import { useEffect } from "react"
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

import { Steps } from "./Steps"
import { StepAdd } from "./StepAdd"
import { TaskRuns } from "./TaskRuns"
import { TaskRunCreate } from "./TaskRunCreate"
import { TaskUpdateAtom } from "@/atoms/Update";
import { TaskAtom } from "@/atoms/Tasks";
import { useRecoilState } from "recoil";
import { toast } from "sonner";
import { TaskSettings } from "./TaskSettings";

type TaskProps = {
    taskID: string
}

export const Task = (props: TaskProps) => {
    const { taskID } = props
    const [task, setTask] = useRecoilState(TaskAtom)
    const [update, setUpdate] = useRecoilState(TaskUpdateAtom)

    useEffect(() => {
        const task = getTask(taskID)
        task.then((task) => {
            setTask(task)
        }).catch((error) => {
            toast.error(`${error.message}: ${error.response.data}`)
        })
        setUpdate(false)
    }, [update])

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
                    <TaskSettings taskID={taskID}/>
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