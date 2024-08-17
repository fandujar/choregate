import { useState } from "react"
import { runTask, getTask } from "@/services/taskApi"
import { Button } from "./ui/button"
import { useEffect } from "react"
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

import { Steps, EditSteps } from "./Steps"
import { TaskRuns } from "./TaskRuns"


type TaskProps = {
    id: string
}

type Task = {
    id: string
    name: string
}

type Step = {
    name: string;
    image: string;
    command: string;
    computeResources: string;
  }

export const Task = (props: TaskProps) => {
    const { id } = props
    const [task, setTask] = useState<Task>({id: '', name: ''})
    const [steps, setSteps] = useState<Step[]>([]);
    const [update, setUpdate] = useState(false)

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
                    <EditSteps taskID={id} steps={steps} setSteps={setSteps} setUpdate={setUpdate}/>
                </div>
                <Button className="ml-2 bg-pink-700 text-white" onClick={() => {runTask(id);setUpdate(true);}}>Run Task</Button>
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
            <Steps taskID={id} update={update} setUpdate={setUpdate} steps={steps} setSteps={setSteps}/>
            <h2 className="text-xl font-semibold mb-4">Runs</h2>
            <TaskRuns taskID={id} update={update} setUpdate={setUpdate}/>
        </section>
    )
}