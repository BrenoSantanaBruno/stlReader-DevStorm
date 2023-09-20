import React, { useState } from 'react';

function FileUploadComponent({ onFileChange }) {
  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (file && file.name.endsWith('.stl')) {
      onFileChange(file);
    } else {
      onFileChange(null);
      alert('Please select a file with .stl extension.');
    }
  };

  return (
      <div className="input-group mb-3">
        <input type="file" accept=".stl" onChange={handleFileChange} className="form-control" />
      </div>
  );
}

function App() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [analysisResult, setAnalysisResult] = useState(null);

  const handleFileChange = (file) => {
    setSelectedFile(file);
  };

  const handleSubmit = async (url) => {
    if (!selectedFile) {
      alert('Please select a file before analyzing.');
      return;
    }

    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await fetch(url, {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Error sending the file.');
      }

      const data = await response.json();
      setAnalysisResult(data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
      <div className="container mt-5">
        <h1 className="text-center">STL File Upload - Analysis</h1>
        <div className="row justify-content-center">
          <div className="col-md-6">
            <FileUploadComponent onFileChange={handleFileChange} />
            <div className="mt-2 d-flex justify-content-between">
              <button className="btn btn-primary" type="button" onClick={() => handleSubmit('http://localhost:8080/process-ascii-stl')}>
                STL ASCII
              </button>
              <button className="btn btn-success" type="button" onClick={() => handleSubmit('http://localhost:8080/process-binary-stl')}>
                STL Binary
              </button>
            </div>
            {analysisResult && (
                <div className="alert alert-success mt-3">
                  <h2>Analysis Result</h2>
                  <p>Number of Triangles: {analysisResult.numTriangles}</p>
                  <p>Total Area: {analysisResult.areaTotal}</p>
                </div>
            )}
          </div>
        </div>
      </div>
  );
}

export default App;