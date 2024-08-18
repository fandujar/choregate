import { createContext, useContext, useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { User } from '@/atoms/User';

const AuthContext = createContext({});

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [user, setUser] = useRecoilState(User);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('jwt');
        if (!token) {
            navigate('/login');
        } else {
            navigate('/');
        }

    }, [setUser]);

    const login = (data: any) => {
        localStorage.setItem('jwt', data.token);
        navigate('/');
    }

    const logout = () => {
        localStorage.removeItem('jwt');
        navigate('/login');
    }

    const value = useMemo(() => ({ user, login, logout }), [user]);

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
}

export const useAuth = () => useContext(AuthContext);