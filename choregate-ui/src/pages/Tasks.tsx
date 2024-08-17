import { Header } from "@/components/Header"
import { Menu } from "@/components/Menu"
import { Tasks } from "@/components/Tasks"

export const TasksPage = () => {
    return (
        <div>
            <Header/>
            <div className="flex">
                <Menu />
                <Tasks />
            </div>
        </div>
    )
}