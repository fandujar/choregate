import { useEffect } from "react";
import { useRecoilState } from "recoil";

import { StepsAtom, StepAtom } from "@/atoms/Tasks";

import { Dialog, DialogClose, DialogTrigger, DialogHeader, DialogContent, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { Label } from "./ui/label";
import { Button } from "./ui/button"

import { getSteps } from "@/services/taskApi";
import { updateSteps } from "@/services/taskApi";

import { StepType } from "@/types/Task";
import { toast } from "sonner";
import { useQuery } from "react-query";

type StepAddProps = {
    taskID: string;
};
  
  
export const StepAdd = (props: StepAddProps) => {
const { taskID } = props;
const [steps, setSteps] = useRecoilState<StepType[]>(StepsAtom(taskID));
const [step, setStep] = useRecoilState<StepType>(StepAtom(`${taskID}--1`));
const { data, isLoading, refetch } = useQuery('steps', () => getSteps(taskID), {staleTime: 1000});

useEffect(() => {
    if (data) {
        setSteps(data);
    }
}, [data]);

if (isLoading) {
    return <div>Loading...</div>
}

const handleStepAdd = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const data = [...steps,{
    name: step.name,
    image: step.image,
    script: step.script
    }]

    updateSteps(taskID, data).then(() => {
        refetch()
    }).catch((err) => {
        console.log(err)
        toast.error(`${err.message}: ${err.response.data}`)
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
                <Button type="submit" className="bg-pink-700 text-white">
                    Add step
                </Button>
            </DialogClose>
        </DialogFooter>
        </form>
    </DialogContent>
</Dialog>
)
}