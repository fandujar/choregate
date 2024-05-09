import { useEffect, useState } from 'react'

interface Task {
    id: string
}

export default function TaskList() {
    const [tasks, setTasks] = useState<Task[]>([])

    useEffect(() => {
        setTasks([])
    }, [])

    return (
        <div>
            <h1>Tasks</h1>
            <ul>
                {tasks?.map((task: Task) => (
                    <li key={task.id}>{task.id}</li>
                ))}
            </ul>
        </div>
    )
}