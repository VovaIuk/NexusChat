import Chat from './pages/Chat'
import Login from './pages/Login'
import Auth from './pages/Auth'
import ProtectedRoute from './components/routing/ProtectedRoute'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css'
import { UserProvider } from './contexts/UserContext';

function App() {
  return (
    <BrowserRouter>
      <UserProvider>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/auth" element={<Auth />} />
          <Route
            path="/"
            element={
              <ProtectedRoute>
                <Chat />
              </ProtectedRoute>
            }
          />
        </Routes>
      </UserProvider>
    </BrowserRouter>
  );
}

export default App
