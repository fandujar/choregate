import { useEffect } from "react";
import { getSteps } from "@/services/taskApi";
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Dialog, DialogClose, DialogTrigger, DialogHeader, DialogContent, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { Label } from "./ui/label";
import { Button } from "./ui/button"
import { updateSteps } from "@/services/taskApi"
import { useRecoilState } from "recoil";

import { StepsUpdateAtom } from "@/atoms/Update";
import { StepsAtom, StepAtom } from "@/atoms/Tasks";
import { Step } from "@/types/Task";

import { SquarePen, Trash2 } from 'lucide-react';

type StepsProps = {
  taskID: string;
};

export const Steps = (props: StepsProps) => {
    const { taskID } = props;
    const [steps, setSteps] = useRecoilState(StepsAtom);
    const [update, setUpdate] = useRecoilState(StepsUpdateAtom)

    useEffect(() => {
      const steps = getSteps(taskID);
      steps.then((steps) => {
        setSteps(steps);
      });
      setUpdate(false);
    }, [update]);

    const handleDeleteStep = (index) => {
      
    } 

    return (
      <div className="space-y-4">
        <Card className="mb-4">
          <CardContent>
            <Table className='mt-5'>
              <TableHeader>
                  <TableRow>
                      <TableHead>ID</TableHead>
                      <TableHead>Step name</TableHead>
                      <TableHead>Image</TableHead>
                      <TableHead>Compute Resources</TableHead>
                      <TableHead>Script</TableHead>
                      <TableHead>Edit/Delete</TableHead>
                  </TableRow>
              </TableHeader>
              <TableBody>
              {steps?.map((step: Step, index) => (
                  <TableRow key={index}>
                      <TableCell>{index}</TableCell>
                      <TableCell>{step.name || "unamed step"}</TableCell>
                      <TableCell>{step.image}</TableCell>
                      <TableCell>{JSON.stringify(step.computeResources) || "None"}</TableCell>
                      <TableCell>{step?.script}</TableCell>
                      <TableCell>
                        <div className="flex">
                          <Button>
                            <SquarePen/>
                          </Button>
                          <Button type="submit" onClick={handleDeleteStep(index)}>
                            <Trash2/>
                          </Button>
                        </div>
                      </TableCell>
                  </TableRow>
              ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    );
};

type AddStepProps = {
  taskID: string;
};


export const AddStep = (props: AddStepProps) => {
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

  const handleEditSteps = (e: any) => {
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
          <DialogHeader>
              <DialogTitle>Add Step</DialogTitle>
              <DialogDescription>
                  Add steps to the task
              </DialogDescription>
          </DialogHeader>
          <form onSubmit={handleEditSteps}>
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