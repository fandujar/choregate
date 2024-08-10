import { useNavigate } from 'react-router-dom'
import { useEffect } from 'react'
import TaskList from '../components/TaskList'
import UserList from '../components/UserList'

export const Home = () => {
  const navigate = useNavigate();
  
  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) {
      navigate('/login')
    }
  }, [navigate])

    return (
        <>
        <h1>Choregate</h1>
        <div className="container" style={{ height: '100vh', width: '200vh', display: 'flex', flexDirection: 'row' }}>
          <div className="column" style={{ flex: 1 }}>
            <UserList />
          </div>
          <div className="column" style={{ flex: 1 }}>
            <TaskList />
          </div>
        </div>
        </>
    )
}