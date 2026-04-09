import { Link, useNavigate } from "react-router-dom";
import { useState, type SubmitEvent } from 'react'
import AuthCard from "./AuthCard";

import {login} from "../../api/auth"
import {useUser} from "../../contexts/UserContext"

export default function LoginForm() {
  console.log("Satrt login form");

  const navigate = useNavigate();
  const [tag, setTag] = useState("");
  const [password, setPassword] = useState("");
  const {setUser} = useUser();
  const [error, setError] = useState("");

  async function handleSubmit(e: SubmitEvent<HTMLFormElement>){
    console.log("press login button");
    e.preventDefault();
    setError("")

    try {
      const data = await login(tag, password);
      localStorage.setItem("token", data.token.refresh);
      setUser(data.user);
      navigate("/");
    } catch{
      setError("Неверный логин или пароль");
    }
  }

  return (
    <AuthCard title="Вход" subtitle="Войдите в аккаунт NexusChat">
      <form className="auth-form" onSubmit={handleSubmit}>
        <label className="auth-label" htmlFor="usertag">
          Usertag
          <input
            className="auth-input"
            id="usertag"
            name="usertag"
            type="text"
            value={tag}
            onChange={(e)=>setTag(e.target.value)}
            placeholder="user"
            autoComplete="username"
            required
          />
        </label>

        <label className="auth-label" htmlFor="password">
          Пароль
          <input
            className="auth-input"
            id="password"
            name="password"
            type="password"
            value={password}
            onChange={(e)=>setPassword(e.target.value)}
            placeholder="Введите пароль"
            autoComplete="current-password"
            required
          />
        </label>

        {error && <p className="auth-error">{error}</p>}
        <button className="auth-button" type="submit">
          Войти
        </button>
      </form>

      <p className="auth-footer">
        Нет аккаунта?{" "}
        <Link className="auth-link" to="/auth">
          Зарегистрироваться
        </Link>
      </p>
    </AuthCard>
  );
}
