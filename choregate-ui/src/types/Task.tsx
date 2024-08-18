export type Task = {
    id: string;
    name: string;
};

export type Step = {
    name: string;
    image: string;
    command: string;
    computeResources: string;
};

export type TaskRun = {
    ID: string;
    TaskID: string;
    Status: string;
    CreatedAt: string;
}