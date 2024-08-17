import { useEffect, useState } from "react";
import { getSteps } from "@/services/taskApi";
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'

type StepsProps = {
  taskID: string;
  update: boolean;
  setUpdate: Function;
};

type Step = {
  name: string;
  image: string;
  command: string[];
  computeResources: string;
}

export const Steps = (props: StepsProps) => {
    const { taskID } = props;
    const [steps, setSteps] = useState([]);
    const { update, setUpdate } = props;

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
                      <TableCell>{step.command.join(' ')}</TableCell>
                  </TableRow>
              ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    );
  };
  