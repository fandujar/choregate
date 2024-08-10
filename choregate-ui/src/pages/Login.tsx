export const Login = () => {
    return (
        <div className="bg-slate-200 max-w-30 h-40 p-4">
            <form action="/user/login" method="post">
                <label>Username</label>
                <input type="text" name="username" />
                <label>Password</label>
                <input type="password" name="password"/>
                
                <button type="submit">Login</button>
            </form>
        </div>
    )
}