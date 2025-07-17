# 🔥 Firebase Developer Cheat Sheet (Basic to Advanced)

## 📦 Getting Started

- Go to: https://console.firebase.google.com
- Create a Firebase Project
- Add Web/Android/iOS App to Firebase
- Install Firebase CLI:

```bash
npm install -g firebase-tools

## 🚀 Firebase CLI Commands

firebase login                     # Log in to Firebase
firebase init                     # Initialize project
firebase deploy                   # Deploy to Firebase
firebase emulators:start          # Start local emulator
firebase use --add                # Add project alias
firebase logout                   # Log out

## ⚙️ SDK Setup (Web)


npm install firebase

// firebase.js
import { initializeApp } from 'firebase/app';
import { getFirestore } from 'firebase/firestore';
import { getAuth } from 'firebase/auth';
import { getStorage } from 'firebase/storage';

const firebaseConfig = {
  apiKey: "API_KEY",
  authDomain: "PROJECT_ID.firebaseapp.com",
  projectId: "PROJECT_ID",
  storageBucket: "PROJECT_ID.appspot.com",
  messagingSenderId: "SENDER_ID",
  appId: "APP_ID"
};

const app = initializeApp(firebaseConfig);

export const db = getFirestore(app);
export const auth = getAuth(app);
export const storage = getStorage(app);

## 🔐 Firebase Authentication

// Signup
createUserWithEmailAndPassword(auth, email, password);

// Login
signInWithEmailAndPassword(auth, email, password);

// Logout
signOut(auth);

// Auth State
onAuthStateChanged(auth, user => {
  console.log(user ? "Logged in" : "Logged out");
});
🧰 Other Auth Methods
Google Auth: GoogleAuthProvider

GitHub, Twitter, Facebook (enable in Firebase Console)

Phone Auth (requires emulator or real device)

## 📦 Firestore (Database)

✅ CRUD Operations

import { collection, addDoc, getDocs, doc, updateDoc, deleteDoc } from 'firebase/firestore';

// Create
await addDoc(collection(db, "users"), {
  name: "Saim", age: 21
});

// Read
const querySnapshot = await getDocs(collection(db, "users"));
querySnapshot.forEach(doc => console.log(doc.id, doc.data()));

// Update
await updateDoc(doc(db, "users", "userId"), {
  age: 22
});

// Delete
await deleteDoc(doc(db, "users", "userId"));

## 🔍 Real-time Listener

import { onSnapshot } from 'firebase/firestore';

onSnapshot(collection(db, "users"), snapshot => {
  snapshot.docs.forEach(doc => console.log(doc.data()));
});

## 📁 Firebase Storage (Images, Files)

import { ref, uploadBytes, getDownloadURL } from 'firebase/storage';

const storageRef = ref(storage, `images/${file.name}`);

// Upload
await uploadBytes(storageRef, file);

// Get URL
const url = await getDownloadURL(storageRef);

## ✉️ Firebase Cloud Messaging (FCM)

Used for push notifications

Requires Service Worker on web

Use getToken() and onMessage() for receiving messages

## ⚙️ Firebase Hosting

🔧 Setup & Deploy

firebase init hosting
firebase deploy

## 🔁 Emulate Locally

firebase emulators:start

## 🔥 Firebase Realtime Database (Optional)


import { getDatabase, ref, set } from 'firebase/database';

const db = getDatabase();

// Write data
set(ref(db, 'users/' + userId), {
  username: "saim",
  email: "saim@example.com"
});

## 🧪 Firebase Emulator Suite

firebase init emulators
firebase emulators:start
Emulates:

## Firestore

- Auth

- Functions

- Hosting

- Realtime DB

- Great for local testing before deploying.


## 🔁 Firebase Functions (Cloud Functions)

firebase init functions
cd functions
npm install
firebase deploy --only functions

// Example (index.js)
exports.helloWorld = functions.https.onRequest((req, res) => {
  res.send("Hello from Firebase!");
});

## 📊 Firebase Analytics

Enabled by default in Firebase Console

Use logEvent('event_name') in frontend (Firebase SDK)

Works well with Google Analytics dashboard

## 🔒 Firebase Security Rules

🔐 Firestore Example

rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /users/{userId} {
      allow read, write: if request.auth.uid == userId;
    }
  }
}
Test rules using the Firebase Emulator Suite.

##🧰 Useful Tools & Links

-🔎 Firebase Console

-🧪 Emulator Docs

-📚 Firestore Docs

-📦 Storage Docs

-🔐 Auth Docs

-🔧 Functions Docs

-💻 Official SDKs

## ✅ Pro Tips

Use .env for sensitive keys (via Vite/Next.js)

Use Firestore indexes for advanced queries

Combine Firebase with React Query or SWR for caching

Use Emulator Suite for local testing

Set up billing for production features like phone auth or FCM


