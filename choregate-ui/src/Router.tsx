import { Routes, Route } from 'react-router-dom';
import { Home } from './pages/Home';
import { Login } from './pages/Login';

export const Router = () => {

    return (
        <Routes>
            <Route path="/" Component={Home} />
            <Route path="/login" Component={Login} />
        </Routes>
    )
}