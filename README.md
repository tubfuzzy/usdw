# USDW - Xero Bank Feeds API Integration

This project is a backend service built using **GoFiber** and integrates with **Xero Bank Feeds API**. It follows **Clean Architecture** and is deployed on **Azure App Service** using **Docker** and **GitHub Actions**.

---

## 🚀 Features

- **OAuth2.0 Authentication** with Xero
- **Bank Feeds API Integration**
- **CRUD Operations on Feed Connections**
- **Transaction Statements Handling**
- **Logging with `x-request-id` for Request Tracking**
- **Dockerized Deployment on Azure**

---

## ⚡ API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST**   | `/api/v1/feed-connections` | Create a new feed connection |
| **GET**    | `/api/v1/feed-connections` | Retrieve all feed connections |
| **GET**    | `/api/v1/feed-connections/:id` | Retrieve a feed connection by ID |
| **DELETE** | `/api/v1/feed-connections/:id` | Delete a feed connection |

### **Statements**
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST**   | `/api/v1/statements` | Post a new bank statement |
| **GET**    | `/api/v1/statements` | Retrieve all statements |
| **GET**    | `/api/v1/statements/:id` | Retrieve a statement by ID |

---

## 📝 API Request Examples (cURL)
```sh
curl -X POST "http://your-app-url/api/v1/feed-connections" \
     -H "Content-Type: application/json" \
     -d '{
           "accountToken": "your-account-token",
           "accountName": "My Bank Account"
         }'