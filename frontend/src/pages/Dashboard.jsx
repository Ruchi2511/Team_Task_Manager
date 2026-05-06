import { useEffect, useState } from "react"

import Navbar from "../components/Navbar"
import api from "../services/api"

function Dashboard() {

  const [stats, setStats] = useState({})

  useEffect(() => {
    fetchStats()
  }, [])

  const fetchStats = async () => {

    try {

      const response = await api.get("/dashboard/stats")

      setStats(response.data.stats)

    } catch (error) {

      console.log(error)
    }
  }

  return (
    <div>

      <Navbar />

      <div className="p-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">

        <div className="bg-slate-900 p-6 rounded-2xl">
          <h2 className="text-xl">
            Total Projects
          </h2>

          <p className="text-4xl font-bold mt-4">
            {stats.total_projects || 0}
          </p>
        </div>

        <div className="bg-slate-900 p-6 rounded-2xl">
          <h2 className="text-xl">
            Total Tasks
          </h2>

          <p className="text-4xl font-bold mt-4">
            {stats.total_tasks || 0}
          </p>
        </div>

        <div className="bg-slate-900 p-6 rounded-2xl">
          <h2 className="text-xl">
            Completed Tasks
          </h2>

          <p className="text-4xl font-bold mt-4">
            {stats.completed_tasks || 0}
          </p>
        </div>

        <div className="bg-slate-900 p-6 rounded-2xl">
          <h2 className="text-xl">
            Overdue Tasks
          </h2>

          <p className="text-4xl font-bold mt-4">
            {stats.overdue_tasks || 0}
          </p>
        </div>

      </div>

    </div>
  )
}

export default Dashboard