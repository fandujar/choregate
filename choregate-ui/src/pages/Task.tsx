import { Header } from "@/components/Header"
import { Menu } from "@/components/Menu"
import { Task } from "@/components/Task"

// FIX: use react-router-dom way to get taskID
export const TaskPage = ({taskID}) => {
    console.log(taskID)
    return (
        <div>
            <Header/>
            <div className="flex">
                <Menu />
                <Task id={taskID}/>
            </div>
        </div>
    )
}