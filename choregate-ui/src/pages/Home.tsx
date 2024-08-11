import TaskList from '../components/TaskList'
import UserList from '../components/UserList'


export const Home = () => {
  return (
        <>
        <div className='bg-slate-950 text-slate-300 flex flex-col p-2'>
        <h1 className='text-2xl text-pink-700 font-bold'>Choregate</h1>
        <div className="flex">
          <div className="column" style={{ flex: 1 }}>
            <UserList />
          </div>
          <div className="column" style={{ flex: 1 }}>
            <TaskList />
          </div>
        </div>
        </div>
        </>
    )
}