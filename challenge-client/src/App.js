import React, { useState } from 'react';

function App() {
  const [file, setFile] = useState(null);
  const [result, setResult] = useState(null);

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
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
              <input type="file" accept=".stl" className="form-control" onChange={handleFileChange} />
              <button className="btn btn-primary" type="button" onClick={handleSubmit}>Analysis</button>
            </div>
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