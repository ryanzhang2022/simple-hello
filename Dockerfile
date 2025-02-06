# Use a lightweight base image
FROM alpine:latest

# Create a working directory
WORKDIR /app

# Copy the pre-built Go executable
COPY simple-hello .

# Make the executable runnable
RUN chmod +x simple-hello

# Set the command to run the executable
ENTRYPOINT ["/app/simple-hello"]