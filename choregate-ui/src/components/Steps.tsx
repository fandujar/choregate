import { useEffect, useState } from "react";
import { Step } from "./Step";
import { getSteps } from "@/services/taskApi";

type StepsProps = {
  taskID: string;
};

export const Steps = (props: StepsProps) => {
    const { taskID } = props;
    const [steps, setSteps] = useState([]);

    useEffect(() => {
      const steps = getSteps(taskID);
      steps.then((steps) => {
        setSteps(steps);
      });
    }, []);

    return (
      <div className="space-y-4">
        {steps?.map((step, index) => (
          <Step key={index} step={step} />
        ))}
      </div>
    );
  };
  