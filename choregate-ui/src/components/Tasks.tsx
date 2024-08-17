import { useEffect, useState } from 'react'
import {createTask, getTasks} from '@/services/taskApi'
import { Task } from './Task'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { useNavigate } from 'react-router-dom'
import { Card, CardContent } from './ui/card'

interface Task {
    id: string
    name: string
}


export function Tasks() {
    const [tasks, setTasks] = useState<Task[]>([])
    const [update, setUpdate] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        const tasks = getTasks()
        tasks.then((tasks) => {
            setTasks(tasks)
        })
        setUpdate(false)
    }, [update])

    return (
        <div className='flex-auto h-screen m-5'>
            <div className='flex'>
            <h2 className="text-xl font-semibold mb-4">Tasks</h2>
            <Button 
                className="ml-auto bg-pink-700 text-white"
                onClick={() => {createTask(); setUpdate(true);}}
            >
                Create Task
            </Button>
            </div>
            <Card>
                <CardContent>
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
                </CardContent>
            </Card>
        </div>
    )
}