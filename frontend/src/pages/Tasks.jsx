import { useEffect, useState } from "react"

import Navbar from "../components/Navbar"
import api from "../services/api"

function Tasks() {

  const [tasks, setTasks] = useState([])

  const [projects, setProjects] = useState([])

  const [users, setUsers] = useState([])

  const role = localStorage.getItem("role")

  const [formData, setFormData] = useState({
    project_id: "",
    title: "",
    description: "",
    assigned_to: "",
    priority: "medium",
    due_date: "",
  })

  useEffect(() => {

    fetchTasks()

    if (role === "admin") {
      fetchProjects()
      fetchUsers()
    }

  }, [])

  const fetchTasks = async () => {

    try {

      const response = await api.get("/tasks")

      setTasks(response.data.tasks)

    } catch (error) {

      console.log(error)
    }
  }

  const fetchProjects = async () => {

    try {

      const response = await api.get("/projects")

      setProjects(response.data.projects)

    } catch (error) {

      console.log(error)
    }
  }

  const fetchUsers = async () => {

    try {

      const response = await api.get("/users")

      setUsers(response.data.users)

    } catch (error) {

      console.log(error)
    }
  }

  const handleChange = (e) => {

    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    })
  }

  const handleSubmit = async (e) => {

    e.preventDefault()

    try {

      await api.post("/tasks", formData)

      alert("Task created successfully")

      setFormData({
        project_id: "",
        title: "",
        description: "",
        assigned_to: "",
        priority: "medium",
        due_date: "",
      })

      fetchTasks()

    } catch (error) {

      alert(
        error.response?.data?.message || "Failed to create task",
      )
    }
  }

  const updateStatus = async (id, status) => {

    try {

      await api.patch(`/tasks/${id}/status`, {
        status,
      })

      fetchTasks()

    } catch (error) {

      alert(
        error.response?.data?.message || "Failed to update status",
      )
    }
  }

  return (
    <div>

      <Navbar />

      <div className="p-6">

        <h1 className="text-3xl font-bold mb-6">
          Tasks
        </h1>

        {role === "admin" && (

          <form
            onSubmit={handleSubmit}
            className="bg-slate-900 p-6 rounded-2xl grid grid-cols-1 md:grid-cols-2 gap-4"
          >

            <select
              name="project_id"
              value={formData.project_id}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800"
            >

              <option value="">
                Select Project
              </option>

              {projects.map((project) => (

                <option
                  key={project.id}
                  value={project.id}
                >
                  {project.title}
                </option>

              ))}

            </select>

            <input
              type="text"
              name="title"
              placeholder="Task Title"
              value={formData.title}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800"
            />

            <textarea
              name="description"
              placeholder="Description"
              value={formData.description}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800 md:col-span-2"
            />

            <select
              name="assigned_to"
              value={formData.assigned_to}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800"
            >

              <option value="">
                Assign User
              </option>

              {users.map((user) => (

                <option
                  key={user.id}
                  value={user.id}
                >
                  {user.name}
                </option>

              ))}

            </select>

            <select
              name="priority"
              value={formData.priority}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800"
            >

              <option value="low">
                Low
              </option>

              <option value="medium">
                Medium
              </option>

              <option value="high">
                High
              </option>

            </select>

            <input
              type="date"
              name="due_date"
              value={formData.due_date}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800"
            />

            <button className="bg-blue-500 hover:bg-blue-600 p-3 rounded-lg font-semibold">
              Create Task
            </button>

          </form>

        )}

      </div>

      <div className="p-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">

        {tasks.map((task) => (

          <div
            key={task.id}
            className="bg-slate-900 p-6 rounded-2xl border border-slate-800"
          >

            <div className="flex justify-between items-center">

              <h2 className="text-2xl font-bold">
                {task.title}
              </h2>

              <span className="bg-slate-800 px-3 py-1 rounded-lg text-sm">
                {task.priority}
              </span>

            </div>

            <p className="mt-4 text-slate-300">
              {task.description}
            </p>

            <div className="mt-5">

              <div className="text-sm text-slate-400">
                Status
              </div>

              <div className="mt-1">
                {task.status}
              </div>

            </div>

            <div className="mt-4">

              <div className="text-sm text-slate-400">
                Due Date
              </div>

              <div className="mt-1">
                {new Date(task.due_date).toLocaleDateString()}
              </div>

            </div>

            {task.status !== "completed" && (

              <button
                onClick={() =>
                  updateStatus(task.id, "completed")
                }
                className="mt-6 w-full bg-green-500 hover:bg-green-600 p-3 rounded-lg font-semibold"
              >
                Mark Completed
              </button>

            )}

          </div>

        ))}

      </div>

    </div>
  )
}

export default Tasks