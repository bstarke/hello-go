FROM scratch
# Copy our static executable.

COPY templates/index.html /app/templates/
COPY hello /app/
WORKDIR /app
# Run the hello binary.
ENTRYPOINT ["/app/hello"]
