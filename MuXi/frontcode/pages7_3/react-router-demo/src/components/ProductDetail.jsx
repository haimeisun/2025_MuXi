// src/components/ProductDetail.jsx
import { useParams } from "react-router-dom";

const ProductDetail = () => {
  const { id } = useParams();
  return (
    <div>
      <h1>商品详情 #{id}</h1>
      <p>这是商品 {id} 的详细信息</p>
    </div>
  );
};

export default ProductDetail;
