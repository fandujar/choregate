import { UserType } from '@/types/User';
import { atom } from 'recoil';

export const UserAtom = atom<UserType>({
  key: 'User',
  default: {
    id: '',
    username: '',
    email: '',
    systemRole: '',
  },
});