import { createContext, useContext, useEffect, useMemo } from 'react';
import { useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { UserAtom } from '@/atoms/User';
import { UserType } from '@/types/User';

type AuthContextType = {
    user: any;
    login: (data: any) => void;
    logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
    user: null,
    login: () => null,
    logout: () => null,
});

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [user, setUser] = useRecoilState<UserType>(UserAtom);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('jwt');
        if (!token) {
            navigate('/login');
        } else {
            navigate('/');
        }

    }, [setUser]);

    useEffect(() => {
        const interval = setInterval(() => {
            const token = localStorage.getItem('jwt');
            if (!token) {
                navigate('/login');
            } else {
                const response = fetch('/user/validate', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                });
                
                response.then((res) => {
                    if (res.status === 401) {
                        localStorage.removeItem('jwt');
                        navigate('/login');
                    }}).catch((error) => {
                        console.error('Error validating token', error);
                    });
                }
            }, 60000);

        return () => clearInterval(interval);
    }, []);
    


    const login = (data: any) => {
        localStorage.setItem('jwt', data.token);
        setUser({
            id: data.user_id,
            username: data.username,
            email: data.email,
            systemRole: data.system_role
        });
        navigate('/');
    }

    const logout = () => {
        localStorage.removeItem('jwt');
        setUser({
            id: '',
            username: '',
            email: '',
            systemRole: '',
        });
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