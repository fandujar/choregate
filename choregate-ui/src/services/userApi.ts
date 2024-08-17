import api from './api';

const getUsers = async () => {
  try {
    const response = await api.get('/users');
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const createUser = async (name:string, slug: string, email: string, password: string) => {
    try {
        const data = {
            name: name,
            slug: slug,
            email: email,
            password: password,
        };
        const response = await api.post('/users', data);
        return response.data;
    } catch (error) {
        console.error(error);
    }
}

const getUser = async (username: string) => {
    try {
        const response = await api.get(`/users/${username}`);
        return response.data;
    } catch (error) {
        console.error(error);
    }
}

export {
    getUsers,
    createUser,
    getUser
};