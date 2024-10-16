import { StepAtom, StepsAtom } from "@/atoms/Tasks";
import { getSteps, updateSteps } from "@/services/taskApi";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { SquarePen } from "lucide-react";

import { Dialog, DialogContent, DialogTrigger, DialogFooter, DialogHeader, DialogClose, DialogTitle } from "./ui/dialog";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { toast } from "sonner";
import { useQuery } from "react-query";

type StepEditProps = {
    taskID: string;
    stepIndex: number;
  }
  
  export const StepEdit = (props: StepEditProps) => {
    const { taskID, stepIndex } = props;
    const [step, setStep] = useRecoilState(StepAtom(`${taskID}-${stepIndex}`));
    const [steps, setSteps] = useRecoilState(StepsAtom(taskID));
    const { data, isLoading, refetch } = useQuery('steps', () => getSteps(taskID), {staleTime: 1000})

    useEffect(() => {
        if (data) {
            setSteps(data)
        }
    }, [data])

    useEffect(() => {
        if (steps) {
            setStep(steps[stepIndex])
        }
    }, [steps])

    if (isLoading) {
        return <div>Loading...</div>
    }
  
    const handleEditStep = (index: number) => {
        const data = [...steps]
        data[index] = step
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
          <Button variant={"ghost"}>
            <SquarePen/>
          </Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Edit Step</DialogTitle>
          </DialogHeader>
          <div>
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              defaultValue="unamed"
              value={step.name}
              onChange={(e) => setStep({...step, name: e.target.value})}
            />
            <Label htmlFor="image" className="text-right">
              Image
            </Label>
            <Input
              id="image"
              defaultValue="ubuntu"
              value={step.image}
              onChange={(e) => setStep({...step, image: e.target.value})}
            />
            <Label htmlFor="script" className="text-right">
              Script
            </Label>
            <Textarea
              id="script"
              defaultValue="echo 'Hello, World!'"
              value={step.script}
              onChange={(e) => setStep({...step, script: e.target.value})}
            />
          </div>
          
          <DialogFooter>
            <DialogClose asChild>
              <Button onClick={() => handleEditStep(stepIndex)} className="bg-pink-700 text-white">
                Save changes
              </Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    )
  }