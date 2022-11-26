Function Move-Pictures {

    [CmdletBinding()]
    Param(
        # Directory in which the media in need of sorting resides
        [Parameter(Mandatory = $true)][String]$FilePath

    )

    # Capture the current path so that we may return to it when finished sorting the photos
    $originalPath = (Get-Location).path

    # Enter the directory in which is to be sorted
    Set-Location $FilePath

    # If the file type is present, create a directory for it. JPG, raw files and videos
    if (!(Test-Path "$FilePath/jpg") -and ((Test-Path "*.jpg") -or (Test-Path "*.png") -or (Test-Path "*.jpeg"))) { New-Item -ItemType Directory jpg }
    if (!(Test-Path "$FilePath/raw") -and ((Test-Path "*.cr3") -or (Test-Path "*.cr2"))) { New-Item -ItemType Directory raw }
    if (!(Test-Path "$FilePath/videos") -and ((Test-Path "*.mp4") -or (Test-Path "*.m4v"))) { New-Item -ItemType Directory videos }

    # Move the files into their respective folders
    $files = Get-ChildItem
    $files | Where-Object { $_.name -like "*.jpg" -or $_.name -like "*.jpeg" -or $_.name -like "*.png" } | Move-Item -Destination ./jpg
    $files | Where-Object { $_.name -like "*.cr3" -or $_.name -like "*.cr2" } | Move-Item -Destination ./raw
    $files | Where-Object { $_.name -like "*.mp4" -or $_.name -like "*.m4v" -or $_.name -like "*.mov*" } | Move-Item -Destination ./videos

    # Move back to the original directory that this script was run from
    Set-Location $originalPath

}