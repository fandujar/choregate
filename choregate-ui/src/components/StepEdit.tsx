import { StepAtom, StepsAtom } from "@/atoms/Tasks";
import { getSteps, updateSteps } from "@/services/taskApi";
import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { SquarePen } from "lucide-react";

import { Dialog, DialogContent, DialogTrigger, DialogFooter, DialogHeader } from "./ui/dialog";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { StepsUpdateAtom } from "@/atoms/Update";
import { DialogClose } from "@radix-ui/react-dialog";
import { toast } from "sonner";

type StepEditProps = {
    taskID: string;
    stepIndex: number;
  }
  
  export const StepEdit = (props: StepEditProps) => {
    const { taskID, stepIndex } = props;
    const [step, setStep] = useRecoilState(StepAtom);
    const [steps, setSteps] = useRecoilState(StepsAtom);
    const [_, setUpdate] = useRecoilState(StepsUpdateAtom)
  
    useEffect(() => {
      let response = getSteps(taskID);
      response.then((steps) => {
        setSteps(steps);
      });
    }, [])
  
    useEffect(() => {
      setStep(steps[stepIndex])
    }, [])
  
    const handleEditStep = (index: number) => {
        let data = [...steps]
        data[index] = step
        updateSteps(taskID, data).then(() => {
          setUpdate(true)
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
            Edit Step
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