# Use the official Node.js v16.15.0 image as the base image
FROM node:16.15.0

# Create a directory named 'app'
RUN mkdir frontend-angular

# Set the working directory to '/app'
WORKDIR /frontend-angular

COPY package.json package.json
COPY package-lock.json package-lock.json

# Install npm dependencies
RUN npm install

# Copy the content of the current directory into the container at '/app'
COPY . .

# Install Angular CLI
RUN npm install -g @angular/cli@v13.3.7

# Expose port 4200 to allow external connections
EXPOSE 4200

# Specify the default command to run when the container starts (assumes "npm watch" is a valid command)
CMD ["npm", "run", "start"]