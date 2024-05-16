type TaskRunProps = {
    id: string
}

const TaskRun = (props: TaskRunProps) => {
    const { id } = props
    return (
        <div key={id} style={{ display: "flex", flexDirection: "row", margin:10 }}>
            <p style={{width: 450}}>{id}</p>
        </div>
    )
}

export default TaskRun