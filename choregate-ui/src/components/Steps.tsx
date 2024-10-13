import { useEffect } from "react";
import { Card, CardContent } from './ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from './ui/table'
import { getSteps } from "@/services/taskApi"
import { useRecoilState } from "recoil";
import { StepsAtom } from "@/atoms/Tasks";
import { StepType } from "@/types/Task";
import { StepEdit } from "./StepEdit";
import { StepDelete } from "./StepDelete";
import { useQuery } from "react-query";

type StepsProps = {
  taskID: string;
};

export const Steps = (props: StepsProps) => {
    const { taskID } = props;
    const [steps, setSteps] = useRecoilState<StepType[]>(StepsAtom(taskID));
    const { data, isLoading } = useQuery('steps', () => getSteps(taskID), {staleTime: 1000});

    useEffect(() => {
        if (data) {
            setSteps(data);
        }
    }, [data]);

    if (isLoading) {
        return <div>Loading...</div>
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
              {steps?.map((step: StepType, index) => (
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