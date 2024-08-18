import { Button } from '@/components/ui/button';
import { useState } from 'react';
import { useAuth } from '@/hooks/Auth';

export const LoginPage = () => {
    const [credentials, setCredentials] = useState({ username: '', password: '' });
    const { login } = useAuth();

    const handleLogin = async (e: any) => {
        e.preventDefault();
        try {
            const response = await fetch('/user/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams(credentials),
            });

            if (response.ok) {
                const data = await response.json();
                login(data);
            } else {
                console.error('Login failed');
            }
        } catch (error) {
            console.error('Error logging in', error);
        }
    };

    return (
        <div className='flex justify-center items-center h-full'>
        <div className="p-4">
            <form onSubmit={handleLogin} className='flex flex-col items-center'>
                <h1 className='text-3xl font-bold text-pink-700 flex'>Choregate</h1>
                <div className='mt-10 flex flex-col'>
                <label>
                    E-mail
                </label>
                <input
                    className='rounded'
                    type="text"
                    onChange={(e) => setCredentials({ ...credentials, username: e.target.value })}
                    value={credentials.username}
                    name="username" 
                />
                <label className='mt-2'>
                    Password
                </label>
                <input 
                    className='rounded'
                    type="password" 
                    onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
                    value={credentials.password} 
                    name="password"
                />
                <Button type="submit" className='bg-pink-700 mt-10'>
                    Login
                </Button>
                </div>
            </form>
        </div>
        </div>
    )
}