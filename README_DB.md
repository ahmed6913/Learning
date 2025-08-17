# NoSQL & SQL db 
---

# 📘 NoSQL Cheatsheet

A complete and concise reference guide to NoSQL databases, covering all core concepts, types, use-cases, and examples — perfect for developers and learners.

---

## 📖 What is NoSQL?

**NoSQL = Not Only SQL**  
NoSQL databases store data in a non-relational, schema-less format. They are designed for:

- 🔁 Flexibility
- ⚡ Speed
- 📈 Scalability (horizontal)
- 🌍 Distributed environments
- 📦 Handling large volumes of unstructured or semi-structured data

---

## 🧩 NoSQL Database Types

| Type        | Description                         | Common Examples        | Use Cases                       |
|-------------|-------------------------------------|------------------------|---------------------------------|
| 📄 Document  | JSON-like documents (BSON)          | MongoDB, CouchDB       | User profiles, CMS, catalogs    |
| 🔑 Key-Value | Key-value pairs                     | Redis, DynamoDB        | Caching, sessions, counters     |
| 📊 Column    | Columnar store                      | Cassandra, HBase       | Time-series, analytics, logs    |
| 🕸️ Graph     | Nodes + relationships (edges)       | Neo4j, ArangoDB        | Social networks, recommendations|

---

## ⚙️ Core Features

- ❌ No rigid schema
- 🔁 Easy replication
- 🧱 High availability
- 📊 Designed for big data
- 🚀 Fast read/write speeds
- 🔄 Auto-sharding (in most engines)

---

## 📦 Use Case Matrix

| Requirement                  | Recommended NoSQL Type   |
|-----------------------------|--------------------------|
| Fast lookups & caching      | Key-Value (Redis)        |
| Flexible JSON storage       | Document (MongoDB)       |
| Complex relationships       | Graph (Neo4j)            |
| Real-time analytics         | Column (Cassandra)       |
| Massive write throughput    | Cassandra, DynamoDB      |
| Schema-less data ingestion  | Document or Key-Value    |

---

- use commands or GUI for NoSQL db 

## 🧪 MongoDB Commands (Document DB)

```bash
# Connect to MongoDB
mongo

# Show databases
show dbs

# Switch/create database
use myDatabase

# Insert a document
db.users.insertOne({ name: "Alice", age: 25 })

# Find documents
db.users.find({ age: { $gt: 20 } })

# Update document
db.users.updateOne({ name: "Alice" }, { $set: { age: 26 } })

# Delete document
db.users.deleteOne({ name: "Alice" })

SELECT * FROM user_emails;      -- Query the view
DROP VIEW user_emails;          -- Remove the view


```

# 🧠 MySQL SQL Cheatsheet

A complete reference guide for learning and using **SQL with MySQL DBMS** — covering syntax, operations, data types, joins, constraints, indexing, and more. Ideal for students, developers, and interview prep.


## 📘 What is SQL?

**SQL (Structured Query Language)** is used to manage data in relational databases.  
MySQL is one of the most popular open-source RDBMS engines that uses SQL for data storage, retrieval, and administration.

---

## 🗂️ Database Basics  
_Commands to manage your database._

```sql
CREATE DATABASE mydb;          -- Create a new database
USE mydb;                      -- Select a database to use
SHOW DATABASES;                -- List all databases
DROP DATABASE mydb;            -- Delete a database

```
## 📁 Table Basics
Commands to create and manage tables.
```sql

CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE,
  age INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);                             -- Create a users table

DESCRIBE users;                -- Show table structure
DROP TABLE users;              -- Delete the table

```
## 🔧 Common Data Types
Key MySQL data types and their usage.


```sql
Type	Example	Description
INT	1, 100	Whole numbers
DECIMAL	10.99	Fixed decimal point numbers
FLOAT	99.99	Approx. floating point numbers
VARCHAR(n)	'John'	Variable-length text (up to n chars)
TEXT	Long article	Long text content
DATE	'2025-07-24'	Stores date
TIMESTAMP	Auto time	Date & time of record creation

```
## 🧪 CRUD Operations
Create, Read, Update, Delete rows in a table.

```sql
INSERT INTO users (name, email, age)
VALUES ('Alice', 'alice@mail.com', 25);       -- Add a row

SELECT * FROM users;                          -- Read all data
SELECT name, age FROM users WHERE age > 18;   -- Read filtered data

UPDATE users SET age = 30 WHERE name = 'Alice';  -- Update data

DELETE FROM users WHERE id = 1;               -- Delete data

```
## 📌 Constraints
Add rules to maintain data integrity.

```sql
Constraint	Description
PRIMARY KEY	Uniquely identifies each row
UNIQUE	All values must be different
NOT NULL	Field must have a value
DEFAULT	Sets a default value if none is given
FOREIGN KEY	Links rows to another table's primary key
CHECK	Enforces a condition on data input

ALTER TABLE orders
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id);   -- Add a foreign key


```
## 🤝 Joins & union
Join multiple tables based on relationships. & union

```sql
-- INNER JOIN: Matches rows in both tables
SELECT orders.id, users.name
FROM orders
INNER JOIN users ON orders.user_id = users.id;

-- LEFT JOIN: All users, even if no orders
SELECT users.name, orders.id
FROM users
LEFT JOIN orders ON users.id = orders.user_id;

-- RIGHT JOIN: All orders, even if no user
SELECT users.name, orders.id
FROM users
RIGHT JOIN orders ON users.id = orders.user_id;

-- union
Query for union
SELECT Name FROM Employees_US
UNION
SELECT Name FROM Employees_UK;

Query for union all 
SELECT Name FROM Employees_US
UNION ALL 
SELECT Name FROM Employees_UK;



```
## 🧠 Subqueries
Nested queries for more complex filtering.


```sql
-- Subquery with IN
SELECT name FROM users
WHERE id IN (SELECT user_id FROM orders WHERE amount > 100);

-- Correlated subquery with EXISTS
SELECT name FROM users u
WHERE EXISTS (
  SELECT 1 FROM orders o WHERE o.user_id = u.id AND o.amount > 100
);


```
## 📊 Aggregate Functions
Summary operations over rows.

```sql
SELECT COUNT(*) FROM users;      -- Total rows
SELECT AVG(age) FROM users;      -- Average age
SELECT MAX(age) FROM users;      -- Maximum age
SELECT MIN(age) FROM users;      -- Minimum age
SELECT SUM(age) FROM users;      -- Total sum of ages

```

## 📑 GROUP BY & HAVING
Group rows and filter grouped results.

```sql
SELECT age, COUNT(*) FROM users
GROUP BY age;                    -- Group users by age

SELECT age, COUNT(*) FROM users
GROUP BY age
HAVING COUNT(*) > 2;             -- Filter grouped results

```
## 🔍 ORDER BY & LIMIT
Sort and restrict result rows.

```sql

SELECT * FROM users
ORDER BY age DESC;               -- Sort by age descending

SELECT * FROM users
LIMIT 5;                         -- Show only 5 results

```
## 🧱 Indexes
Improve search performance on large tables.

```sql
CREATE INDEX idx_name ON users(name);   -- Add index on name
DROP INDEX idx_name ON users;           -- Remove index


```

## 📋 Views
Save complex queries as virtual tables.
```

CREATE VIEW user_emails AS
SELECT name, email FROM users;


```







