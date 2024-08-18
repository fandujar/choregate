import { atom, selector } from "recoil";

export const TasksAtom = atom({
    key: 'tasks',
    default: []
})

export const TaskAtom = atom({
    key: 'task',
    default: {}
})

export const StepsAtom = atom({
    key: 'steps',
    default: []
})

export const StepAtom = atom({
    key: 'step',
    default: {}
})

export const TaskRunsAtom = atom({
    key: 'taskRuns',
    default: []
})

export const TaskRunAtom = atom({
    key: 'taskRun',
    default: {}
})

export const TaskRunLogsAtom = atom({
    key: 'taskRunLogs',
    default: ''
})

export const TaskSelector = selector({
    key: 'taskSelector',
    get: ({get}) => (id: string) => {
        const tasks = get(TasksAtom)
        return tasks.find((task: any) => task.id === id)
    }
})