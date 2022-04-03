import axios from 'axios';
import { useEffect, useState } from 'react'
import './App.css'
import Menu from './components/Menu';
import ProductDetail from './components/ProductDetail';

const baseURL = "http://localhost:8000/products";

function App() {
  const [menuStatus, setMenuStatus] = useState(false);
  const [detailStatus, setDetailStatus] = useState(false);
  const [products, setProducts] = useState<any>(null);
  const [productId, setProductId] = useState<string>("");

  const showMenu = (productIdToShow: string) => {
    setProductId(productIdToShow)
    setMenuStatus(!menuStatus)
  }

  const showDetail = (productIdToShow: string) => {
    setProductId(productIdToShow)
    setDetailStatus(!detailStatus)
  }

  useEffect(() => {
    axios.get(baseURL).then((response) => {
      setProducts(response.data);
    });
  })
  

  if (!products) return (
    <div className="flex flex-col justify-center items-center min-h-screen bg-gray-800">
      <div className="col-span-12 text-white">
        No results found. Please check if your API server is running.
      </div>
    </div>
  )

  return (
    <div className="flex flex-col justify-center items-center min-h-screen bg-gray-800">
      <div className="col-span-12">
        <div className="overflow-auto lg:overflow-visible ">
          <table className="table text-gray-400 border-separate space-y-6 text-sm">
            <thead className="bg-gray-800 text-gray-500">
              <tr>
                <th className="p-3">Product</th>
                <th className="p-3 text-left">Customer</th>
                <th className="p-3 text-left">Address</th>
                <th className="p-3 text-left">Delivery Date</th>
                <th className="p-3 text-left">Status</th>
                <th className="p-3 text-left">Action</th>
              </tr>
            </thead>
            <tbody>
              {products.map((pro: any) => (
                <tr className="bg-gray-700" key={pro.product_type_name}>
                  <td className="p-3">
                    <div className="flex items-center">
                        <div className="">{pro.product_type_name}</div>
                    </div>
                  </td>
                  <td className="p-3">
                    {pro.customer_name}
                  </td>
                  <td className="p-3">
                    {pro.delivery_address}
                  </td>
                  <td className="p-3">
                    {new Date(pro.estimated_delivery_date).toDateString()}
                  </td>
                  <td className="p-3 flex justify-center">
                    <button 
                      onClick={() => showMenu(pro.product_id)}
                      type="button" 
                      className="bg-blue-400 hover:bg-blue-600 text-gray-50 rounded-md px-2 py-1">
                      {pro.delivery_status}
                    </button>
                    {menuStatus && <Menu showMenu={showMenu} productId={productId} />}
                  </td>
                  <td className="p-3">
                    <button 
                      onClick={() => showDetail(pro.product_id)}
                      type='button' 
                      className="text-gray-400 hover:text-gray-100 mr-2">
                      <i className="material-icons-outlined text-base">visibility</i>
                    </button>
                    {detailStatus && <ProductDetail showDetail={showDetail} productId={productId} />}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
	    </div>
    </div>
  )
}

export default App
