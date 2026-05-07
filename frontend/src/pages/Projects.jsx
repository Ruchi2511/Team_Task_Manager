import { useEffect, useState } from "react"

import Navbar from "../components/Navbar"
import api from "../services/api"

function Projects() {

  const [projects, setProjects] = useState([])

  const [users, setUsers] = useState([])

  const [selectedUsers, setSelectedUsers] = useState({})

  const [search, setSearch] = useState("")

  const role = localStorage.getItem("role")

  const [formData, setFormData] = useState({
    title: "",
    description: "",
  })

  useEffect(() => {

    fetchProjects()

    if (role === "admin") {
      fetchUsers()
    }

  }, [search])

  const fetchProjects = async () => {

    try {

      const response = await api.get(
        `/projects?title=${search}`
      )

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

      await api.post("/projects", formData)

      alert("Project created successfully")

      setFormData({
        title: "",
        description: "",
      })

      fetchProjects()

    } catch (error) {

      alert(
        error.response?.data?.message || "Failed to create project",
      )
    }
  }

  const addMember = async (projectID) => {

    try {

      await api.post(
        `/projects/${projectID}/members`,
        {
          user_id: selectedUsers[projectID],
        },
      )

      alert("Member added successfully")

    } catch (error) {

      alert(
        error.response?.data?.message || "Failed to add member",
      )
    }
  }

  return (
    <div>

      <Navbar />

      <div className="p-6">

        <h1 className="text-3xl font-bold mb-6">
          Projects
        </h1>

        {role === "admin" && (

          <form
            onSubmit={handleSubmit}
            className="bg-slate-900 p-6 rounded-2xl flex flex-col gap-4 mb-8"
          >

            <input
              type="text"
              name="title"
              placeholder="Project Title"
              value={formData.title}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800 outline-none"
            />

            <textarea
              name="description"
              placeholder="Project Description"
              value={formData.description}
              onChange={handleChange}
              className="p-3 rounded-lg bg-slate-800 outline-none"
            />

            <button className="bg-blue-500 hover:bg-blue-600 p-3 rounded-lg font-semibold">
              Create Project
            </button>

          </form>

        )}

        <input
          type="text"
          placeholder="Search Projects..."
          value={search}
          onChange={(e) =>
            setSearch(e.target.value)
          }
          className="w-full p-3 mb-6 rounded-lg bg-slate-800 outline-none"
        />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">

          {projects.map((project) => (

            <div
              key={project.id}
              className="bg-slate-900 p-6 rounded-2xl border border-slate-800"
            >

              <h2 className="text-2xl font-bold">
                {project.title}
              </h2>

              <p className="mt-4 text-slate-300">
                {project.description}
              </p>

              <div className="mt-5 text-sm text-slate-400">
                Created At
              </div>

              <div className="text-slate-200">
                {new Date(project.created_at).toLocaleDateString()}
              </div>

              {role === "admin" && (

                <div className="mt-6 flex flex-col gap-3">

                  <select
                    value={selectedUsers[project.id] || ""}
                    onChange={(e) =>
                      setSelectedUsers({
                        ...selectedUsers,
                        [project.id]: e.target.value,
                      })
                    }
                    className="p-3 rounded-lg bg-slate-800"
                  >

                    <option value="">
                      Select User
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

                  <button
                    onClick={() =>
                      addMember(project.id)
                    }
                    className="bg-green-500 hover:bg-green-600 p-3 rounded-lg font-semibold"
                  >
                    Add Member
                  </button>

                </div>

              )}

            </div>

          ))}

        </div>

      </div>

    </div>
  )
}

export default Projects