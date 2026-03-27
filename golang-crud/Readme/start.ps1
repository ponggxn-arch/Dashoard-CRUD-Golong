# Golang CRUD System - Quick Start Script for Windows
# This script helps you get started quickly

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Golang CRUD System - Quick Start" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Check if Go is installed
Write-Host "[1/5] Checking Go installation..." -ForegroundColor Yellow
try {
    $goVersion = go version
    Write-Host "✓ Go is installed: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "✗ Go is not installed!" -ForegroundColor Red
    Write-Host "Please download and install Go from: https://go.dev/dl/" -ForegroundColor Red
    exit 1
}

# Check if MySQL is running
Write-Host ""
Write-Host "[2/5] Checking MySQL service..." -ForegroundColor Yellow
$mysqlService = Get-Service -Name "MySQL*" -ErrorAction SilentlyContinue
if ($mysqlService) {
    if ($mysqlService.Status -eq "Running") {
        Write-Host "✓ MySQL service is running" -ForegroundColor Green
    } else {
        Write-Host "! MySQL service is not running. Starting..." -ForegroundColor Yellow
        Start-Service $mysqlService.Name
        Write-Host "✓ MySQL service started" -ForegroundColor Green
    }
} else {
    Write-Host "! MySQL service not found. Please install MySQL." -ForegroundColor Yellow
}

# Download Go dependencies
Write-Host ""
Write-Host "[3/5] Downloading Go dependencies..." -ForegroundColor Yellow
go mod download
if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Dependencies downloaded successfully" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to download dependencies" -ForegroundColor Red
    exit 1
}

# Tidy up Go modules
Write-Host ""
Write-Host "[4/5] Tidying Go modules..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Go modules tidied successfully" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to tidy Go modules" -ForegroundColor Red
    exit 1
}

# Display next steps
Write-Host ""
Write-Host "[5/5] Setup Complete!" -ForegroundColor Green
Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Next Steps:" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "1. Create MySQL database:" -ForegroundColor White
Write-Host "   mysql -u root -p < database.sql" -ForegroundColor Gray
Write-Host ""
Write-Host "2. Update database credentials in main.go if needed" -ForegroundColor White
Write-Host "   (Default: root:@tcp(127.0.0.1:3306)/golang_crud)" -ForegroundColor Gray
Write-Host ""
Write-Host "3. Run the application:" -ForegroundColor White
Write-Host "   go run main.go" -ForegroundColor Gray
Write-Host ""
Write-Host "4. Open browser:" -ForegroundColor White
Write-Host "   http://localhost:8080" -ForegroundColor Gray
Write-Host ""
Write-Host "5. Login with default credentials:" -ForegroundColor White
Write-Host "   Username: admin" -ForegroundColor Gray
Write-Host "   Password: admin123" -ForegroundColor Gray
Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Ask if user wants to run the application now
$run = Read-Host "Do you want to run the application now? (y/n)"
if ($run -eq "y" -or $run -eq "Y") {
    Write-Host ""
    Write-Host "Starting application..." -ForegroundColor Yellow
    Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
    Write-Host ""
    go run main.go
}
