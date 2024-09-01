# Choregate

Choregate is a task manager to help you handle and automate your tasks. It has a intuitive interface to manage TektonCD tasks.

## Build and Run

Make sure you have Docker, Docker Compose, Go, Kind and Node.js installed.

### Setup

First, clone the repository:

```sh
git clone https://github.com/fandujar/choregate.git
cd choregate
```

Then, set up the environment. This will create a local Kubernetes cluster using `kind` and start all the services with Docker Compose:

```sh
make up
```

### Accessing Choregate

Once everything is running, go to login page `http://localhost:8080`.

### Using Choregate

Open your browser and go to `http://localhost:8080`. Log in with the default credentials:

- **Email**: `email@admin.com`
- **Password**: `password`

After logging in, you can create tasks in the "Tasks" section by clicking "Create Task". To run a task, click the "Run Task" button next to it. You can check the progress and logs of task runs in the "Runs" section of each task.

### Stopping the Services

To stop the services, run:

```sh
make down
```

## License

Choregate is licensed under the [MIT License](LICENSE).
