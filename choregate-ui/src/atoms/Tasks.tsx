import { StepType, TaskRunType, TaskType } from "@/types/Task";
import { atom, atomFamily, selector } from "recoil";

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

export const StepsAtom = atomFamily<StepType[], string>({
    key: 'steps',
    default: []
})

export const StepAtom = atomFamily<StepType, string>({
    key: 'step',
    default: { name: '', image: 'ubuntu', script: 'echo "Hello, World!"' }
})

export const TaskRunsAtom = atomFamily<TaskRunType[], string>({
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