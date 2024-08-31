import { Button } from '@/components/ui/button';
import { useState } from 'react';
import { useAuth } from '@/hooks/Auth';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Checkbox } from '@/components/ui/checkbox';
import { toast } from 'sonner';

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
                toast.error('email or password is incorrect');
            }
        } catch (error) {
            console.error('Error logging in', error);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <Card className="w-full max-w-md">
            <CardHeader className="space-y-1">
            <div className="flex items-center justify-center mb-4">
                <span className="ml-2 text-2xl font-bold text-pink-700">Choregate</span>
            </div>
            <CardTitle className="text-2xl font-bold text-center">Sign in to your account</CardTitle>
            <CardDescription className="text-center">
                Enter your email below to login to your account
            </CardDescription>
            </CardHeader>
            <form onSubmit={handleLogin}>
            <CardContent className="space-y-4">
            <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input  
                    id="email"
                    placeholder="m@example.com"
                    required type="email"
                    onChange={(e) => setCredentials({ ...credentials, username: e.target.value })}
                    value={credentials.username}
                />
            </div>
            <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input 
                    id="password"
                    required type="password"
                    onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
                    value={credentials.password} 
                />
            </div>
            <div className="flex items-center space-x-2">
                <Checkbox id="remember" />
                <label
                htmlFor="remember"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                >
                Remember me
                </label>
            </div>
            </CardContent>
            <CardFooter className="flex flex-col space-y-4">
            <Button className="w-full bg-pink-700 text-white">Sign In</Button>
            <div className="flex justify-between w-full text-sm">
                <a href="#" className="text-pink-500 hover:underline">
                Forgot password?
                </a>
                <a href="#" className="text-pink-500 hover:underline">
                Create an account
                </a>
            </div>
            </CardFooter>
            </form>
        </Card>
        </div>
    )
}