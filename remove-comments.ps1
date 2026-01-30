# Script to remove comments from source files

function Remove-JSComments {
    param([string]$content)
    
    # Remove multi-line comments /* */
    $content = $content -replace '(?s)/\*.*?\*/', ''
    
    # Remove single-line comments // but preserve URLs like https://
    $content = $content -replace '(?m)(?<!:)//.*$', ''
    
    # Clean up extra blank lines (more than 2 consecutive)
    $content = $content -replace '(?m)^\s*$(\r?\n){2,}', "`r`n`r`n"
    
    return $content
}

function Remove-GoComments {
    param([string]$content)
    
    # Remove multi-line comments /* */
    $content = $content -replace '(?s)/\*.*?\*/', ''
    
    # Remove single-line comments // but preserve URLs
    $content = $content -replace '(?m)(?<!:)//.*$', ''
    
    # Clean up extra blank lines
    $content = $content -replace '(?m)^\s*$(\r?\n){2,}', "`r`n`r`n"
    
    return $content
}

function Remove-VueComments {
    param([string]$content)
    
    # Remove HTML comments <!-- -->
    $content = $content -replace '(?s)<!--.*?-->', ''
    
    # Remove multi-line comments /* */
    $content = $content -replace '(?s)/\*.*?\*/', ''
    
    # Remove single-line comments // but preserve URLs
    $content = $content -replace '(?m)(?<!:)//.*$', ''
    
    # Clean up extra blank lines
    $content = $content -replace '(?m)^\s*$(\r?\n){2,}', "`r`n`r`n"
    
    return $content
}

function Remove-CSSComments {
    param([string]$content)
    
    # Remove CSS comments /* */
    $content = $content -replace '(?s)/\*.*?\*/', ''
    
    # Clean up extra blank lines
    $content = $content -replace '(?m)^\s*$(\r?\n){2,}', "`r`n`r`n"
    
    return $content
}

# Get all source files
$projectRoot = "c:\Users\admin\Desktop\hello\Split-Chill"
$extensions = @("*.js", "*.ts", "*.vue", "*.go", "*.css", "*.scss", "*.jsx", "*.tsx")
$excludeDirs = @("node_modules", ".git", "storage", "log", "data")

$files = @()
foreach ($ext in $extensions) {
    $foundFiles = Get-ChildItem -Path $projectRoot -Filter $ext -Recurse -File | 
        Where-Object { 
            $exclude = $false
            foreach ($dir in $excludeDirs) {
                if ($_.FullName -like "*\$dir\*") {
                    $exclude = $true
                    break
                }
            }
            -not $exclude
        }
    $files += $foundFiles
}

Write-Host "Found $($files.Count) files to process"

$processedCount = 0
foreach ($file in $files) {
    try {
        $content = Get-Content -Path $file.FullName -Raw -Encoding UTF8
        $originalContent = $content
        
        $extension = $file.Extension.ToLower()
        
        switch ($extension) {
            {$_ -in ".js", ".ts", ".jsx", ".tsx"} {
                $content = Remove-JSComments $content
            }
            ".go" {
                $content = Remove-GoComments $content
            }
            ".vue" {
                $content = Remove-VueComments $content
            }
            {$_ -in ".css", ".scss"} {
                $content = Remove-CSSComments $content
            }
        }
        
        # Only write if content changed
        if ($content -ne $originalContent) {
            Set-Content -Path $file.FullName -Value $content -Encoding UTF8 -NoNewline
            $processedCount++
            Write-Host "Processed: $($file.FullName)"
        }
    }
    catch {
        Write-Host "Error processing $($file.FullName): $_" -ForegroundColor Red
    }
}

Write-Host "`nCompleted! Processed $processedCount files."
