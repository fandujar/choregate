import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const AuthContext = React.createContext({});

export const useAuth = () => React.useContext(AuthContext);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const navigate = useNavigate();
    const [user, setUser] = React.useState({});

    useEffect(() => {
        if (!user) {
            navigate('/login');
        }
        setUser({ username: 'test' });
    }, []);

    return (
        <AuthContext.Provider value={user}>
            {children}
        </AuthContext.Provider>
    );
}