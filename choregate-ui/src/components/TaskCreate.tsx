import { createTask, getTasks } from "@/services/taskApi"
import { Dialog, DialogClose, DialogContent, DialogTrigger, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { toast } from "sonner";
import { useState } from "react";
import { TaskType } from "@/types/Task";
import { useQuery } from "react-query";

export const TaskCreate = () => {
    const [task, setTask] = useState<TaskType>({name: ""})
    const {refetch} = useQuery('tasks', getTasks, {staleTime: 1000})

    const handleTaskCreate = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const response  = createTask(task)
        response.then(() => {
            refetch()
        }).catch((err) => {
            console.log(err)
            toast.error(`${err.message}: ${err.response.data}`)
        })
    }

    return (
    <Dialog>
        <DialogTrigger>
            <Button className="bg-pink-700 text-white">Create Task</Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
            <form onSubmit={handleTaskCreate}>
            <DialogHeader>
                <DialogTitle>Create task</DialogTitle>
                <DialogDescription>
                    Create a new task
                </DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
                <div className="grid grid-cols-4 items-center gap-4">
                    <Label htmlFor="name" className="text-right">
                        Name
                    </Label>
                    <Input
                        id="name"
                        defaultValue="unamed"
                        className="col-span-3"
                        value={task.name}
                        onChange={(e) => setTask({...task, name: e.target.value})}
                    />
                </div>
            </div>
            <DialogFooter>
                <DialogClose asChild>
                    <Button type="submit" className="bg-pink-700 text-white">
                        Create task
                    </Button>
                </DialogClose>
            </DialogFooter>
            </form>
        </DialogContent>
    </Dialog>
    )
}