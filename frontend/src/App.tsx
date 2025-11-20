import { Routes, Route, BrowserRouter } from "react-router-dom"
import Login from "./pages/Login"
import Registration from "./pages/Registration"
import Chat from "./pages/Chat"
import ProtectedRoute from "./components/ProtectedRoute"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login/>}/>
        <Route path="/registration" element={<Registration/>}/>
        <Route path="/chat" element={
        <ProtectedRoute>
          <Chat/>
        </ProtectedRoute>}/>
      </Routes>
    </BrowserRouter>
  )
}

export default App
