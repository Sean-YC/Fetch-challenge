# **Receipt Processor API 🥾🚀**
A simple RESTful API to process retail receipts and calculate reward points based on predefined rules.

## **📌 Features**
- Process receipts and assign unique receipt IDs.
- Calculate reward points based on specific rules.
- Retrieve points using receipt ID.
- Lightweight in-memory storage.

---

## **📚 API Endpoints**
### **1⃣ Process a Receipt**
#### **Request**
- **Method:** `POST`
- **URL:** `/receipts/process`
- **Headers:** `Content-Type: application/json`
- **Body Example:**
  ```json
  {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "total": "35.35",
    "items": [
      { "shortDescription": "Item A", "price": "10.00" },
      { "shortDescription": "Item B", "price": "25.35" }
    ]
  }
  ```
#### **Response**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```

---

### **2⃣ Get Points for a Receipt**
#### **Request**
- **Method:** `GET`
- **URL:** `/receipts/{id}/points`
- **Example:** `/receipts/123e4567-e89b-12d3-a456-426614174000/points`

#### **Response**
```json
{
  "points": 45
}
```

---

## **🚀 Getting Started**
### **🔹 Prerequisites**
- **Go** (>= 1.18)
- **cURL** (for testing API)

### **🔹 Installation**
1. Clone the repository:
   ```sh
   git clone https://github.com/YOUR_GITHUB/Fetch-challenge.git
   cd Fetch-challenge
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. The API runs on **`http://localhost:8080`** 🎉

---

## **🧪 Testing the API**
### **1⃣ Process a Receipt**
```sh
curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d '{
           "retailer": "Target",
           "purchaseDate": "2022-01-01",
           "purchaseTime": "13:01",
           "total": "35.35",
           "items": [
               { "shortDescription": "Item A", "price": "10.00" },
               { "shortDescription": "Item B", "price": "25.35" }
           ]
         }'
```
### **2⃣ Get Points for a Receipt**
```sh
curl -X GET "http://localhost:8080/receipts/123e4567-e89b-12d3-a456-426614174000/points"
```