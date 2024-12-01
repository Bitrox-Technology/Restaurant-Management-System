# Restaurant Management System

A comprehensive RESTful API built with Go for managing restaurant operations. This system includes functionalities for managing food items, menus, invoices, tables, users, orders, and order items.

---

## 🚀 Features

- **Food Management**: Add, update, and remove food items.
- **Menu Management**: Manage menus linked to specific food items.
- **Table Management**: Manage table availability and reservations.
- **Order Management**: Handle customer orders and associated items.
- **Invoice Generation**: Generate and manage invoices for orders.
- **User Management**: Add and authenticate restaurant users.

---

## 📋 Prerequisites

1. **Go**: Version 1.19 or later.
2. **Database**: MongoDB (or your preferred database).
3. **Tools**: Postman (for API testing).

---

## 🛠 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Bitrox-Technology/Restaurant-Management-System.git
cd Restaurant-Management-System

.env
DB_URI=mongodb://localhost:27017
DB_NAME=restaurant
PORT=8000
JWT_SECRET=your_jwt_secret

run server
go run main.go

