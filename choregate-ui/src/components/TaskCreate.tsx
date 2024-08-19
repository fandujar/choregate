import { createTask } from "@/services/taskApi"

import { TaskAtom } from "@/atoms/Tasks";
import { useRecoilState } from "recoil";
import { Dialog, DialogClose, DialogContent, DialogTrigger, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { TasksUpdateAtom } from "@/atoms/Update";

export const TaskCreate = () => {
    const [task, setTask] = useRecoilState(TaskAtom)
    const [update, setUpdate] = useRecoilState(TasksUpdateAtom)

    const handleTaskCreate = (e: any) => {
        e.preventDefault()
        let response  = createTask(task)
        response.then(() => {
            setUpdate(true)
        })
        setUpdate(true)
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
                        onChange={(e) => setTask({name: e.target.value})}
                    />
                </div>
            </div>
            <DialogFooter>
                <DialogClose asChild>
                    <Button type="submit">Create</Button>
                </DialogClose>
            </DialogFooter>
            </form>
        </DialogContent>
    </Dialog>
    )
}