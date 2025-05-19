# LinkBush – Serverless User Page Hosting with AWS

LinkBush is a serverless application that allows users to create and manage personalized link pages, similar to Linktree. Users can create pages with links, profile information, and a profile picture. The backend is powered by AWS Lambda and DynamoDB, and the frontend is hosted as a static site.

---

## 🚀 Features

- Create a user page with profile information and links
- View a user's page by ID
- Delete a user page
- Upload and serve profile pictures

---

## 🛠️ Technologies Used

### 🧩 AWS Services

| Service       | Purpose                                                  |
|---------------|----------------------------------------------------------|
| **Lambda**    | Hosts the Go backend using Function URLs                 |
| **API Gateway (implicit via Lambda URL)** | Accepts public HTTPS requests             |
| **DynamoDB**  | Stores user page data (user ID, links, profile info)     |
| **S3**        | Stores and serves user profile pictures (public access)  |

### 📁 Frontend

- HTML + JavaScript (hosted as a static site)
- Uses Fetch API to call Lambda Function URLs
- Hosted publicly (in my case S3)

---

## 📦 Endpoints

### Create Page
`POST /create_page`

- Accepts JSON data with user info and links
- Stores data in DynamoDB

### Get Page
`GET /get_page/:user_id`

- Fetches and returns user page content based on user ID

### Delete Page
`POST /delete_page/:user_id`

- Deletes user page by ID
- Requires body with necessary confirmation data

---

## 🖼 Profile Pictures

- Profile images are uploaded manually to an S3 bucket (publicly accessible)
- Image URL is referenced in the user’s page data

---

