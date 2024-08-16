import api from './api';

const createTask = async () => {
  try {
    const data = {
      "name": "mock",
    };

    const response = await api.post('/tasks', data);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const getTask = async (taskID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const getSteps = async (taskID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/steps`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const getTasks = async () => {
  try {
    const response = await api.get('/tasks');
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const runTask = async (taskID: string) => {
  try {
    const response = await api.post(`/tasks/${taskID}/runs`);
    console.log(response);
    return response.data;
  } catch (error) {
    console.error(error);
    return error;
  }
};

const deleteTask = () => {

}

const addSteps = async (taskID: string) => {
  try {
    const data = [{"image": "ubuntu", "command": ["echo", "mock"]}];
    const response = await api.put(`/tasks/${taskID}/steps`, data);
    console.log(response);
    return response.data;
  } catch (error) {
    console.error(error);
    return error;
  }
};

const getTaskRuns = async (taskID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/runs`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

export {
  createTask,
  getTask,
  getTasks,
  runTask,
  deleteTask,
  addSteps,
  getSteps,
  getTaskRuns
};