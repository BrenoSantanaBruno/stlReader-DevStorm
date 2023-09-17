// Import React and ReactDOM libraries for creating and rendering React components
import React from 'react';
import ReactDOM from 'react-dom/client';

// Import custom CSS styles and the main 'App' component
import './index.css';
import App from './App';

// Import a function for measuring web vitals
import reportWebVitals from './reportWebVitals';

// Import Bootstrap CSS styles for styling the application
import 'bootstrap/dist/css/bootstrap.min.css';

// Create a root element in the HTML with the id 'root'
const root = ReactDOM.createRoot(document.getElementById('root'));

// Render the main React application within a 'StrictMode' component
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals

// Measure and report web vitals for performance monitoring
reportWebVitals();
