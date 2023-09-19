import React, { useState } from 'react';

// Define a FileUpload component to handle file input
function FileUploadComponent({ onFileChange }) {
    // Check if the uploaded file is binary STL
    const isBinarySTL = (file) => {
        if (!file.name.endsWith('.stl')) {
            return false;
        }

        const reader = new FileReader();
        reader.onload = (event) => {
            const dataView = new DataView(event.target.result);
            for (let i = 0; i < 5; i++) {
                if (dataView.getUint8(i) === 0x73 || dataView.getUint8(i) === 0x53) {
                    return true;
                }
            }
            return false;
        };
        reader.readAsArrayBuffer(file.slice(0, 80));
        return false;
    };

    // Handle file change event
    const handleFileChange = (e) => {
        const file = e.target.files[0];

        // Check if the selected file is binary STL
        const isBinary = isBinarySTL(file);

        if (isBinary) {
            // Send the binary file to the binary STL route on the server
            sendBinarySTL(file);
        } else {
            // Send the file to the ASCII STL route on the server
            sendASCIISTL(file);
        }
    };

    // Function to send binary STL file to the server
    const sendBinarySTL = async (file) => {
        const formData = new FormData();
        formData.append('file', file);

        try {
            const response = await fetch('http://localhost:8080/process-binary-stl', {
                method: 'POST',
                body: formData,
            });

            if (!response.ok) {
                throw new Error('Error sending the file.');
            }

            const data = await response.json();
            onFileChange(data); // Callback with the analysis result
        } catch (error) {
            console.error(error);
        }
    };

    // Function to send ASCII STL file to the server
    const sendASCIISTL = async (file) => {
        const formData = new FormData();
        formData.append('file', file);

        try {
            const response = await fetch('http://localhost:8080/process-ascii-stl', {
                method: 'POST',
                body: formData,
            });

            if (!response.ok) {
                throw new Error('Error sending the file.');
            }

            const data = await response.json();
            onFileChange(data); // Callback with the analysis result
        } catch (error) {
            console.error(error);
        }
    };

    // Render a file input element
    return (
        <div>
            <input type="file" accept=".stl" onChange={handleFileChange} />
        </div>
    );
}

export default FileUploadComponent;