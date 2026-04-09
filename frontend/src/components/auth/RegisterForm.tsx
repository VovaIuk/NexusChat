import { Link, useNavigate } from "react-router-dom";
import { useState, type SubmitEvent } from "react";
import AuthCard from "./AuthCard";
import { register } from "../../api/auth";

export default function RegisterForm() {
  const navigate = useNavigate();
  const [tag, setTag] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [passwordRepeat, setPasswordRepeat] = useState("");
  const [error, setError] = useState("");

  async function handleSubmit(e: SubmitEvent<HTMLFormElement>) {
    e.preventDefault();
    setError("");

    if (password !== passwordRepeat) {
      setError("Пароли не совпадают");
      return;
    }

    try {
      await register(tag, name, password);
      navigate("/login");
    } catch {
      setError("Не удалось зарегистрироваться");
    }
  }

  return (
    <AuthCard title="Регистрация" subtitle="Создайте аккаунт в Народном-чате">
      <form className="auth-form" onSubmit={handleSubmit}>
        <label className="auth-label" htmlFor="register-usertag">
          Usertag
          <input
            className="auth-input"
            id="register-usertag"
            name="usertag"
            type="text"
            value={tag}
            onChange={(e) => setTag(e.target.value)}
            placeholder="new_user"
            autoComplete="username"
            minLength={8}
            required
          />
        </label>

        <label className="auth-label" htmlFor="username">
          Username
          <input
            className="auth-input"
            id="username"
            name="username"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Введите имя пользователя"
            autoComplete="name"
            required
          />
        </label>

        <label className="auth-label" htmlFor="register-password">
          Пароль
          <input
            className="auth-input"
            id="register-password"
            name="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Введите пароль"
            autoComplete="new-password"
            minLength={8}
            required
          />
        </label>

        <label className="auth-label" htmlFor="register-password-repeat">
          Повторите пароль
          <input
            className="auth-input"
            id="register-password-repeat"
            name="passwordRepeat"
            type="password"
            value={passwordRepeat}
            onChange={(e) => setPasswordRepeat(e.target.value)}
            placeholder="Повторите пароль"
            autoComplete="new-password"
            minLength={8}
            required
          />
        </label>

        {error && <p className="auth-error">{error}</p>}

        <button className="auth-button" type="submit">
          Зарегистрироваться
        </button>
      </form>

      <p className="auth-footer">
        Уже есть аккаунт?{" "}
        <Link className="auth-link" to="/login">
          Войти
        </Link>
      </p>
    </AuthCard>
  );
}
