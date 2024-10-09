import { useLocation } from "react-router-dom";
import Card from 'react-bootstrap/Card';
import {Link} from 'react-router-dom';

const SearchShowPage = () => {
  const location = useLocation();
  const { results, term } = location.state || { results: [], term: "" };

  return (
    <div className='pt-5'>
      <div className='mt-5 px-5'>
        <h3>Hasil dari: {term}</h3>
        <div className='mt-4 px-3 d-flex flex-wrap justify-content-start'>
          {results.length > 0 ? (
            results.map((product) => (
              <Card key={product.id} style={{ width: '14rem' }} className='mb-4 me-3'>
                <Link to={`/product/${product.id}`}>
                <Card.Img variant="top" style={{ width: '13.9rem', height:'14rem' }} src={product.image || "https://via.placeholder.com/150"} />
                </Link>
                <Card.Body>
                  <Card.Title style={{ fontSize:"16px"}}>{product.name}</Card.Title>
                  <Card.Text style={{ fontWeight:"bold", color:"#20bf6b" }}>
                    Rp. {product.price.toLocaleString('id-ID')}
                  </Card.Text>
                  <Card.Text style={{ fontSize:"14px", color:"#a5b1c2" }}>{product.like} Suka</Card.Text>
                </Card.Body>
              </Card>
            ))
          ) : (
            <p>Tidak ada hasil yang ditemukan.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default SearchShowPage;
