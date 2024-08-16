import { ListTodo } from 'lucide-react';
import { Avatar, AvatarFallback, AvatarImage } from './ui/avatar';

export const Header = () => {
    return (
        <div className='flex p-2 bg-slate-100'>
            <h1 className='text-3xl text-pink-700 font-bold flex'>
                <ListTodo className="mt-2 mr-2"/>
                Choregate
            </h1>

            <Avatar className='ml-auto'>
                <AvatarImage src="https://github.com/fandujar.png" alt="User" />
                <AvatarFallback>US</AvatarFallback>
            </Avatar>
        </div>
    )
}