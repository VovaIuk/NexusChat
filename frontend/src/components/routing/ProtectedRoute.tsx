import { Navigate } from 'react-router-dom';
import { useUser } from "../../contexts/UserContext"

const TOKEN_KEY = 'token';

export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const token = localStorage.getItem(TOKEN_KEY);
  const {user, isAuthReady} = useUser();

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  if (!isAuthReady){
    return <div>Загрузка…</div>;
  }

  if (!user){
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
}