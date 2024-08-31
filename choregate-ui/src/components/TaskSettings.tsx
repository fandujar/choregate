import { Button } from "./ui/button"

type TaskSettingsProps = {
    taskID: string
}

export const TaskSettings = (_: TaskSettingsProps) => {
    return (
        <Button className="bg-pink-700 text-white">
            Settings
        </Button>    
    )
}