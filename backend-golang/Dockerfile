FROM golang:1.18.10-buster

WORKDIR /

# Create celer-api folder
RUN mkdir backend-golang

# Set working directory
WORKDIR /backend-golang


#COPY ../.env ./

# Copy only dependency files
COPY go.mod go.sum ./

# Compile the application
RUN go mod download

# Copy the rest of the code
COPY . .

# Install air
RUN go install github.com/cosmtrek/air

# In the following port the application will be available
EXPOSE 5000

# Use air to run the application during development
CMD ["air", "-c", ".air.toml"]