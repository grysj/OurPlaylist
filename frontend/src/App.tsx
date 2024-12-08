
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './Routes/Home';
import User from './Routes/User';
import Profile from './Routes/Profile';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="user" element={<User />} />
        <Route path="profile/:id" element={<Profile />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
