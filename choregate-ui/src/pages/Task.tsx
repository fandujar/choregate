import { useParams } from "react-router-dom"
import { Header } from "@/components/Header"
import { Menu } from "@/components/Menu"
import { Task } from "@/components/Task"

export const TaskPage = () => {
    const { taskID } = useParams();
    
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