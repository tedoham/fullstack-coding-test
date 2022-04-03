import axios from 'axios';
import React, { useEffect, useState } from 'react'

interface MenuType {
    showMenu: any
    productId: string
}

const baseURLForUpdate = "http://localhost:8000/product";

export default function Menu({showMenu, productId}: MenuType) {

    const [productsDetail, setProductsDetail] = useState<any>(null);
    
    useEffect(() => {
        axios.get(`${baseURLForUpdate}/${productId}`).then((response) => {
            setProductsDetail(response.data);
        });
    }, [])

    const updateDeliveryStatus = (statusId: string) => {
        const statusToUpdate = { delivery_status_id: statusId };
        axios.put(`${baseURLForUpdate}/${productId}`, statusToUpdate).then((response) => {
            showMenu(false);
        });
    }
    
    return (
        <div className="min-w-screen h-screen animated fadeIn faster  fixed  left-0 top-0 flex justify-center items-center inset-0 z-50 outline-none focus:outline-none bg-no-repeat bg-center bg-cover">
            <div className="absolute text-center bg-white-800 opacity-80 inset-0 z-0">
                <h1 className='text-white mt-20 text-xl'>Update Delivery Status</h1>
            </div>
            <div className="w-full  max-w-lg p-5 relative mx-auto my-auto rounded-xl shadow-lg  bg-gray-600">
                <div className="text-center p-5 flex-auto justify-center">
                    <div className="flex flex-col divide-y divide-gray-300">
                        <button 
                            onClick={() => updateDeliveryStatus("1")}
                            type="button" 
                            className={`p-4 hover:bg-gray-50 cursor-pointer ${productsDetail?.delivery_status === 'PENDING' && 'bg-blue-400 text-black'}`}>
                            PENDING
                        </button>
                        <button 
                            onClick={() => updateDeliveryStatus("2")}
                            type="button" 
                            className={`p-4 hover:bg-gray-50 cursor-pointer ${productsDetail?.delivery_status === 'ORDERED' && 'bg-blue-400 text-black'}`}>
                            ORDERED
                        </button>
                        <button 
                            onClick={() => updateDeliveryStatus("3")}
                            type="button" 
                            className={`p-4 hover:bg-gray-50 cursor-pointer ${productsDetail?.delivery_status === 'SHIPPED' && 'bg-blue-400 text-black'}`}>
                            SHIPPED
                        </button>
                        <button 
                            onClick={() => updateDeliveryStatus("4")}
                            type="button" 
                            className={`p-4 hover:bg-gray-50 cursor-pointer ${productsDetail?.delivery_status === 'CANCELLED' && 'bg-blue-400 text-black'}`}>
                            CANCELLED
                        </button>
                    </div>
                </div>
                <div className="p-3  mt-2 text-center space-x-4 md:block">
                    <button 
                        onClick={() => showMenu(false)}
                        className="mb-2 md:mb-0 bg-white px-5 py-2 text-sm shadow-sm font-medium tracking-wider border text-gray-600 rounded-full hover:shadow-lg hover:bg-gray-100">
                        Cancel
                    </button>
                </div>
            </div>
        </div>
    )
}
