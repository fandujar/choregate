import { useEffect } from "react";
import { useRecoilState } from "recoil";

import { StepsUpdateAtom } from "@/atoms/Update";
import { StepsAtom, StepAtom } from "@/atoms/Tasks";

import { Dialog, DialogClose, DialogTrigger, DialogHeader, DialogContent, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { Label } from "./ui/label";
import { Button } from "./ui/button"

import { getSteps } from "@/services/taskApi";
import { updateSteps } from "@/services/taskApi";

type StepAddProps = {
    taskID: string;
};
  
  
export const StepAdd = (props: StepAddProps) => {
const { taskID } = props;
const [steps, setSteps] = useRecoilState(StepsAtom);
const [update, setUpdate] = useRecoilState(StepsUpdateAtom)
const [step, setStep] = useRecoilState(StepAtom);

useEffect(() => {
    let response = getSteps(taskID);
    response.then((steps) => {
    setSteps(steps);
    });
    setUpdate(false);
}, [update]);

useEffect(() => {
    setStep({name: '', image: 'ubuntu', script: 'echo "Hello, World!"'})
}, [])

const handleStepAdd = (e: any) => {
    e.preventDefault()

    let data = [...steps,{
    name: step.name,
    image: step.image,
    script: step.script
    }]

    updateSteps(taskID, data).then(() => {
        setUpdate(true)
    })
}

return (
    <Dialog>
    <DialogTrigger>
        <Button className="bg-pink-700 text-white">Add Step</Button>
    </DialogTrigger>
    <DialogContent className="sm:max-w-[425px]">
        <form onSubmit={handleStepAdd}>
        <DialogHeader>
            <DialogTitle>Add Step</DialogTitle>
            <DialogDescription>
                Add steps to the task
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
                value={step.name}
                onChange={(e) => setStep({name: e.target.value, image: step.image, script: step.script})}
            />
        </div>
        <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="image" className="text-right">
                Image
            </Label>
            <Input
                id="image"
                defaultValue="ubuntu"
                className="col-span-3"
                value={step.image}
                onChange={(e) => setStep({name: step.name, image: e.target.value, script: step.script})}
            />
        </div>
        <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="script" className="text-right">
                Script
            </Label>
            <Textarea
                id="script"
                defaultValue="echo 'Hello, World!'"
                value={step.script}
                className="col-span-3"
                onChange={(e) => setStep({name: step.name, image: step.image, script: e.target.value})}
            />
        </div>
        </div>
        <DialogFooter>
            <DialogClose asChild>
                <Button type="submit">Save changes</Button>
            </DialogClose>
        </DialogFooter>
        </form>
    </DialogContent>
</Dialog>
)
}