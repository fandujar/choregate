import { getTaskRuns, runTask } from "@/services/taskApi"
import { Button } from "./ui/button"
import { toast } from "sonner"
import { useQuery } from "react-query"

type TaskRunCreateProps = {
    taskID: string
}

export const TaskRunCreate = (props: TaskRunCreateProps) => {
    const { taskID } = props
    const { refetch } = useQuery('taskRuns', () => getTaskRuns(taskID), {staleTime: 1000})

    const handleTaskRunCreate = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault()
        runTask(taskID).then(() => {
            refetch()
        }).catch((err) => {
            console.log(err)
            toast.error(`${err.message}: ${err.response.data}`)
        })
    }

    return (
        <Button className="bg-pink-700 text-white" onClick={handleTaskRunCreate}>
            Run Task
        </Button>
    )
}