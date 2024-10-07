import { TaskRunsUpdateAtom } from "@/atoms/Update"
import { runTask } from "@/services/taskApi"
import { useRecoilState } from "recoil"
import { Button } from "./ui/button"
import { toast } from "sonner"

type TaskRunCreateProps = {
    taskID: string
}

export const TaskRunCreate = (props: TaskRunCreateProps) => {
    const { taskID } = props
    const [, setUpdate] = useRecoilState(TaskRunsUpdateAtom)

    const handleTaskRunCreate = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault()
        runTask(taskID).then(() => {
            setUpdate(true)
        }).catch((err) => {
            console.log(err)
            toast.error(`${err.message}: ${err.response.data}`)
            setUpdate(true)
        })
    }

    return (
        <Button className="bg-pink-700 text-white" onClick={handleTaskRunCreate}>
            Run Task
        </Button>
    )
}