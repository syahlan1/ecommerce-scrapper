import { useState, useEffect } from "react";
import { Navbar, Container, Form, Button, Dropdown } from "react-bootstrap";
import { FaSearch, FaSignOutAlt } from "react-icons/fa";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

const NavbarComponent = () => {
  const [username, setUsername] = useState(null);
  const [setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState(""); // State untuk term pencarian
  const navigate = useNavigate();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get("http://localhost:3000/api/user", {
          withCredentials: true,
        });
        setUsername(response.data.username); // Menyimpan nama pengguna jika berhasil
      } catch (err) {
        console.error(err);
        setUsername(null); // Jika gagal, anggap user belum login
        setError("Not logged in");
      }
    };

    fetchUser();
  }, []);

  const handleLogout = async () => {
    try {
      await axios.post("http://localhost:3000/api/logout", {}, {
        withCredentials: true,
      });
      setUsername(null); // Hapus nama user dari state
      navigate("/login"); // Redirect ke halaman login setelah logout
    } catch (err) {
      console.error("Logout failed", err);
    }
  };

  // Fungsi untuk menangani pencarian
  const handleSearch = async (e) => {
    e.preventDefault();
    try {
      // Kirim term pencarian ke backend
      const response = await axios.post("http://localhost:3000/api/get-search", {
        product_name: searchTerm
      }, {
        withCredentials: true
      });

      // Simpan hasil pencarian ke state dan arahkan ke halaman hasil pencarian
      navigate('/search', { state: { results: response.data.search_results, term: searchTerm } });
    } catch (err) {
      console.error("Search failed", err);
    }
  };

  return (
    <div>
      <Navbar expand="lg" bg="white" data-bs-theme="light" fixed="top">
        <Container>
          <Navbar.Brand href="/">
            <img
              alt=""
              src="iconn.png"
              width="30"
              height="30"
              className="d-inline-block align-top"
            />{" "}
            Online Shop
          </Navbar.Brand>

          <Navbar.Collapse id="navbarScroll" className="justify-content-end">
            <Form className="d-flex" onSubmit={handleSearch}>
              <Form.Control
                type="search"
                placeholder="Search.."
                className="me-2"
                aria-label="Search"
                style={{ width: "600px" }}
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)} // Update term pencarian
              />
              <Button variant="outline-secondary" type="submit">
                <FaSearch />
              </Button>
            </Form>
          </Navbar.Collapse>

          <Navbar.Collapse className="justify-content-end">
            <Navbar.Text>
              {username ? (
                <Dropdown>
                  <Dropdown.Toggle variant="success" style={{ background:"none", color:"black", border:"none" }} id="dropdown-basic">
                  <img
                    src="https://static.vecteezy.com/system/resources/previews/020/911/736/original/profile-icon-user-icon-person-icon-free-png.png"
                    alt="User avatar"
                    style={{ width: "30px", height: "30px", borderRadius: "50%", marginRight: "10px" }}
                    />
                    {username}
                  </Dropdown.Toggle>

                  <Dropdown.Menu>
                    <Dropdown.Item onClick={handleLogout}>Logout <FaSignOutAlt/></Dropdown.Item>
                  </Dropdown.Menu>
                </Dropdown>
              ) : (
                <Link to="/login">
                  <Button variant="outline-success">Login</Button>
                </Link>
              )}
            </Navbar.Text>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </div>
  );
};

export default NavbarComponent;
