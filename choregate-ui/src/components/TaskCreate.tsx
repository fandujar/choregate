import { createTask } from "@/services/taskApi"
import { useRecoilState } from "recoil";
import { Dialog, DialogClose, DialogContent, DialogTrigger, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { TasksUpdateAtom } from "@/atoms/Update";
import { toast } from "sonner";
import { useState } from "react";
import { TaskType } from "@/types/Task";

export const TaskCreate = () => {
    const [task, setTask] = useState<TaskType>({name: ""})
    const [_, setUpdate] = useRecoilState(TasksUpdateAtom)

    const handleTaskCreate = (e: any) => {
        e.preventDefault()
        let response  = createTask(task)
        response.then(() => {
            setUpdate(true)
        }).catch((err) => {
            console.log(err)
            toast.error(`${err.message}: ${err.response.data}`)
            setUpdate(true)
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