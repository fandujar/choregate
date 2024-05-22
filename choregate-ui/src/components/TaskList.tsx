import { useEffect, useState } from 'react'
import {createTask, getTasks} from '../services/taskApi'
import Task from './Task'
import TaskRunList from './TaskRunList'

interface Task {
    id: string
}


export default function TaskList() {
    const [tasks, setTasks] = useState<Task[]>([])
    const [update, setUpdate] = useState(true)

    useEffect(() => {
        const tasks = getTasks()
        tasks.then((tasks) => {
            setTasks(tasks)
        })
        setUpdate(false)
    }, [update])

    return (
        <div>
            <h1>Tasks</h1>
            <button onClick={() => {createTask(); setUpdate(true)}}>Create Task</button>
            <ul>
                {tasks?.map((task: Task) => (
                    <li key={task.id}>
                        <Task 
                            key={`task-${task.id}`}
                            id={task.id}
                            setUpdate={setUpdate}
                        />
                        <TaskRunList
                            key={`taskrun-${task.id}`}
                            taskID={task.id}
                            update={update}
                            setUpdate={setUpdate}
                        />
                    </li>
                ))}
            </ul>
        </div>
    )
}