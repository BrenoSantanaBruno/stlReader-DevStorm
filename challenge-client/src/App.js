import React, { useState } from 'react';

function App() {
  // Define two state variables, 'file' and 'result', using the 'useState' hook
  const [file, setFile] = useState(null);
  const [result, setResult] = useState(null);

  // Event handler for file input change
  const handleFileChange = (e) => {
    // Get the selected file from the input
    const selectedFile = e.target.files[0];
    // Update the 'file' state variable with the selected file
    setFile(selectedFile);
  };

  // Event handler for form submission
  const handleSubmit = async () => {
    // Create a FormData object to send the file
    const formData = new FormData();
    formData.append('file', file);

    try {
      // Send a POST request to the server with the selected file
      const response = await fetch('http://localhost:8080/upload', {
        method: 'POST',
        body: formData,
      });

      // Check if the response is successful (status code 2xx)
      if (!response.ok) {
        throw new Error('Error sending the file.');
      }

      // Parse the response JSON data
      const data = await response.json();
      // Update the 'result' state variable with the received data
      setResult(data);
    } catch (error) {
      // Handle any errors that occur during the process and log them
      console.error(error);
    }
  };

  return (
      <div className="App">
        <h1>STL File Upload</h1>
        {/* Input element for file selection */}
        <input type="file" accept=".stl" onChange={handleFileChange} />
        {/* Button to trigger file upload */}
        <button onClick={handleSubmit}>Upload File</button>

        {/* Display analysis result if 'result' is not null */}
        {result && (
            <div>
              <h2>Analysis Result</h2>
              <p>Number of Triangles: {result.numTriangles}</p>
              <p>Total Area: {result.areaTotal}</p>
            </div>
        )}
      </div>
  );
}

export default App;