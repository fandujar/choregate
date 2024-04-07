import { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { faBell, faGear } from '@fortawesome/free-solid-svg-icons';

type User = {
    name: string;
    email: string;
};

const getUser = async (): Promise<User> => {
    const response = new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve({
                name: 'John Doe',
                email: 'johndoe@example.com'
            });
        }, 2000);
    });
    return response as Promise<User>;
}

const Notifications = () => {
    return (
        <>
            <FontAwesomeIcon icon={faBell} />
        </>
    );
}

const Settings = () => {
    return (
        <>
            <FontAwesomeIcon icon={faGear} />
        </>
    );
}

const UserMenu = ({ user, loading }: { user: User | null, loading: boolean }) => {
    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            {user ? (
                <div>
                    <p>{user.name}</p>
                    <p>{user.email}</p>
                </div>
            ) : (
                <div>
                    <p>Not logged in</p>
                </div>
            )}
        </div>
    );
}

const UserHelperBar = () => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const user = await getUser();
                setUser(user);
            } catch (e) {
                console.error(e);
            } finally {
                setLoading(false);
            }
        };

        fetchUser();
    }, []);

    return (
        <>
            <Notifications />
            <Settings />
            <UserMenu user={user} loading={loading} />
        </>
    );

}

export { 
    UserHelperBar,
};