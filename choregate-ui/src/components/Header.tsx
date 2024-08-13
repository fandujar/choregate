import { ListTodo } from 'lucide-react';

export const Header = () => {
    return (
        <div className='flex flex-col p-2'>
            <h1 className='text-3xl text-pink-700 font-bold flex'>
                <ListTodo className="mt-2 mr-2"/>
                Choregate
            </h1>
        </div>
    )
}