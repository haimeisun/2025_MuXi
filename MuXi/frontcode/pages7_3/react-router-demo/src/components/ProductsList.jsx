// src/components/ProductsList.jsx
const ProductsList = () => {
  return (
    <div>
      <h1>商品列表</h1>
      <ul>
        <li>
          <Link to="/products/1">商品一</Link>
        </li>
        <li>
          <Link to="/products/2">商品二</Link>
        </li>
        <li>
          <Link to="/products/3">商品三</Link>
        </li>
      </ul>
    </div>
  );
};

export default ProductsList;
