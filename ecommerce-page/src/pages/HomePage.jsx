import { useState, useEffect } from 'react';
import { Carousel, Card } from 'react-bootstrap';
import axios from 'axios';
import { Link } from 'react-router-dom'; // Import Link dari react-router-dom

function HomePage() {
  const [products, setProducts] = useState([]);
  const [recommendedProducts, setRecommendedProducts] = useState([]);
  const [lastViewedProducts, setLastViewedProducts] = useState([]); // State untuk produk terakhir dilihat
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Cek apakah user sudah login
        const userResponse = await axios.get('http://localhost:3000/api/user', { withCredentials: true });
        if (userResponse.data) {
          setIsLoggedIn(true);

          // Jika user sudah login, ambil rekomendasi produk
          const recommendResponse = await axios.get('http://localhost:3000/api/products/recommend', { withCredentials: true });
          setRecommendedProducts(recommendResponse.data.recommended_products);

          // Ambil produk terakhir dilihat
          const lastViewedResponse = await axios.get('http://localhost:3000/api/products/last-viewed', { withCredentials: true });
          setLastViewedProducts(lastViewedResponse.data.last_viewed_products);
        }
      } catch (error) {
        console.error(error);
        setIsLoggedIn(false); // Jika user tidak login
      }

      // Ambil semua produk, baik user login maupun tidak
      const allProductsResponse = await axios.get('http://localhost:3000/api/get-all-product');
      setProducts(allProductsResponse.data);
    };

    fetchData();
  }, []);

  const renderProducts = (productList) => {
    return productList.map((product) => (
      <Card key={product.id} style={{ width: '14rem' }} className='mb-4 me-3'>
        <Link to={`/product/${product.id}`}>
          <Card.Img variant="top" style={{ width: '13.9rem', height: '14rem' }} src={product.image} alt={product.name} />
        </Link>
        <Card.Body>
          <Card.Title style={{ fontSize: "16px" }}>{product.name}</Card.Title>
          <Card.Text style={{ fontWeight: "bold", color: "#20bf6b" }}>
            Rp. {product.price.toLocaleString()}
          </Card.Text>
          <Card.Text style={{ fontSize: "14px", color: "#a5b1c2" }}>{product.like} Suka</Card.Text>
        </Card.Body>
      </Card>
    ));
  };

  return (
    <div className='mt-5'>
      {/* Carousel */}
      <Carousel>
        <Carousel.Item>
          <img className="d-block w-100" src="../../img/slide/slide1.png" alt="First slide" />
        </Carousel.Item>
        <Carousel.Item>
          <img className="d-block w-100" src="../../img/slide/slide2.png" alt="Second slide" />
        </Carousel.Item>
        <Carousel.Item>
          <img className="d-block w-100" src="../../img/slide/slide3.png" alt="Third slide" />
        </Carousel.Item>
      </Carousel>

      {/* Produk Terakhir Dilihat */}
      {isLoggedIn && lastViewedProducts.length > 0 && (
        <div className='mt-5 px-5'>
          <h3>Lanjut Cek Produkmu</h3>
          <div className='mt-4 px-3 d-flex flex-wrap justify-content-start'>
            {renderProducts(lastViewedProducts)}
          </div>
        </div>
      )}

      {/* Rekomendasi Produk */}
      {isLoggedIn && recommendedProducts.length > 0 && (
        <div className='mt-5 px-5'>
          <h3>Rekomendasi Produk</h3>
          <div className='mt-4 px-3 d-flex flex-wrap justify-content-start'>
            {renderProducts(recommendedProducts)}
          </div>
        </div>
      )}

      {/* Semua Produk */}
      <div className='mt-5 px-5'>
        <h3>Semua Produk</h3>
        <div className='mt-4 px-3 d-flex flex-wrap justify-content-start'>
          {renderProducts(products)}
        </div>
      </div>
    </div>
  );
}

export default HomePage;
