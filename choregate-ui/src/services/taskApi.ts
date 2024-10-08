import api from './api';
import { AxiosError } from 'axios';

const createTask = async (data: any) => {
  try {
    const response = await api.post('/tasks', data);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

const getTask = async (taskID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

const getSteps = async (taskID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/steps`);
    return response.data;
  } catch (error: any) {
    if (error.response.status === 404) {
      return [];
    }
    throw error;
  }
}

const getTasks = async () => {
  try {
    const response = await api.get('/tasks');
    return response.data;
  } catch (error: any) {
    if (error.response.status === 404) {
      return [];
    }
    console.error(error);
    throw error;
  }
}

const runTask = async (taskID: string) => {
  try {
    const response = await api.post(`/tasks/${taskID}/runs`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

const deleteTask = () => {

}

const updateSteps = async (taskID: string, data: any) => {
  try {
    const response = await api.put(`/tasks/${taskID}/steps`, data);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
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

const getTaskRun = async (taskID: string, runID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/runs/${runID}`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

const getTaskRunLogs = async (taskID: string, runID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/runs/${runID}/logs`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

const getTaskRunStatus = async (taskID: string, runID: string) => {
  try {
    const response = await api.get(`/tasks/${taskID}/runs/${runID}/status`);
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
  updateSteps,
  getSteps,
  getTaskRuns,
  getTaskRun,
  getTaskRunLogs,
  getTaskRunStatus
};