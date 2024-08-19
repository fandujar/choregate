import { useEffect } from 'react'
import {createTask, getTasks} from '@/services/taskApi'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Button } from './ui/button'
import { useNavigate } from 'react-router-dom'
import { Card, CardContent } from './ui/card'
import { useRecoilState } from 'recoil'
import { TasksAtom } from '@/atoms/Tasks'
import { TasksUpdateAtom } from '@/atoms/Update'
import { TaskCreate } from './TaskCreate'

export function Tasks() {
    const [tasks, setTasks] = useRecoilState(TasksAtom)
    const [update, setUpdate] = useRecoilState(TasksUpdateAtom)
    const navigate = useNavigate()

    useEffect(() => {
        const tasks = getTasks()
        tasks.then((tasks) => {
            setTasks(tasks)
        })
        setUpdate(false)
    }, [update])

    return (
        <div className='flex-auto h-full m-5'>
            <div className='flex'>
                <h2 className="text-xl font-semibold mb-4">Tasks</h2>
                <div className='ml-auto'>
                    <TaskCreate/>
                </div>
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
                        {tasks?.map((task) => (
                            <TableRow key={task.id} onClick={() => {navigate(`/tasks/${task.id}`);}} role="button">
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