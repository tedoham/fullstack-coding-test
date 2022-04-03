import axios from 'axios';
import React, { useEffect, useState } from 'react'

const baseURL = "http://localhost:8000/product";

interface DetailType {
    productId: string
    showDetail: any
}

export default function ProductDetail({showDetail, productId}: DetailType) {
    const [productsDetail, setProductsDetail] = useState<any>(null);
    useEffect(() => {
        axios.get(`${baseURL}/${productId}`).then((response) => {
            setProductsDetail(response.data);
        });
    }, [])
    
    return (
        <div className="min-w-screen h-screen animated fadeIn faster  fixed  left-0 top-0 flex justify-center items-center inset-0 z-50 outline-none focus:outline-none bg-no-repeat bg-center bg-cover">
            <div className="absolute text-center bg-gray-800 opacity-80 inset-0 z-0">
                <h1 className='text-white mt-20 text-xl'>Product Detail</h1>
            </div>
            <div className="w-full  max-w-lg p-5 relative mx-auto my-auto rounded-xl shadow-lg  bg-gray-600">
                <div className="text-center p-5 flex flex-col justify-center">
                    <div className="flex flex-row justify-center hover:bg-gray-500">
                        <span className="p-4 w-full text-right  cursor-pointer">Product Name</span>
                        <span className="p-4 w-full text-left  cursor-pointer">
                            {productsDetail?.product_type_name ?? ""}
                        </span>
                    </div>
                    <div className="flex flex-row justify-center hover:bg-gray-500">
                        <span className="p-4 w-full text-right  cursor-pointer">Customer Name</span>
                        <span className="p-4 w-full text-left  cursor-pointer">
                            {productsDetail?.customer_name ?? ""}
                        </span>
                    </div>
                    <div className="flex flex-row justify-center hover:bg-gray-500">
                        <span className="p-4 w-full text-right  cursor-pointer">Address</span>
                        <span className="p-4 w-full text-left  cursor-pointer">
                            {productsDetail?.delivery_address ?? ""}
                        </span>
                    </div>
                    <div className="flex flex-row justify-center hover:bg-gray-500">
                        <span className="p-4 w-full text-right  cursor-pointer">Delivery Date</span>
                        <span className="p-4 w-full text-left  cursor-pointer">
                            {new Date(productsDetail?.estimated_delivery_date).toDateString() ?? ""}
                        </span>
                    </div>
                    <div className="flex flex-row justify-center hover:bg-gray-500">
                        <span className="p-4 w-full text-right  cursor-pointer">Status</span>
                        <span className="p-4 w-full text-left  cursor-pointer">
                            {productsDetail?.delivery_status ?? ""}
                        </span>
                    </div>  
                </div>
                <div className="p-3  mt-2 text-center md:block">
                    <button 
                        onClick={() => showDetail(false)}
                        className="mb-2 md:mb-0 bg-white px-5 py-2 text-sm shadow-sm font-medium tracking-wider border text-gray-600 rounded-full hover:shadow-lg hover:bg-gray-100">
                        Close
                    </button>
                </div>
            </div>
        </div>
    )
}
