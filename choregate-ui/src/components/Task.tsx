import { getTask } from "@/services/taskApi"
import { useEffect } from "react"
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

import { Steps, AddStep } from "./Steps"
import { TaskRuns, RunTask } from "./TaskRuns"
import { TaskUpdateAtom } from "@/atoms/Update";
import { TaskAtom } from "@/atoms/Tasks";
import { useRecoilState } from "recoil";


type TaskProps = {
    id: string
}

export const Task = (props: TaskProps) => {
    const { id } = props
    const [task, setTask] = useRecoilState(TaskAtom)
    const [update, setUpdate] = useRecoilState(TaskUpdateAtom)

    useEffect(() => {
        let task = getTask(id)
        task.then((task) => {
            setTask(task)
        })
        setUpdate(false)
    }, [update])

    return (
        <section className="flex-auto m-5">
            <div className="flex">
                <h2 className="text-xl font-semibold mb-4">Task</h2>
                <div className="ml-auto">
                    <AddStep taskID={id}/>
                </div>
                <div className="ml-2">
                    <RunTask taskID={id}/>
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
                            <TableRow key={id}>
                                <TableCell>{id}</TableCell>
                                <TableCell>{task.name}</TableCell>
                            </TableRow>
                        </TableBody>
                    </Table>
                </CardContent>
            </Card>
            <h2 className="text-xl font-semibold mb-4">Steps</h2>
            <Steps taskID={id}/>
            <h2 className="text-xl font-semibold mb-4">Runs</h2>
            <TaskRuns taskID={id}/>
        </section>
    )
}