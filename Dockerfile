# Use a lightweight base image
FROM cloud-boe-cn-beijing.cr.volces.com/custom_ns_1267075819788007/zxytest:zxy-alpine

# Create a working directory
WORKDIR /app

# Copy the pre-built Go executable
COPY simple-hello .

# Make the executable runnable
RUN chmod +x simple-hello

# Set the command to run the executable
ENTRYPOINT ["/app/simple-hello"]