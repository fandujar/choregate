import { addSteps, runTask } from "@/services/taskApi"

type TaskProps = {
    id: string
    setUpdate: (update: boolean) => void
}

const Task = (props: TaskProps) => {
    const { id, setUpdate } = props
    return (
        <div key={id} style={{ display: "flex", flexDirection: "row", margin:10 }}>
            <p style={{width: 450}}>{id}</p>
            <button style={{margin: 5}} onClick={() => {addSteps(id)}}>Add Steps</button>
            <button style={{margin: 5}} onClick={() => {runTask(id); setUpdate(true)}}>Run Task</button>
        </div>
    )
}

export default Task