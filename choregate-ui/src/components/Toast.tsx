type ToastProps = {
    message: string
    type: string
}

const Toast = ({ message, type }: ToastProps) => {
    return (
        <div className={`toast ${type}`}>
            {message}
        </div>
    )
}

export default Toast