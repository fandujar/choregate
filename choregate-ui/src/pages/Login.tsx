import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export const Login = () => {
    const [credentials, setCredentials] = useState({ username: '', password: '' });
    const navigate = useNavigate();

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
                localStorage.setItem('jwt', data.token);
                navigate('/');
            } else {
                console.error('Login failed');
            }
        } catch (error) {
            console.error('Error logging in', error);
        }
    };

    return (
        <div className="bg-slate-200 max-w-30 h-40 p-4">
            <form onSubmit={handleLogin}>
                <label>Username</label>
                <input
                    type="text"
                    onChange={(e) => setCredentials({ ...credentials, username: e.target.value })}
                    value={credentials.username}
                    name="username" 
                />
                <label>Password</label>
                <input 
                    type="password" 
                    onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
                    value={credentials.password} 
                    name="password"
                />
                
                <button type="submit">Login</button>
            </form>
        </div>
    )
}