# Simple Go CI/CD Learning Project

This is a beginner-friendly project to learn CI/CD with GitHub Actions and Go.

## What's in this project?

- `main.go` - Simple web server that says "Hello, CI/CD World!"
- `main_test.go` - Test file to verify our code works
- `.github/workflows/pipeline.yml` - CI/CD pipeline configuration
- `go.mod` - Go module file

## How to run locally

```bash
# Run the server
go run main.go

# Visit http://localhost:8080 to see "Hello, CI/CD World!"

# Run tests
go test -v

# Build the app
go build -o app main.go
```

## What happens with CI/CD?

1. **When you push code to GitHub:**
   - GitHub Actions automatically runs
   - Tests your code
   - Builds your app
   - If everything passes, simulates deployment

2. **CI (Continuous Integration):**
   - Runs tests
   - Builds the application
   - Fails if tests don't pass

3. **CD (Continuous Deployment):**
   - Only runs if CI passes
   - Only runs on main branch
   - Simulates deploying your app

## Next Steps

1. Push this code to GitHub
2. Watch GitHub Actions run automatically
3. Make changes and see the pipeline run again!
