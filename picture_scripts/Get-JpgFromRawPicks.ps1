Function Match-Photos {

    [CmdletBinding()]
    Param(

        [Parameter(Mandatory = $true)][String]$SourcePath,
        [Parameter(Mandatory = $true)][String]$DestinationPath

    )

    $files = Get-Content $SourcePath

    foreach ($file in (Get-ChildItem $DestinationPath)) {
        foreach ($file2 in $files) {
            if ((($file.name) -replace $file.Extension,"") -like (($file2.name) -replace $file2.Extension,"")) {
                Copy-Item -LiteralPath $file.Name -Destination 'C:/strom/photos/chief_mann_retirement/file2s/file/' -WhatIf
            }
        }
    }
}