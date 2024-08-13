import { Routes, Route } from 'react-router-dom';
import { DashboardPage } from './pages/Dashboard';
import { LoginPage } from './pages/Login';
import { TasksPage } from './pages/Tasks';
import { AdminPage } from './pages/Admin';

export const Router = () => {

    return (
        <Routes>
            <Route path="/" Component={DashboardPage} />
            <Route path="/tasks" Component={TasksPage} />
            <Route path="/admin" Component={AdminPage} />
            <Route path="/login" Component={LoginPage} />
        </Routes>
    )
}