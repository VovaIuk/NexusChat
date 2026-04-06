import { Link } from "react-router-dom";
import AuthCard from "./AuthCard";

export default function LoginForm() {
  return (
    <AuthCard title="Вход" subtitle="Войдите в аккаунт NexusChat">
      <form className="auth-form">
        <label className="auth-label" htmlFor="usertag">
          Usertag
          <input
            className="auth-input"
            id="usertag"
            name="usertag"
            type="text"
            placeholder="@user"
            autoComplete="username"
          />
        </label>

        <label className="auth-label" htmlFor="password">
          Пароль
          <input
            className="auth-input"
            id="password"
            name="password"
            type="password"
            placeholder="Введите пароль"
            autoComplete="current-password"
          />
        </label>

        <button className="auth-button" type="button">
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
