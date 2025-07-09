# 🔥 Firebase Developer Cheat Sheet (Basic to Advanced)

## 📦 Getting Started
- Go to: https://console.firebase.google.com
- Create a Firebase Project
- Add Web/Android/iOS App to Firebase
- Install Firebase CLI:
  ```bash
  npm install -g firebase-tools
🚀 Firebase CLI Commands
bash
Copy
Edit
firebase login                     # Log in to Firebase
firebase init                     # Initialize project
firebase deploy                   # Deploy to Firebase
firebase emulators:start          # Start local emulator
firebase use --add                # Add project alias
firebase logout                   # Log out
⚙️ SDK Setup (Web)
bash
Copy
Edit
npm install firebase
js
Copy
Edit
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
🔐 Firebase Authentication
js
Copy
Edit
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

GitHub, Twitter, Facebook (enabled via Firebase Console)

Phone Auth (requires Firebase emulator or real device)

📦 Firestore (Database)
✅ CRUD Operations
js
Copy
Edit
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
🔍 Real-time Listener
js
Copy
Edit
import { onSnapshot } from 'firebase/firestore';

onSnapshot(collection(db, "users"), snapshot => {
  snapshot.docs.forEach(doc => console.log(doc.data()));
});
📁 Firebase Storage (Images, Files)
js
Copy
Edit
import { ref, uploadBytes, getDownloadURL } from 'firebase/storage';

const storageRef = ref(storage, `images/${file.name}`);

// Upload
await uploadBytes(storageRef, file);

// Get URL
const url = await getDownloadURL(storageRef);
✉️ Firebase Cloud Messaging (FCM)
Used for push notifications

Requires Service Worker on web

Use getToken() and onMessage() for receiving messages

⚙️ Firebase Hosting
🔧 Setup & Deploy
bash
Copy
Edit
firebase init hosting
firebase deploy
🔁 Emulate Locally
bash
Copy
Edit
firebase emulators:start
🔥 Firebase Realtime Database (Optional)
js
Copy
Edit
import { getDatabase, ref, set } from 'firebase/database';

const db = getDatabase();

// Write data
set(ref(db, 'users/' + userId), {
  username: "saim",
  email: "saim@example.com"
});
🧪 Firebase Emulator Suite
bash
Copy
Edit
firebase init emulators
firebase emulators:start
Emulates:

Firestore

Auth

Functions

Hosting

Realtime DB

Useful for local testing without deploying.

🔁 Firebase Functions (Cloud Functions)
bash
Copy
Edit
firebase init functions
cd functions
npm install
firebase deploy --only functions
js
Copy
Edit
// Example (index.js)
exports.helloWorld = functions.https.onRequest((req, res) => {
  res.send("Hello from Firebase!");
});
📊 Firebase Analytics
Enabled by default in Firebase Console

Use logEvent('event_name') in frontend (Firebase SDK)

Works well with Google Analytics dashboard

🔒 Firebase Security Rules
🔐 Firestore Example:
js
Copy
Edit
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /users/{userId} {
      allow read, write: if request.auth.uid == userId;
    }
  }
}
Test rules using the Firebase Emulator Suite.

🧰 Useful Tools & Links
🔎 Firebase Console

🧪 Emulator Docs

📚 Firestore Docs

📦 Storage Docs

🔐 Auth Docs

🔧 Functions Docs

💻 Official SDKs

✅ Pro Tips
Use .env for sensitive keys (via Vite/Next.js)

Use Firestore indexes for advanced queries

Combine Firebase with React Query or SWR for caching

Use Emulator for local testing before production deploys

Setup billing for production features like phone auth, FCM

🔥 Save this as FIREBASE_CHEATSHEET.md in your project root or docs folder. Use it as a quick reference!

yaml
Copy
Edit

---

Let me know if you'd like a **PDF version** or want to **combine this with the Go + Firebase integration guid
