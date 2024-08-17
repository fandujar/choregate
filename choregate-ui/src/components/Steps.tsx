import { useEffect, useState } from "react";
import { getSteps } from "@/services/taskApi";
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { Dialog, DialogClose, DialogTrigger, DialogHeader, DialogContent, DialogTitle, DialogDescription, DialogFooter } from "./ui/dialog";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Button } from "./ui/button"
import { addSteps } from "@/services/taskApi"

type StepsProps = {
  taskID: string;
  update: boolean;
  setUpdate: Function;
  steps: Step[];
  setSteps: Function;
};

export const Steps = (props: StepsProps) => {
    const { taskID } = props;
    const { update, setUpdate } = props;
    const { steps, setSteps } = props;

    useEffect(() => {
      const steps = getSteps(taskID);
      steps.then((steps) => {
        setSteps(steps);
      });
      setUpdate(false);
    }, [update]);

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
                      <TableHead>Commands</TableHead>
                  </TableRow>
              </TableHeader>
              <TableBody>
              {steps?.map((step: Step, index) => (
                  <TableRow key={index} role="button">
                      <TableCell>{index}</TableCell>
                      <TableCell>{step.name || "unamed step"}</TableCell>
                      <TableCell>{step.image}</TableCell>
                      <TableCell>{JSON.stringify(step.computeResources) || "None"}</TableCell>
                      <TableCell>{step?.command?.join(' ')}</TableCell>
                  </TableRow>
              ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    );
};

type EditStepsProps = {
  taskID: string;
  steps: Step[];
  setSteps: Function;
  setUpdate: Function;
};


export const EditSteps = (props: EditStepsProps) => {
  const { taskID } = props;
  const { steps, setSteps } = props;
  // TODO remove this once state is properly managed
  const [step, setStep] = useState<Step>({name: '', image: '', command: '', computeResources: ''});
  const { setUpdate } = props;

  useEffect(() => {
    const steps = getSteps(taskID);
    steps.then((steps) => {
      setSteps(steps);
      setStep(steps[0]);
    });
    setUpdate(false);
  }, []);

  const handleEditSteps = (e: any) => {
    e.preventDefault()
    const data = [{
            name: step.name,
            image: step.image,
            command: step.command.split(' '),
        }]

    addSteps(taskID, data).then(() => {
        setUpdate(true)
    })
  }

  return (
    <Dialog>
      <DialogTrigger>
          <Button className="bg-pink-700 text-white">Edit Steps</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
              <DialogTitle>Edit Steps</DialogTitle>
              <DialogDescription>
                  Edit task steps here.
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
                  onChange={(e) => setStep({name: e.target.value, image: step.image, command: step.command})}
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
                  onChange={(e) => setStep({name: step.name, image: e.target.value, command: step.command})}
              />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="command" className="text-right">
                  Command
              </Label>
              <Input
                  id="command"
                  defaultValue="echo 'Hello, World!'"
                  value={step.command}
                  className="col-span-3"
                  onChange={(e) => setStep({name: step.name, image: step.image, command: e.target.value})}
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