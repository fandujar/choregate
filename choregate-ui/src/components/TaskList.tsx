import { useEffect, useState } from 'react'
import {createTask, getTasks} from '@/services/taskApi'
import Task from './Task'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { Link } from 'react-router-dom'

interface Task {
    id: string
    name: string
}


export function TaskList() {
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
        <div className='flex-auto flex-col'>
            <div className='flex'>
            <h1 className='text-pink-700 text-xl'>Tasks</h1>
            <Button 
                className="ml-auto bg-pink-700 text-white"
                onClick={() => {createTask(); setUpdate(true)}}
            >
                Create Task
            </Button>
            </div>
            <Table className='mt-5'>
                <TableHeader>
                    <TableRow>
                        <TableHead>ID</TableHead>
                        <TableHead>Task name</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                {tasks?.map((task: Task) => (
                    <TableRow key={task.id}>
                        <TableCell>
                            <Link to={`/tasks/${task.id}`}>{task.id}</Link>
                        </TableCell>
                        <TableCell>{task.name}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
        </div>
    )
}