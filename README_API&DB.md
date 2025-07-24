# SQL & NoSQL 

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

