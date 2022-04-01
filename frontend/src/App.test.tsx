import axios from "axios";
import { render, screen } from "@testing-library/react"
import App from './App';

// jest.mock('axios');

const fakeProducts = [
    {
        "product_id": "2",
        "product_type_name": "Product-type-2",
        "customer_name": "Cutomer-2",
        "delivery_status": "ORDERED",
        "delivery_address": "address-two",
        "estimated_delivery_date": "2022-03-31T05:43:53.774977Z"
    }, 
    {
        "product_id": "3",
        "product_type_name": "Product-type-3",
        "customer_name": "Cutomer-3",
        "delivery_status": "SHIPPED",
        "delivery_address": "address-three",
        "estimated_delivery_date": "2022-03-31T05:43:53.774977Z"
    }
];

describe("Render table page", () => {
    test('renders table', async () => {
        jest.spyOn(axios, 'get').mockResolvedValue({ data: fakeProducts });
        render(<App />);
    }); 
})