// Define a function called 'reportWebVitals' that takes a callback function 'onPerfEntry'
const reportWebVitals = onPerfEntry => {
  // Check if 'onPerfEntry' is a function and not undefined
  if (onPerfEntry && onPerfEntry instanceof Function) {
    // Dynamically import the 'web-vitals' library
    import('web-vitals').then(({ getCLS, getFID, getFCP, getLCP, getTTFB }) => {
      // Call each of the 'get' functions from 'web-vitals' with 'onPerfEntry' as a callback
      getCLS(onPerfEntry);   // Cumulative Layout Shift
      getFID(onPerfEntry);   // First Input Delay
      getFCP(onPerfEntry);   // First Contentful Paint
      getLCP(onPerfEntry);   // Largest Contentful Paint
      getTTFB(onPerfEntry);  // Time to First Byte
    });
  }
};

// Export the 'reportWebVitals' function for use in other parts of the application
export default reportWebVitals;
