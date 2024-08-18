import { atom } from 'recoil';

export const User = atom({
  key: 'User',
  default: {
    id: 0,
    email: '',
    name: '',
    role: '',
  },
});