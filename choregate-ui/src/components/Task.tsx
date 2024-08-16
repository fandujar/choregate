import { useState } from "react"
import { addSteps, runTask, getTask } from "@/services/taskApi"
import { Button } from "./ui/button"
import TaskRunList from "./TaskRunList"
import { useEffect } from "react"
import { Steps } from "./Steps"

type TaskProps = {
    id: string
}

export const Task = (props: TaskProps) => {
    const { id } = props
    const [task, setTask] = useState({})

    useEffect(() => {
        let task = getTask(id)
        setTask(task)
    }, [])
    return (
        <div className="flex flex-col h-screen p-10">
            <div className="flex">
                <h1 className='text-pink-700 text-xl'>Task {id}</h1>
                <div className="ml-auto">
                    <Button className="ml-auto bg-pink-700 text-white" onClick={() => {addSteps(id)}}>Add Steps</Button>
                    <Button className="ml-auto bg-pink-700 text-white" onClick={() => {runTask(id);}}>Run Task</Button>
                </div>
            </div>
            <div className="mt-5">
                <h2 className="text-pink-700 text-lg">Steps</h2>
                <Steps taskID={id}/>
            </div>
            <div className="mt-5">
                <h2 className="text-pink-700 text-lg">Runs</h2>
                <TaskRunList taskID={id} />
            </div>
        </div>
    )
}