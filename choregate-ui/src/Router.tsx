import { Routes, Route } from 'react-router-dom';
import { DashboardPage } from './pages/Dashboard';
import { LoginPage } from './pages/Login';
import { TasksPage } from './pages/Tasks';
import { TaskPage } from './pages/Task';
import { AdminPage } from './pages/Admin';

export const Router = () => {

    return (
        <Routes>
            <Route path="/" Component={DashboardPage} />
            <Route path="/admin" Component={AdminPage} />
            <Route path="/login" Component={LoginPage} />
            <Route path="/tasks" Component={TasksPage} />
            <Route path='/tasks/:taskID' Component={TaskPage} />
        </Routes>
    )
}