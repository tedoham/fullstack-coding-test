# check this video for demo https://www.loom.com/share/cfbdc18c78c3429ba8f9e99831f32449

## Tasks

Below will be your set of tasks to accomplish. This assignment is to build a dashboard for admins to manage their customer's product shipment status. You do not have to complete the tasks in order, but please state clearly in your PR what is solved. You are encouraged to use Golang + React for this assignment.

Success criteria will be defined clearly for each task

1. [Database](#database)
1. [Backend](#backend)
1. [Frontend](#frontend)
1. [Deployment](#deployment)

Out of scope: you don't need to handle authentication


### Database
Please implement the following
- A database that can store the information:
    - product type (`name`)
    - customer (`name`, `contact_number`)
    - product (`product_type_id`, `customer_id`, `delivery_status`, `delivery address`, `estimated_delivery_date`)
- The product `delivery_status` values are from a fixed range: `PENDING`, `ORDERED`, `SHIPPED` and `CANCELLED`
- You may use either a relational or non-relational database, that is up to your preference.

Success Criteria
- Create a pull request against `master` of your fork with a clear description of the change and purpose and merge it
- Generate the table script

### Backend
Please implement the following
- Documentation on how to start the backend application
- A REST API to retrieve a list of `product` with its fields, plus extra fields
    - customer name
    - product type name
- A REST API to update the product's `delivery_status`
- Automated tests
- BONUS: pagination
- BONUS: prevents SQL injection

Success Criteria
- Create a pull request against `master` of your fork with a clear description of the change and purpose and merge it
- The backend should be connected to the database
- Screenshot of test summary should be attached to the PR

### Frontend
Please implement the following
- A page to display a list of products and the product status, the customer name, the product type name, and the estimated delivery date
- A button to manually update the `delivery_status` of the product
- Documentation on how to start the frontend application
- Automated tests
- BONUS: A page to display the product by ID

Success Criteria
- Create a pull request against `master` of your fork with a clear description of the change and purpose and merge it
- Screenshots or videos of the web application should be attached to the PR
- Screenshot of test summary should be attached to the PR

### Deployment
Please implement the following
- A deployment pipeline that builds and deploys the web application

Success Criteria
- Create a pull request against `master` of your fork with a clear description of the change and purpose and merge it
- The CI/CD file should be in the pull request
- BONUS: a link such that makes the web application publicly accessibly
