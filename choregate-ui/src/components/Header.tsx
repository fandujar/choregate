import { ListTodo } from 'lucide-react';
import { Avatar, AvatarFallback, AvatarImage } from './ui/avatar';
import { DropdownMenu, DropdownMenuContent, DropdownMenuTrigger } from './ui/dropdown-menu';
import { useAuth } from '@/hooks/Auth';
// import { useRecoilState } from 'recoil';
// import { UserAtom } from '@/atoms/User';

export const Header = () => {
    // const [user, _] = useRecoilState(UserAtom);
    const { logout } = useAuth();

    const handleLogout = async (e: React.MouseEvent<HTMLAnchorElement>) => {
        e.preventDefault();
        logout();    
    }

    return (
        <div className='flex p-4'>
            <h1 className='text-3xl text-pink-700 font-bold flex'>
                <ListTodo className="mt-2 mr-2"/>
                Choregate
            </h1>

            <DropdownMenu>
                <DropdownMenuTrigger className='ml-auto'>
                    <Avatar className='ml-auto' role="button">
                        <AvatarImage src="https://github.com/fandujar.png" alt="User" />
                        <AvatarFallback>US</AvatarFallback>
                    </Avatar>
                </DropdownMenuTrigger>
                <DropdownMenuContent className='absolute right-0'>
                    <div className='p-2 bg-white rounded shadow-md'>
                        <a href='/profile' className='block'>Profile</a>
                    </div>
                    <div className='p-2 bg-white rounded shadow-md'>
                        <a onClick={handleLogout} className='block' role="button">Logout</a>
                    </div>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}