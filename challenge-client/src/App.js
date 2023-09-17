import React, { useState } from 'react';

// Move the definition of the FileUpload component outside of the App component
function FileUpload({ onFileChange }) {
  const handleFileChange = (e) => {
    const file = e.target.files[0];

    if (file && file.name.endsWith('.stl')) {
      onFileChange(file);
    } else {
      onFileChange(null);
      alert('Please select a file with the .stl extension.');
    }
  };

  return (
      <div>
        <input type="file" accept=".stl" onChange={handleFileChange} />
      </div>
  );
}

function App() {
  const [file, setFile] = useState(null);
  const [result, setResult] = useState(null);

  const handleFileChange = (selectedFile) => {
    setFile(selectedFile);
  };

  const handleSubmit = async () => {
    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await fetch('http://localhost:8080/upload', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Error sending the file.');
      }

      const data = await response.json();
      setResult(data);
    } catch (error) {
      console.error(error);
    }
  };

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