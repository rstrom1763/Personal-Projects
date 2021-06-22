Function Get-Picks {

    param(
        [Parameter(Mandatory = $true)]$filePath,
        [Parameter(Mandatory = $true)]$pickList,
        [Parameter(Mandatory = $true)]$finalDir
    )

    if (!(Test-Path $finalDir)) {
        New-Item -Path $finalDir -ItemType Directory
    }

    $picklist = Get-Content $pickList
    $found = $false

    foreach ($pick in $pickList) {
        
        $found = $false
        foreach ($photo in (Get-ChildItem $filePath)) {

            if ($photo.Name -like "*$pick*") {
                Copy-Item -Path "$filePath\$photo" -Destination $finalDir
                $found = $true
            }
            if ($found) {
                break
            }

        }

        if (!$found) {
            Write-Host "Could not find $pick :( "
        }

    }
    
}