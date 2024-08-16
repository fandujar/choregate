import { Card, CardHeader, CardTitle, CardContent } from './ui/card';

type StepsProps = {
    step: Step;
}

type Step = {
    name: string;
    image: string;
    command: string[];
    computeResources: string;
}

export const Step = (props: StepsProps) => {
    const { step } = props;
    return (
    <Card className="mb-4">
      <CardHeader>
        <CardTitle>{step.name || "Unnamed Step"}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex items-center space-x-4">
          <div>
            <p><strong>Image:</strong> {step.image}</p>
            <p><strong>Command:</strong> {step.command.join(' ')}</p>
            <p><strong>Compute Resources:</strong> {JSON.stringify(step.computeResources) || "None"}</p>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};