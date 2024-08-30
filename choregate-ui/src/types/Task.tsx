export type TaskType = {
    id?: string;
    name: string;
};

export type StepType = {
    name: string;
    image: string;
    script: string;
    command?: string;
    computeResources?: string;
};

export type TaskRunType = {
    ID: string;
    TaskID: string;
    Status?: string;
    CreatedAt?: string;
}

export type TaskRunLogsType = {
    [key: string]: string;
}