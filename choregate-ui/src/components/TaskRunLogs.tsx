import { TaskRunLogsAtom } from "@/atoms/Tasks"
import { getTaskRunLogs } from "@/services/taskApi"
import { useRecoilState } from "recoil"
import { toast } from "sonner"
import { Sheet, SheetClose, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger } from './ui/sheet';
import { ScrollArea } from './ui/scroll-area';
import { Tabs, TabsList, TabsTrigger } from './ui/tabs';
import { TabsContent } from '@radix-ui/react-tabs';
import { Button } from "./ui/button";

type TaskRunLogsProps = {
    taskID: string
    taskRunID: string
}

export const TaskRunLogs = (props: TaskRunLogsProps) => {
    const { taskID, taskRunID } = props
    const [taskRunLogs, setTaskRunLogs] = useRecoilState<TaskRunLogsType>(TaskRunLogsAtom)

    const handleViewLogs = () => {
        let response = getTaskRunLogs(taskID, taskRunID)
        response.then((logs) => {
            setTaskRunLogs(logs)
        }).catch((err) => {
            toast.error(`${err.message}: ${err.response.data}`)
        })
    }

    return (
        <Sheet>
            <SheetTrigger>
                <Button onClick={handleViewLogs} variant={"ghost"}>View Logs</Button>
            </SheetTrigger>
            <SheetContent className="w-[540px] sm:w-[940px] sm:max-w-[940px]">
                <SheetHeader>
                    <SheetTitle>Task Run Logs</SheetTitle>
                    <SheetDescription>logs from the task execution</SheetDescription>
                </SheetHeader>
                <ScrollArea className='h-full w-full'>
                    <Tabs className="p-5">
                        <TabsList className="grid w-full grid-cols-5">
                            {Object.keys(taskRunLogs)?.map((key, index) => (
                                <TabsTrigger key={index} value={key}>{key}</TabsTrigger>
                            ))}
                        </TabsList>
                        {Object.keys(taskRunLogs)?.map((key, index) => (
                            <TabsContent key={index} value={key}>
                                <pre>{taskRunLogs[key]}</pre>
                            </TabsContent>
                        ))}
                    </Tabs>
                </ScrollArea>
                <SheetFooter>
                    <SheetClose asChild>
                        <Button type="submit">Close</Button>
                    </SheetClose>
                </SheetFooter>
            </SheetContent>
        </Sheet>
    )
}