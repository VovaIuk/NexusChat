import { Link } from "react-router-dom";
import AuthCard from "./AuthCard";

export default function RegisterForm() {
  return (
    <AuthCard title="Регистрация" subtitle="Создайте аккаунт в NexusChat">
      <form className="auth-form">
        <label className="auth-label" htmlFor="register-usertag">
          Usertag
          <input
            className="auth-input"
            id="register-usertag"
            name="usertag"
            type="text"
            placeholder="@new_user"
            autoComplete="username"
          />
        </label>

        <label className="auth-label" htmlFor="username">
          Username
          <input
            className="auth-input"
            id="username"
            name="username"
            type="text"
            placeholder="Введите имя пользователя"
            autoComplete="name"
          />
        </label>

        <label className="auth-label" htmlFor="register-password">
          Пароль
          <input
            className="auth-input"
            id="register-password"
            name="password"
            type="password"
            placeholder="Введите пароль"
            autoComplete="new-password"
          />
        </label>

        <label className="auth-label" htmlFor="register-password-repeat">
          Повторите пароль
          <input
            className="auth-input"
            id="register-password-repeat"
            name="passwordRepeat"
            type="password"
            placeholder="Повторите пароль"
            autoComplete="new-password"
          />
        </label>

        <button className="auth-button" type="button">
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
