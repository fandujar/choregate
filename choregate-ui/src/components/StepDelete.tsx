import { StepsAtom } from "@/atoms/Tasks";
import { updateSteps } from "@/services/taskApi";
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogFooter, AlertDialogHeader, AlertDialogTrigger } from "./ui/alert-dialog";
import { useRecoilState } from "recoil";
import { Button } from "./ui/button";
import { Trash2 } from "lucide-react";

type StepDeleteProps = {
    taskID: string;
    stepIndex: number;
  }
  
  // Use AlertDialog to confirm deletion
  export const StepDelete = (props: StepDeleteProps) => {
    const { taskID, stepIndex } = props;
    const [steps, setSteps] = useRecoilState(StepsAtom);
  
    const handleDeleteStep =  (index: number) => {
        let data = steps.filter((_, i) => i !== index)
        updateSteps(taskID, data).then(() => {
            setSteps(data)
        })
    } 
  
    return (
      <AlertDialog>
        <AlertDialogTrigger>
          <Button>
            <Trash2/>
          </Button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            Are you sure you want to delete this step?
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={() => handleDeleteStep(stepIndex)}>Confirm</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    )
  }