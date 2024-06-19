import { useEffect, useState } from "react";
import { getUsers, createUser } from "../services/userApi";

type User = {
    id: string,
    slug: string,
    name: string,
    email: string,
    password: string,
}

interface UserListProps {}

const UserList = (_: UserListProps) => {
    
    const [update, setUpdate] = useState<boolean>(true);
    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        const users = getUsers();
        users.then((users) => {
            setUsers(users);
        })
        setUpdate(false);
    }, [update])

    return (
        <div>
            <h1>User List</h1>
            <button onClick={() => {
                const random = Math.floor(Math.random() * 1000);
                const slug = `mockslug${random}`;
                const name = `mockname${random}`;
                const email = `mockmail${random}`;
                const password = "mockpassword";
                createUser(name, slug, email, password);
                setUpdate(true);
            }}
            >
                Create User
            </button>
            <ul>
                {users?.map((user: User) => (
                    <li key={user.id}>
                        {user.id} - {user.name}
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default UserList