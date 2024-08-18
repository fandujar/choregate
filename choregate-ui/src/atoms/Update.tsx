import { atom } from 'recoil';

export const TaskUpdateAtom = atom({
    key: 'taskUpdate',
    default: false
})

export const StepsUpdateAtom = atom({
    key: 'stepsUpdate',
    default: false
})

export const TasksUpdateAtom = atom({
    key: 'tasksUpdate',
    default: false
})

export const TaskRunsUpdateAtom = atom({
    key: 'taskRunsUpdate',
    default: false
})