Function Match-Jpg {

    [CmdletBinding()]
    Param(
        
        [Parameter(Mandatory = $true)][String]$rawSourcePath,
        [Parameter(Mandatory = $true)][String]$jpgSourcePath,
        [Parameter(Mandatory = $true)][String]$DestinationPath

    )

    $jpgFiles = Get-ChildItem $jpgSourcePath

    foreach ($rawFile in (Get-ChildItem $rawSourcePath)) {
        foreach ($jpgFile in $jpgFiles) {
            if ((($rawFile.name) -replace $rawFile.Extension,"") -like (($jpgFile.name) -replace $jpgFile.Extension,"")) {
                Copy-Item $jpgFile.FullName -Destination $DestinationPath
            }
        }
    }
}

Function Match-Raw {

    [CmdletBinding()]
    Param(
        
        [Parameter(Mandatory = $true)][String]$rawSourcePath,
        [Parameter(Mandatory = $true)][String]$jpgSourcePath,
        [Parameter(Mandatory = $true)][String]$DestinationPath

    )

    $rawFiles = Get-ChildItem $rawSourcePath

    foreach ($jpgFile in (Get-ChildItem $jpgSourcePath)) {
        foreach ($rawFile in $rawFiles) {
            if ((($jpgFile.name) -replace $jpgFile.Extension,"") -like (($rawFile.name) -replace $rawFile.Extension,"")) {
                Copy-Item $rawFile.FullName -Destination $DestinationPath
            }
        }
    }
}