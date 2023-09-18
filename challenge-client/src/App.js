import React, { useState } from 'react';

// Define a FileUpload component to handle file input
function FileUpload({ onFileChange }) {
  // Handle file change event
  const handleFileChange = (e) => {
    const file = e.target.files[0];

    // Check if the selected file has the .stl extension
    if (file && file.name.endsWith('.stl')) {
      onFileChange(file); // Call the provided onFileChange callback with the file
    } else {
      onFileChange(null); // If the file doesn't have the .stl extension, set file to null
      alert('Please select a file with the .stl extension.'); // Display an alert to the user
    }
  };

  // Render a file input element
  return (
      <div>
        <input type="file" accept=".stl" onChange={handleFileChange} />
      </div>
  );
}

// Define the main App component
function App() {
  const [file, setFile] = useState(null); // State to store the selected file
  const [result, setResult] = useState(null); // State to store analysis result

  // Callback function to handle file change
  const handleFileChange = (selectedFile) => {
    setFile(selectedFile); // Set the selected file in the component's state
  };

  // Function to handle form submission
  const handleSubmit = async () => {
    const formData = new FormData();
    formData.append('file', file); // Append the selected file to the form data

    try {
      // Send a POST request with the file to the server
      const response = await fetch('http://localhost:8080/upload', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Error sending the file.'); // Handle errors if the response is not OK
      }

      const data = await response.json(); // Parse the response JSON
      setResult(data); // Set the analysis result in the component's state
    } catch (error) {
      console.error(error); // Log any errors that occur during the request
    }
  };

  // Render the main application UI
  return (
      <div className="container mt-5">
        <h1 className="text-center">STL File Upload - Analysis</h1>
        <div className="row justify-content-center">
          <div className="col-md-6">
            <div className="input-group mb-3">
              {/* Use the FileUpload component here */}
              <FileUpload onFileChange={handleFileChange} />
            </div>
            <button className="btn btn-primary" type="button" onClick={handleSubmit}>
              Analysis
            </button>
            {result && (
                <div className="alert alert-success mt-3">
                  <h2>Analysis Result</h2>
                  <p>Number of Triangles: {result.numTriangles}</p>
                  <p>Total Area: {result.areaTotal}</p>
                </div>
            )}
          </div>
        </div>
      </div>
  );
}

export default App;