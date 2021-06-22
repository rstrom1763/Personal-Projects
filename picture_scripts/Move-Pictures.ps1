﻿Function Move-Pictures {

    [CmdletBinding()]
    Param(

        [Parameter(Mandatory = $true)][String]$FilePath

    )
    $originalPath = (Get-Location).path
    Set-Location $FilePath

    if (!(Test-Path "$FilePath/jpg") -and ((Test-Path "*.jpg") -or (Test-Path "*.png") -or (Test-Path "*.jpeg"))) { New-Item -ItemType Directory jpg }
    if (!(Test-Path "$FilePath/raw") -and (Test-Path "*.cr3")) { New-Item -ItemType Directory raw }
    if (!(Test-Path "$FilePath/videos") -and (Test-Path "*.mp4")) { New-Item -ItemType Directory videos }

    Get-ChildItem | Where-Object { $_.name -like "*.jpg" -or $_.name -like "*.jpeg" -or $_.name -like "*.png"} | Move-Item -Destination ./jpg
    Get-ChildItem | Where-Object { $_.name -like "*.cr3" } | Move-Item -Destination ./raw
    Get-ChildItem | Where-Object { $_.name -like "*.mp4" } | Move-Item -Destination ./videos

    Set-Location $originalPath

}