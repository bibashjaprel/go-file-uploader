# Go File Uploader

A simple and efficient file upload API built with Go and Gin framework, supporting both local storage and AWS S3.

## Features

- ✅ Local file upload with size limit (10MB)
- ✅ File download from local storage
- ✅ AWS S3 file upload
- ✅ RESTful API endpoints
- ✅ Multipart form data support

## Requirements

- Go 1.24.2 or higher
- AWS credentials (for S3 upload)

## Installation

```bash
git clone https://github.com/bibashjaprel/go-file-uploader.git
cd go-file-uploader
go mod download
```

## Configuration

### Environment Variables

The application uses environment variables for AWS configuration. You can set these in two ways:

1. **Using a `.env` file** (recommended for development):
   ```bash
   cp .env.example .env
   ```
   Then edit `.env` with your actual values:
   ```
   AWS_REGION=us-east-1
   AWS_BUCKET_NAME=your-actual-bucket-name
   ```

2. **Using system environment variables**:
   ```bash
   export AWS_REGION=us-east-1
   export AWS_BUCKET_NAME=your-actual-bucket-name
   export AWS_ACCESS_KEY_ID=your-access-key
   export AWS_SECRET_ACCESS_KEY=your-secret-key
   ```

### AWS Credentials

AWS credentials can be configured in multiple ways (in order of precedence):

1. Environment variables: `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`
2. AWS credentials file: `~/.aws/credentials`
3. IAM role (when running on AWS infrastructure)

For more information, see the [AWS SDK for Go configuration guide](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/).

## Usage

Run the server:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### Local Storage

**Upload File**
```bash
POST /upload
Content-Type: multipart/form-data

curl -X POST http://localhost:8080/upload \
  -F "file=@/path/to/your/file"
```

**Download File**
```bash
GET /uploads/:filename

curl http://localhost:8080/uploads/example.txt
```

### S3 Storage

**Upload to S3**
```bash
POST /s3/upload
Content-Type: multipart/form-data

curl -X POST http://localhost:8080/s3/upload \
  -F "file=@/path/to/your/file"
```

## Project Structure

```
go-file-uploader/
├── main.go              # Main application entry point
├── localstorage/        # Local file storage handlers
│   └── basic.go
├── s3storage/           # AWS S3 storage handlers
│   └── basic.go
├── go.mod               # Go module definition
└── README.md
```

## Future Enhancements

- [ ] Multiple file upload
- [ ] Streaming for large files
- [ ] Retry & circuit breaker mechanism
- [ ] Rate limiting and MIME type filtering
- [ ] Background file processing
- [ ] File compression
- [ ] Authentication and authorization
