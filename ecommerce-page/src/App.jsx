import NavbarComponent from "./components/NavbarComponent";
import FooterComponent from "./components/FooterComponent";


import { Routes, Route, useLocation } from 'react-router-dom';
import HomePage from "./pages/HomePage";
import DetailPage from "./pages/DetailPage";
import SearchShowPage from "./pages/SearchShowPage";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";

function App() {
  const location = useLocation();

  // Periksa apakah path saat ini adalah "/login"
  const isLoginPage = location.pathname === '/login' || location.pathname === '/register' ;
 return <div>
  {!isLoginPage && <NavbarComponent />}
    <Routes>
      <Route path="/" Component={HomePage}/>
      <Route path="/product/:id" Component={DetailPage}/>
      <Route path="/search" Component={SearchShowPage}/>
      <Route path="/login" Component={LoginPage}/>
      <Route path="/register" Component={RegisterPage}/>
    </Routes>
  {!isLoginPage && <FooterComponent />}
 </div>;
}

export default App
