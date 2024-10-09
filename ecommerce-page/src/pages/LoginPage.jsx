import { useState } from 'react';
import { Form, Button, Card, FloatingLabel, Alert } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    try {
      // Kirim request login ke backend
      const response = await axios.post('http://localhost:3000/api/login', {
        username,
        password,
      },{
        withCredentials: true,
      });
      
      
      if (response.data.success) {
        navigate('/');
      } else {
        setError(response.data.error);
      }
      // Jika login berhasil, redirect ke halaman lain, misalnya ke dashboard
      navigate('/'); // Anda bisa ganti '/dashboard' dengan route lain yang Anda inginkan
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        // Tampilkan pesan error dari backend
        setError(err.response.data.error);
      } else {
        // Tampilkan pesan error umum
        setError('Login failed. Please try again.');
      }
    }
  };

  return (
    <div className="d-flex justify-content-center align-items-center vh-100">
      <Card style={{ width: '24rem' }}>
        <Card.Body>
          <Link to="/">
            <Card.Img
              className="d-block mx-auto mb-3"
              src="iconn.png"
              style={{ width: '50px', height: '50px' }}
            ></Card.Img>
          </Link>
          <h2 className="text-center mb-4">Login</h2>

          {error && <Alert variant="danger">{error}</Alert>}

          <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3" controlId="formBasicUsername">
              <FloatingLabel
                controlId="floatingInput"
                label="Username"
                className="mb-3"
              >
                <Form.Control
                  type="text"
                  placeholder="Enter username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  required
                />
              </FloatingLabel>
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <FloatingLabel controlId="floatingPassword" label="Password">
                <Form.Control
                  type="password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
              </FloatingLabel>
            </Form.Group>

            <Button variant="success" type="submit" className="w-100">
              Login
            </Button>
          </Form>

          <p className="text-center mt-3">
            Dont have an account? <Link to="/register">Register</Link>
          </p>
        </Card.Body>
      </Card>
    </div>
  );
};

export default LoginPage;
