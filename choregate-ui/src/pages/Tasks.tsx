import { Header } from "@/components/Header"
import { Menu } from "@/components/Menu"
import TaskList from "@/components/TaskList"

export const TasksPage = () => {
    return (
        <div>
            <Header/>
            <div className="flex">
                <Menu />
                <TaskList />
            </div>
        </div>
    )
}