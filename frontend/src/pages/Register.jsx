import { useState } from "react"
import { useNavigate, Link } from "react-router-dom"

import api from "../services/api"

function Register() {

  const navigate = useNavigate()

  const [formData, setFormData] = useState({
    name: "",
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

      await api.post("/register", formData)

      alert("Registration successful")

      navigate("/")

    } catch (error) {

      alert(
        error.response?.data?.message || "Registration failed",
      )
    }
  }

  return (
    <div className="min-h-screen flex justify-center items-center bg-slate-950">

      <form
        onSubmit={handleSubmit}
        className="bg-slate-900 p-8 rounded-2xl w-[400px] flex flex-col gap-4"
      >

        <h1 className="text-3xl font-bold text-center">
          Register
        </h1>

        <input
          type="text"
          name="name"
          placeholder="Name"
          onChange={handleChange}
          className="p-3 rounded-lg bg-slate-800 outline-none"
        />

        <input
          type="email"
          name="email"
          placeholder="Email"
          onChange={handleChange}
          className="p-3 rounded-lg bg-slate-800 outline-none"
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          onChange={handleChange}
          className="p-3 rounded-lg bg-slate-800 outline-none"
        />

        <button className="bg-green-500 p-3 rounded-lg font-semibold hover:bg-green-600">
          Register
        </button>

        <Link
          to="/"
          className="text-center text-blue-400"
        >
          Already have an account?
        </Link>

      </form>

    </div>
  )
}

export default Register