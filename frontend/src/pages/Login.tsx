import "./Auth.css"
import {loginUser} from "../api/authApi"
import {useState} from "react";
import { useNavigate } from "react-router-dom";

function Login() {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [remember, setRemember] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try{
      const response = await loginUser({});
      if(remember){
        localStorage.setItem("token", "test_token");
      } else{
        sessionStorage.setItem("token", "test_token");
      }
      navigate("/chat");
    } catch(e){
      console.error(e);
    } finally{
      setLoading(false);
    }
  }

  return (
    <div className="wrapper-center">
      <main className="card">
        <div className="brand">
          <div className="logo">LF</div>
          <div>
            <h1 className="form-title">Вход в приложение</h1>
            <p className="lead">Введите учётные данные для доступа к чату</p>
          </div>
        </div>

        <form onSubmit={handleSubmit}>
          <div className="field">
            <label htmlFor="username">Тег пользователя</label>
            <input
              id="username"
              type="text"
              placeholder="Введите тег пользователя"
              required
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>

          <div className="field">
            <label htmlFor="password">Пароль</label>
            <input
              id="password"
              type="password"
              placeholder="Введите пароль"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="action">
            <label className="remember">
              <input 
                type="checkbox"
                checked={remember}
                onChange={(e)=>setRemember(e.target.checked)}
              />
              Запомнить
            </label>
            <button className="btn" type="submit" disabled={loading}>
              {loading ? "Вход...": "Войти"}
            </button>
          </div>
          <p className="helper">
            Нет аккаунта? <a href="./registration">Зарегестрироваться</a>
          </p>
        </form>
      </main>
    </div>
  )
}

export default Login
