import { useState, useEffect } from 'react';
import { Row, Col, Button, Card } from 'react-bootstrap';
import { FaArrowLeft, FaEye, FaHeart } from "react-icons/fa";
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';

const DetailPage = () => {
  const navigate = useNavigate();
  const { id } = useParams(); // Ambil product ID dari URL
  const [product, setProduct] = useState(null); // State untuk menyimpan detail produk
  const [error, setError] = useState(null); // State untuk menangani error

  // Fungsi untuk kembali ke halaman sebelumnya
  const handleBack = () => {
    navigate(-1);
  };

  // useEffect untuk mengambil detail produk dari backend
  useEffect(() => {
    const fetchProductDetails = async () => {
      try {
        const response = await axios.get(`http://localhost:3000/api/product/${id}`, {
          withCredentials: true, // Kirim cookie untuk autentikasi
        });
        setProduct(response.data.product); // Simpan produk ke state
      } catch (err) {
        console.error(err)
        setError("Product not found or you are not logged in");
      }
    };

    fetchProductDetails();
  }, [id]); // Jalankan ulang saat product ID berubah

  if (error) {
    return <div className="container mt-5 pt-5"><h3>{error}</h3></div>;
  }

  if (!product) {
    return <div className="container mt-5 pt-5"><h3>Loading...</h3></div>;
  }

  return (
    <div className="container mt-5 pt-5">
      <button onClick={handleBack} style={{ background: "none", color: "black", border: "none" }}>
        <FaArrowLeft className="mb-5" style={{ fontSize: "20px" }} />
      </button>
      <Row>
        {/* Bagian Gambar Produk */}
        <Col md={6}>
          <Card style={{ border: "none" }}>
            <Card.Img style={{ width: "450px", height: "450px" }} src={product.image} alt={product.name} />
          </Card>
        </Col>

        {/* Bagian Detail Produk */}
        <Col md={6}>
          <h2>{product.name}</h2>
          <div className="d-flex">
            <p className="text-muted me-5">
              <FaEye /> {product.view} dilihat
            </p>
            <p className="text-muted">
              <FaHeart /> {product.like} Disukai
            </p>
          </div>
          <p className="text-muted mt-2">
            {product.description}
          </p>
          <h3 className="text-success">Rp. {product.price.toLocaleString('id-ID')}</h3>

          {/* Tombol Beli */}
          <div className="d-flex mt-4">
            <Button variant="success" className="me-2">Beli Sekarang</Button>
            <Button variant="outline-secondary">Tambahkan ke Keranjang</Button>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default DetailPage;
