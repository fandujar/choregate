import { StepType, TaskRunType, TaskType } from "@/types/Task";
import { atom, selector } from "recoil";

export const TasksAtom = atom<TaskType[]>({
    key: 'tasks',
    default: []
})

export const TaskAtom = atom<TaskType>({
    key: 'task',
    default: {
        name: ''
    }
})

export const StepsAtom = atom<StepType[]>({
    key: 'steps',
    default: []
})

export const StepAtom = atom<StepType>({
    key: 'step',
    default: { name: '', image: 'ubuntu', script: 'echo "Hello, World!"' }
})

export const TaskRunsAtom = atom<TaskRunType[]>({
    key: 'taskRuns',
    default: []
})

export const TaskRunAtom = atom<TaskRunType>({
    key: 'taskRun',
    default: {
        ID: '',
        TaskID: ''
    }
})

export const TaskRunLogsAtom = atom({
    key: 'taskRunLogs',
    default: {}
})

export const TaskSelector = selector({
    key: 'taskSelector',
    get: ({get}) => (id: string) => {
        const tasks = get(TasksAtom)
        return tasks.find((task: TaskType) => task.id === id)
    }
})