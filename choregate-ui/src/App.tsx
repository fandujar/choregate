import TaskList from './components/TaskList'
import UserList from './components/UserList'
import './App.css'

function App() {
  return (
    <>
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

export default App
