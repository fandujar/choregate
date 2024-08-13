import { useEffect, useState } from 'react'
import {getTaskRuns} from '@/services/taskApi'
import TaskRun from './TaskRun'

interface TaskRun {
    ID: string
}

type TaskRunListProps = {
    taskID: string
    update: boolean
    setUpdate: (update: boolean) => void
}

export default function TaskRunList(props: TaskRunListProps) {
    const { taskID, update, setUpdate } = props
    const [taskRuns, setTaskRuns] = useState<TaskRun[]>([])

    useEffect(() => {
        const tasksRuns = getTaskRuns(taskID)
        tasksRuns.then((tasksRuns) => {
            setTaskRuns(tasksRuns)
        })
        setUpdate(false)
    }, [update])

    return (
        <div>
            <h4>Task Runs</h4>
            <ul>
                {taskRuns?.map((taskRun: TaskRun) => (
                    <li key={taskRun.ID}>
                        <TaskRun key={taskRun.ID} id={taskRun.ID}/>
                    </li>
                ))}
            </ul>
        </div>
    )
}