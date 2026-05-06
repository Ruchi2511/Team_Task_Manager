import { useNavigate } from "react-router-dom"

function Navbar() {

  const navigate = useNavigate()

  const role = localStorage.getItem("role")

  const name = localStorage.getItem("name")

  const handleLogout = () => {

    localStorage.clear()

    navigate("/")
  }

  return (
    <div className="flex justify-between items-center p-4 bg-slate-900 border-b border-slate-700">

      <div>

        <h1 className="text-2xl font-bold">
          Team Task Manager
        </h1>

        <p className="text-sm text-slate-400 mt-1">
          Welcome, {name} ({role})
        </p>

      </div>

      <div className="flex gap-4">

        <button
          onClick={() => navigate("/dashboard")}
          className="bg-slate-700 hover:bg-slate-600 px-4 py-2 rounded-lg"
        >
          Dashboard
        </button>

        <button
          onClick={() => navigate("/projects")}
          className="bg-slate-700 hover:bg-slate-600 px-4 py-2 rounded-lg"
        >
          Projects
        </button>

        <button
          onClick={() => navigate("/tasks")}
          className="bg-slate-700 hover:bg-slate-600 px-4 py-2 rounded-lg"
        >
          Tasks
        </button>

        <button
          onClick={handleLogout}
          className="bg-red-500 hover:bg-red-600 px-4 py-2 rounded-lg"
        >
          Logout
        </button>

      </div>

    </div>
  )
}

export default Navbar