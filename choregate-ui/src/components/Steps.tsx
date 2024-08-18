import { useEffect } from "react";
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { getSteps } from "@/services/taskApi"
import { useRecoilState } from "recoil";
import { StepsUpdateAtom } from "@/atoms/Update";
import { StepsAtom } from "@/atoms/Tasks";
import { Step } from "@/types/Task";
import { StepEdit } from "./StepEdit";
import { StepDelete } from "./StepDelete";

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
                          <div>
                            <StepEdit taskID={taskID} stepIndex={index}/>
                          </div>
                          <div className="ml-2">
                            <StepDelete taskID={taskID} stepIndex={index}/>
                          </div>
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