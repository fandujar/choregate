import { atom } from 'recoil';

export const TaskUpdateAtom = atom({
    key: 'taskUpdate',
    default: false
})

export const TasksUpdateAtom = atom({
    key: 'tasksUpdate',
    default: false
})