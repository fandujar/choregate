import { useEffect, useState } from 'react'
import {createTask, getTasks} from '@/services/taskApi'
import { Task } from './Task'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { useNavigate } from 'react-router-dom'

interface Task {
    id: string
    name: string
}


export function TaskList() {
    const [tasks, setTasks] = useState<Task[]>([])
    const navigate = useNavigate()

    useEffect(() => {
        const tasks = getTasks()
        tasks.then((tasks) => {
            setTasks(tasks)
        })
    }, [])

    return (
        <div className='flex-auto h-screen p-10'>
            <div className='flex'>
            <h1 className='text-pink-700 text-xl'>Tasks</h1>
            <Button 
                className="ml-auto bg-pink-700 text-white"
                onClick={() => {createTask()}}
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
                    <TableRow key={task.id} onClick={() => navigate(`/tasks/${task.id}`)} role="button">
                        <TableCell>{task.id}</TableCell>
                        <TableCell>{task.name}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
        </div>
    )
}