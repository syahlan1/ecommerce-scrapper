import { useState } from 'react';
import { Form, Button, Card, FloatingLabel, Alert } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

const RegisterPage = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    // Kirim request ke backend menggunakan Axios
    try {
      const response = await axios.post('http://localhost:3000/api/register', {
        email,
        username,
        password,
      });

      // Jika sukses
      setSuccess(response.data.success);
      setTimeout(() => {
        navigate('/login');
      }, 2000); // Redirect ke halaman login setelah 2 detik
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        setError(err.response.data.error); // Menampilkan error dari backend
      } else {
        setError('Something went wrong, please try again.'); // Error umum
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
          <h2 className="text-center mb-4">Register</h2>

          {error && <Alert variant="danger">{error}</Alert>}
          {success && <Alert variant="success">Registration successful! Redirecting...</Alert>}

          <Form onSubmit={handleSubmit}>
            {/* Input Email */}
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <FloatingLabel controlId="floatingEmail" label="Email address" className="mb-3">
                <Form.Control
                  type="email"
                  placeholder="name@example.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </FloatingLabel>
            </Form.Group>

            {/* Input Username */}
            <Form.Group className="mb-3" controlId="formBasicUsername">
              <FloatingLabel controlId="floatingUsername" label="Username" className="mb-3">
                <Form.Control
                  type="text"
                  placeholder="Enter username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  required
                />
              </FloatingLabel>
            </Form.Group>

            {/* Input Password */}
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
              Register
            </Button>
          </Form>
          <div className="text-center mt-3">
            Already have an account? <Link to="/login">Login</Link>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
};

export default RegisterPage;

