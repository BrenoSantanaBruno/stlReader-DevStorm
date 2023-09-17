# Use the latest Node.js image as the base
FROM node:latest

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json (if it exists)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire source code into the working directory in the container
COPY . .

# Build the React application
RUN npm run build

# Expose the port where the React application will run (usually 3000)
EXPOSE 3000

# Start the React application when the container is started
CMD [ "npm", "start" ]
