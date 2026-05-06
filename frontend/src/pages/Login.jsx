import { useState } from "react"
import { useNavigate, Link } from "react-router-dom"

import api from "../services/api"

function Login() {

  const navigate = useNavigate()

  const [formData, setFormData] = useState({
    email: "",
    password: "",
  })

  const handleChange = (e) => {

    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    })
  }

  const handleSubmit = async (e) => {

    e.preventDefault()

    try {

      const response = await api.post("/login", formData)

      localStorage.setItem(
        "token",
        response.data.access_token,
      )

      localStorage.setItem(
        "role",
        response.data.data.role,
      )

      localStorage.setItem(
        "name",
        response.data.data.name,
      )

      navigate("/dashboard")

    } catch (error) {

      alert(
        error.response?.data?.message || "Login failed",
      )
    }
  }

  return (
    <div className="min-h-screen flex justify-center items-center bg-slate-950">

      <form
        onSubmit={handleSubmit}
        className="bg-slate-900 p-8 rounded-2xl w-[400px] flex flex-col gap-4 shadow-2xl"
      >

        <h1 className="text-4xl font-bold text-center">
          Login
        </h1>

        <input
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          className="p-3 rounded-lg bg-slate-800 outline-none border border-slate-700"
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          className="p-3 rounded-lg bg-slate-800 outline-none border border-slate-700"
        />

        <button
          className="bg-blue-500 hover:bg-blue-600 p-3 rounded-lg font-semibold transition"
        >
          Login
        </button>

        <Link
          to="/register"
          className="text-center text-blue-400 hover:text-blue-300"
        >
          Create Account
        </Link>

      </form>

    </div>
  )
}

export default Login